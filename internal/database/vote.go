package database

import (
	"database/sql"
	"errors"
)

func GetPostVote(tx *sql.Tx, userID, postID int) (int, error) {
	var vote int

	err := tx.QueryRow(
		"SELECT vote FROM post_vote WHERE user_id=? AND post_id=?", userID, postID,
	).Scan(&vote)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return vote, nil
}

func InsertPostVote(tx *sql.Tx, userID, postID, vote int) error {
	_, err := tx.Exec(
		"INSERT INTO post_vote (user_id, post_id, vote) VALUES(?, ?, ?)", userID, postID, vote,
	)

	return err
}

func UpdatePostVote(tx *sql.Tx, userID, postID, vote int) error {
	_, err := tx.Exec(
		"UPDATE post_vote SET vote=? WHERE user_id=? AND post_id=?", vote, userID, postID,
	)

	return err
}

func DeletePostVote(tx *sql.Tx, userID, postID int) error {
	_, err := tx.Exec(
		"DELETE FROM post_vote WHERE user_id=? AND post_id=?", userID, postID,
	)

	return err
}
