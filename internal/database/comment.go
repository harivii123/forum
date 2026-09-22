package database

import (
	"database/sql"
	"forum/internal/models"
)

func CreateComment(tx *sql.Tx, newComment models.Comment)error{

	_, err :=tx.Exec(`
		INSERT INTO comment (user_id, post_id, body, created_at) VALUES (?, ?, ?, ?)
		`, newComment.User_id, newComment.Post_id, newComment.Body, newComment.Created,
	)
	if err != nil {
		return err
	}	
	
	return nil
}
