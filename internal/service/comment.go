package service

import (
	"database/sql"
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
	"time"
	"context"
)

func CreateComment(ctx context.Context, newComment models.Comment, cookie *http.Cookie, db *sql.DB) (error, int) {
	if newComment.Body==""{
		return nil, http.StatusSeeOther
	}
	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}
	defer tx.Rollback()

	userID, err := database.GetUserIDWithToken(ctx, tx, cookie.Value)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	newComment.User_id, newComment.Created = userID, time.Now()
	err = database.CreateComment(tx, newComment)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	err = tx.Commit()
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusSeeOther
}
