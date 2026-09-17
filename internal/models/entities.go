package models

type Post struct {
	ID         int
	UserID     int
	Title      string
	Body       string
	Created    time.Time
	Categories []string
}