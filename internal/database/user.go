package database

import (
	"context"
	"database/sql"
	"time"
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

func GetUserPostImages(ctx context.Context, tx *sql.Tx, userID int) ([]string, error) {
	var image string
	var images []string
	rows, err := tx.QueryContext(ctx, "SELECT image FROM post WHERE user_id = ? AND image IS NOT NULL", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&image); err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, rows.Err()
}

func DeleteUser(ctx context.Context, tx *sql.Tx, userID int) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM user WHERE id=?", userID)
	return err
}
