package database

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

func CreatePost(tx *sql.Tx, newPost models.Post) (err error) {
	fmt.Println("here")
	_, err = tx.Exec(`
		INSERT INTO post (user_id, title, body, image, created_at) VALUES (?, ?, ?, ?, ?)
		`, newPost.UserID, newPost.Title, newPost.Body, newPost.Image, newPost.Created,
	)
	if err != nil {
		return err
	}
	//if post does not appear in front page automatically. We need to return it from here maybe hehe

	return nil
}
