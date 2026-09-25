package service

import (
	"database/sql"
	"errors"
	"forum/internal/database"
)

var ErrPostNotFound = errors.New("post not found")
var ErrInvalidVote = errors.New("invalid vote value")

func SetPostVote(db *sql.DB, userID, postID, vote int) error {

	if vote != 1 && vote != -1 {
		return ErrInvalidVote
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	exists, err := database.PostExists(tx, postID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrPostNotFound
	}

	current, err := database.GetPostVote(tx, userID, postID)
	if err != nil {
		return err
	}

	switch {
	case current == 0:
		err = database.InsertPostVote(tx, userID, postID, vote)
	case current == vote:
		err = database.DeletePostVote(tx, userID, postID)
	default:
		err = database.UpdatePostVote(tx, userID, postID, vote)
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}
