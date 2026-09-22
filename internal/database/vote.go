package database

import (
	"database/sql"
)

func GetPostVote(tx *sql.Tx, userID, postID int) (int, error) {
	var vote int

	err := tx.QueryRow(
		"SELECT user_id FROM post_vote WHERE post_id=? AND vote=?",
	)
}
