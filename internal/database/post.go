package database

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

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
