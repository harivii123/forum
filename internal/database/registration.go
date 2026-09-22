package database

import (
	"database/sql"
	"forum/internal/models"
)

func CreateUser(db *sql.DB, user models.User) (int, error) {
	result, err := db.Exec(`INSERT INTO user (username, email, password_hash) VALUES (?, ?, ?)`, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), nil
}
