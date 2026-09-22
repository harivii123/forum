package service

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
)

func CreatePost(newPost models.Post, db *sql.DB) (error, int) {

	if newPost.Body == "" || newPost.Title == "" {
		return errors.New("Empty post"), http.StatusBadRequest // do nothing
	}

	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	defer tx.Rollback()

	err = database.CreatePost(tx, newPost)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	err = tx.Commit()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	fmt.Println(newPost)
	return nil, http.StatusOK
}
