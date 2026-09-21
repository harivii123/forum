package models

import (
	"time"
)

type Post struct {
	ID         int
	UserID     int
	Title      string
	Body       string
	Created    time.Time
	Categories []string
}

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
}

type Session struct {
	ID         int
	Token      string
	UserID     int
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LastActive time.Time
}
