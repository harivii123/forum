package database

import (
	"database/sql"
	"forum/internal/models"
	"time"
)

func CreateSession(db *sql.DB, session models.Session) error {
	_, err := db.Exec(`INSERT INTO session (token, user_id, created_at, expires_at, last_active) VALUES (?, ?, ?, ?, ?)`, session.Token, session.UserID, session.CreatedAt, session.ExpiresAt, session.LastActive)
	return err
}

// update session activity if it's not expired
func UpdateSessionActivity(db *sql.DB, token string, lastActive time.Time, expiresAt time.Time) error {
	result, err := db.Exec(`UPDATE session SET last_active = ?, expires_at = ? WHERE token = ? AND expires_at > ? `,
		lastActive, expiresAt, token, lastActive)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func DeleteSession(db *sql.DB, token string) error {
	_, err := db.Exec(
		`DELETE FROM session WHERE token = ?`,
		token,
	)
	return err
}

func DeleteExpiredSessions(db *sql.DB, now time.Time) error {
	_, err := db.Exec(
		`DELETE FROM session WHERE expires_at <= ?`,
		now,
	)
	return err
}
