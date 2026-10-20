package database

import (
	"database/sql"
	"time"
	"context"
)

func GetUserIDWithToken(ctx context.Context, tx *sql.Tx, token string) (int, error) {
	var userID int

	err := tx.QueryRowContext(ctx,
		`SELECT user_id FROM session WHERE token = ? AND expires_at > ?`,
		token, time.Now()).Scan(&userID)
	if err != nil {
		return userID, err
	}

	return userID, nil
}
