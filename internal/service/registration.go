package service

import (
	"database/sql"
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"net/mail"
	"strings"
)

func RegisterUser(db *sql.DB, username, email, password string) (*models.User, error) {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	// Validate required fields
	if username == "" || email == "" || password == "" {
		return nil, ErrMissingFields
	}

	//Validate email
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrInvalidEmail
	}
	// Check whether username already exists.
	_, err := database.GetUserByUsername(db, username)
	if err == nil {
		return nil, ErrUsernameTaken
	}

	// Check whether eemail already exists
	_, err = database.GetUserByEmail(db, email)
	if err == nil {
		return nil, ErrEmailTaken
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	//Hash password before storing
	passwordHash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}

	user.ID, err = database.CreateUser(db, user)
	if err != nil {
		// Recheck in case another request registered the same value.
		_, usernameErr := database.GetUserByUsername(db, username)
		if usernameErr == nil {
			return nil, ErrUsernameTaken
		}

		_, emailErr := database.GetUserByEmail(db, email)
		if emailErr == nil {
			return nil, ErrEmailTaken
		}

		return nil, err
	}
	return &user, nil
}
