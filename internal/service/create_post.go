package service

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
	"time"
)

func CreatePost(cookie *http.Cookie, newPost models.Post, db *sql.DB) (error, int) {

	if newPost.Body == "" || newPost.Title == "" {
		return errors.New("Empty post"), http.StatusBadRequest // do nothing
	}

	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	defer tx.Rollback()

	userID, err := database.GetUserIDWithToken(tx, cookie.Value)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	newPost.UserID, newPost.Created = userID, time.Now()

	err = database.CreatePost(tx, newPost)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	err = tx.Commit()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	fmt.Println(newPost)
	return nil, http.StatusSeeOther
}
