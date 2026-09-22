package service

import (
	"database/sql"
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
	"time"
)

func CreateComment(newComment models.Comment, cookie *http.Cookie, db *sql.DB) (error, int) {

	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	defer tx.Rollback()

	userID, err := database.GetUserIDWithToken(tx, cookie.Value)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	newComment.User_id, newComment.Created = userID, time.Now()

	err = database.CreateComment(tx, newComment)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusSeeOther
}
