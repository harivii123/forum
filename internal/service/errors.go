package service

import "errors"

var (
	ErrMissingFields      = errors.New("missing required fields")
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrUsernameTaken      = errors.New("username already taken")
	ErrEmailTaken         = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidSession     = errors.New("invalid or expired session")
)
