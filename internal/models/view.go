package models

import "time"

type PostView struct {
	ID            int
	Poster        UserView
	Title         string
	Body          string
	Image         string
	Created       time.Time
	PostLikers    []UserView
	PostDislikers []UserView
	Comments      []CommentView
	Categories    []PostCategory
}

type UserView struct {
	ID             int
	Username       string
	ProfilePicture string
}

type CommentView struct {
	ID               int
	Commenter        UserView
	Body             string
	Created          time.Time
	CommentLikers    []UserView
	CommentDislikers []UserView
}

type PostCategory string
