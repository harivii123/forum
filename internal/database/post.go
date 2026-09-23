package database

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/models"
	"time"
)

type PostSearch struct {
	ID               int
	Title            string
	PosterID         int
	Body             string
	Image            string
	Created          time.Time
	Poster           string
	PosterAvatar     string
	Comments         string
	CommentLikers    string
	CommentDislikers string
	PostLikers       string
	PostDislikers    string
	PostCategories   string
}

func FilterPosts(filters string, ctx context.Context, tx *sql.Tx, args ...any) ([]PostSearch, error) {
	query := `SELECT p.id, p.title, p.user_id, p.body, COALESCE(p.image, ''), p.created_at, u.username, COALESCE(u.profile_picture, '/static/avatars/default-avatar.png'),
		COALESCE((SELECT group_concat(c.id || char(31) || cu.id || char(31) || cu.username || char(31) || COALESCE(cu.profile_picture, '/static/avatars/default-avatar.png') || char(31) || c.body, char(30))
				FROM comment c JOIN user cu ON cu.id = c.user_id
				WHERE c.post_id = p.id), '') AS comments,
		COALESCE((SELECT group_concat(cv.comment_id || char(31) || cvu.id || char(31) || cvu.username || char(31) || COALESCE(cvu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
				FROM comment_vote cv JOIN user cvu ON cv.user_id = cvu.id JOIN comment c ON c.id = cv.comment_id
				WHERE c.post_id = p.id AND cv.vote = 1), '') AS comment_likers,
		COALESCE((SELECT group_concat(cv.comment_id || char(31) || cvu.id || char(31) || cvu.username || char(31) || COALESCE(cvu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
				FROM comment_vote cv JOIN user cvu ON cv.user_id = cvu.id JOIN comment c ON c.id = cv.comment_id
				WHERE c.post_id = p.id AND cv.vote = -1), '') AS comment_dislikers,
		COALESCE((SELECT group_concat(vu.id || char(31) || vu.username || char(31) || COALESCE(vu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
			FROM post_vote pv JOIN user vu ON vu.id = pv.user_id
			WHERE pv.post_id = p.id AND pv.vote = 1), '') AS likers,
		COALESCE((SELECT group_concat(vu.id || char(31) || vu.username || char(31) || COALESCE(vu.profile_picture, '/static/avatars/default-avatar.png'), char(30))
			FROM post_vote pv JOIN user vu ON vu.id = pv.user_id
			WHERE pv.post_id = p.id AND pv.vote = -1), '') AS dislikers,
		COALESCE((SELECT group_concat(cat.id || char(31) || cat.name || char(31) || cat.type, char(30))
			FROM category_post cp JOIN category cat ON cat.id = cp.category_id
			WHERE cp.post_id = p.id), '') AS categories
		FROM post p
		JOIN user u ON u.id = p.user_id
		JOIN post_fts fts ON fts.rowid = p.id`

	rows, err := tx.QueryContext(ctx, query+filters, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []PostSearch{}

	for rows.Next() {
		var postSearch PostSearch
		err := rows.Scan(&postSearch.ID, &postSearch.Title, &postSearch.PosterID, &postSearch.Body, &postSearch.Image, &postSearch.Created, &postSearch.Poster, &postSearch.PosterAvatar, &postSearch.Comments, &postSearch.CommentLikers, &postSearch.CommentDislikers, &postSearch.PostLikers, &postSearch.PostDislikers, &postSearch.PostCategories)
		if err != nil {
			return nil, err
		}
		result = append(result, postSearch)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}

func GetPostsByCategory(db *sql.DB, categoryName string) ([]models.Post, error) {
	rows, err := db.Query(`
        SELECT p.id, p.user_id, p.title, p.body, p.created
        FROM post p
        JOIN category_post cp ON cp.post_id = p.id
        JOIN category c       ON c.id = cp.category_id
        WHERE c.name = ?
        ORDER BY p.created DESC`, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.Created); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

// creates a post to db
func CreatePost(tx *sql.Tx, newPost models.Post) (err error) {
	// fmt.Println("here")
	_, err = tx.Exec(`
		INSERT INTO post (user_id, title, body, created_at) VALUES (?, ?, ?, ?)
		`, newPost.UserID, newPost.Title, newPost.Body, newPost.Created,
	)
	if err != nil {
		fmt.Println("er")
		return err
	}
	//if post does not appear in front page automatically. We need to return it from here

	return nil
}
