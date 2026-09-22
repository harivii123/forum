package database

import (
	"database/sql"
	"time"
)

func GetUserIDWithToken(tx *sql.Tx, token string) (int, error) {
	var userID int

	err := tx.QueryRow(
		`SELECT user_id FROM session WHERE token = ? AND expires_at > ?`,
		token, time.Now()).Scan(&userID)
	if err != nil {
		return userID, err
	}

	return userID, nil
}
