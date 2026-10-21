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

func CommentExists(tx *sql.Tx, commentID int) (bool, error) {
	var n int
	err := tx.QueryRow("SELECT 1 FROM comment WHERE id = ?", commentID).Scan(&n)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetCommentVote(tx *sql.Tx, userID, commentID int) (int, error) {
	var vote int

	err := tx.QueryRow(
		"SELECT vote FROM comment_vote WHERE user_id=? AND comment_id=?", userID, commentID,
	).Scan(&vote)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return vote, nil
}

func InsertCommentVote(tx *sql.Tx, userID, commentID, vote int) error {
	_, err := tx.Exec(
		"INSERT INTO comment_vote (user_id, comment_id, vote) VALUES(?, ?, ?)", userID, commentID, vote,
	)

	return err
}

func UpdateCommentVote(tx *sql.Tx, userID, commentID, vote int) error {
	_, err := tx.Exec(
		"UPDATE comment_vote SET vote=? WHERE user_id=? AND comment_id=?", vote, userID, commentID,
	)
	return err
}

func DeleteCommentVote(tx *sql.Tx, userID, commentID int) error {
	_, err := tx.Exec(
		"DELETE FROM comment_vote WHERE user_id=? AND comment_id=?", userID, commentID,
	)
	return err
}