package database

import (
	"database/sql"
	"forum/internal/models"
	"time"
)

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var user models.User

	err := db.QueryRow(`
        SELECT id, username, email, password_hash
        FROM user
        WHERE email = ?
    `, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateSession(db *sql.DB, session models.Session) error {
	_, err := db.Exec(`INSERT INTO session (token, user_id, created_at, expires_at, last_active) VALUES (?, ?, ?, ?, ?)`, session.Token, session.UserID, session.CreatedAt, session.ExpiresAt, session.LastActive)
	return err
}

func GetSessionUserID(db *sql.DB, token string) (int, error) {
	var userID int
	err := db.QueryRow(`
		SELECT user_id
		FROM session
		WHERE token = ? AND expires_at > ?
	`, token, time.Now()).Scan(&userID)

	return userID, err
}

func GetUserByID(db *sql.DB, userID int) (*models.User, error) {
	var user models.User
	err := db.QueryRow(`SELECT id, username, email FROM user WHERE id = ?`, userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, err
}
func GetUserByUsername(
	db *sql.DB,
	username string,
) (*models.User, error) {
	// TODO
}

func UpdateSessionActivity(db *sql.DB, token string, lastActive time.Time, expiresAt time.Time) error {
	_, err := db.Exec(`UPDATE session SET last_active = ?, expires_at = ? WHERE token = ?`,
		lastActive,
		expiresAt,
		token,
	)
	return err
}
