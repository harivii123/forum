package service

import (
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// CheckPassword checks if the provided password matches the hashed password
func CheckPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func AuthenticateUser(db *sql.DB, email, password string) (*models.User, error) {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return nil, ErrMissingFields
	}
	user, err := database.GetUserByEmail(db, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if err = CheckPassword(user.PasswordHash, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
