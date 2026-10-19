package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// every seeded user logs in with this password
const seedPassword = "password123"

type seedCategory struct {
	name string
	kind string // "genres" or "authors", the filters page shows these two types
}

type seedUser struct {
	username string
	email    string
	bio      string
	avatar   string // empty = default avatar
}

type seedComment struct {
	user string
	body string
}

type seedPost struct {
	user       string
	title      string
	body       string
	image      string // empty = post without an image
	daysAgo    int
	categories []string
	comments   []seedComment
}

// Seed fills an empty database with example data.
// If posts already exist nothing is inserted, delete database.db to seed again.
func Seed(db *sql.DB) error {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM post").Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		log.Println("Seed skipped: database already has posts")
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	categoryIDs, err := insertSeedCategories(tx)
	if err != nil {
		return fmt.Errorf("seeding categories: %w", err)
	}

	userIDs, err := insertSeedUsers(tx)
	if err != nil {
		return fmt.Errorf("seeding users: %w", err)
	}

	// fixed seed so every run creates the same "random" votes and times
	rng := rand.New(rand.NewSource(42))
	if err := insertSeedPosts(tx, rng, userIDs, categoryIDs); err != nil {
		return fmt.Errorf("seeding posts: %w", err)
	}

	return tx.Commit()
}

func insertSeedCategories(tx *sql.Tx) (map[string]int, error) {
	ids := map[string]int{}
	for _, c := range seedCategories {
		res, err := tx.Exec(`INSERT INTO category (name, type) VALUES (?, ?)`, c.name, c.kind)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", c.name, err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		ids[c.name] = int(id)
	}
	return ids, nil
}

func insertSeedUsers(tx *sql.Tx) (map[string]int, error) {
	ids := map[string]int{}
	for _, u := range seedUsers {
		// every user gets their own hash because bcrypt adds a random salt
		hash, err := bcrypt.GenerateFromPassword([]byte(seedPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		var avatar any // nil -> NULL -> default avatar is shown
		if u.avatar != "" {
			avatar = u.avatar
		}

		res, err := tx.Exec(`INSERT INTO user (username, email, password_hash, profile_picture, bio) VALUES (?, ?, ?, ?, ?)`,
			u.username, u.email, string(hash), avatar, u.bio)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", u.username, err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		ids[u.username] = int(id)
	}
	return ids, nil
}

func insertSeedPosts(tx *sql.Tx, rng *rand.Rand, userIDs, categoryIDs map[string]int) error {
	// a slice keeps the voter order the same on every run, map order is random
	allUsers := make([]int, 0, len(seedUsers))
	for _, u := range seedUsers {
		allUsers = append(allUsers, userIDs[u.username])
	}

	postVote, err := tx.Prepare(`INSERT INTO post_vote (user_id, post_id, vote) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer postVote.Close()

	commentVote, err := tx.Prepare(`INSERT INTO comment_vote (user_id, comment_id, vote) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer commentVote.Close()

	now := time.Now()

	for _, p := range seedPosts {
		authorID, ok := userIDs[p.user]
		if !ok {
			return fmt.Errorf("post %q: unknown user %q", p.title, p.user)
		}

		var image any
		if p.image != "" {
			image = p.image
		}

		created := now.Add(-time.Duration(p.daysAgo)*24*time.Hour - time.Duration(rng.Intn(600))*time.Minute)
		res, err := tx.Exec(`INSERT INTO post (user_id, title, body, image, created_at) VALUES (?, ?, ?, ?, ?)`,
			authorID, p.title, p.body, image, created)
		if err != nil {
			return fmt.Errorf("post %q: %w", p.title, err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		postID := int(id)

		for _, name := range p.categories {
			categoryID, ok := categoryIDs[name]
			if !ok {
				return fmt.Errorf("post %q: unknown category %q", p.title, name)
			}
			if _, err := tx.Exec(`INSERT INTO category_post (category_id, post_id) VALUES (?, ?)`, categoryID, postID); err != nil {
				return err
			}
		}

		if err := randomVotes(postVote, postID, authorID, allUsers, 25, rng); err != nil {
			return err
		}

		// comments are spread evenly between the post time and now (max 4 days),
		// so they stay in the order they are written in seed_data.go
		window := min(now.Sub(created)-30*time.Minute, 4*24*time.Hour)
		step := window / time.Duration(len(p.comments)+1)

		for i, c := range p.comments {
			commenterID, ok := userIDs[c.user]
			if !ok {
				return fmt.Errorf("comment on %q: unknown user %q", p.title, c.user)
			}

			at := created.Add(step*time.Duration(i+1) + time.Duration(rng.Int63n(int64(step/2)+1)))
			res, err := tx.Exec(`INSERT INTO comment (user_id, post_id, body, created_at) VALUES (?, ?, ?, ?)`,
				commenterID, postID, c.body, at)
			if err != nil {
				return fmt.Errorf("comment on %q: %w", p.title, err)
			}
			commentID, err := res.LastInsertId()
			if err != nil {
				return err
			}

			if err := randomVotes(commentVote, int(commentID), commenterID, allUsers, 10, rng); err != nil {
				return err
			}
		}
	}
	return nil
}

// randomVotes lets a random group of users like or dislike one post or comment.
// Every target gets its own like chance, so some end up loved and some divisive.
func randomVotes(stmt *sql.Stmt, targetID, authorID int, users []int, maxVoters int, rng *rand.Rand) error {
	likeChance := 0.5 + rng.Float64()*0.45
	voters := rng.Intn(min(maxVoters, len(users)) + 1)

	for _, i := range rng.Perm(len(users))[:voters] {
		if users[i] == authorID {
			continue // nobody votes on their own writing
		}
		vote := 1
		if rng.Float64() > likeChance {
			vote = -1
		}
		if _, err := stmt.Exec(users[i], targetID, vote); err != nil {
			return err
		}
	}
	return nil
}
