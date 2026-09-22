package models

import (
	"time"
)

type Post struct {
	ID         int
	UserID     int       `json:"userid"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created    time.Time `json:"created"`
	Categories []string  `json:"categories"`
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
