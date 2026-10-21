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

func GetSessionUserID(db *sql.DB, token string) (int, error) {
	var userID int
	err := db.QueryRow(`
		SELECT user_id
		FROM session
		WHERE token = ? AND julianday(expires_at) > julianday(?)
	`, token, time.Now()).Scan(&userID)

	return userID, err
}

func GetUserByID(db *sql.DB, userID int) (*models.User, error) {
	var user models.User
	err := db.QueryRow(`SELECT id, username, email, COALESCE(bio,''), COALESCE(profile_picture,'') FROM user WHERE id = ?`, userID).Scan(&user.ID, &user.Username, &user.Email, &user.Bio, &user.ProfilePicture)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	// TODO
	var user models.User

	err := db.QueryRow(`
        SELECT id, username, email, password_hash
        FROM user
        WHERE username = ?
    `, username).Scan(
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
