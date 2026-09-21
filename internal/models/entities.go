package models

import "time"

type Post struct {
	ID         int
	UserID     int       `json:"userid"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created    time.Time `json:"created"`
	Categories []string  `json:"categories"`
}
