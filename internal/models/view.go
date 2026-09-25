package models

import (
	"net/url"
	"strconv"
	"time"
)

type Filters struct {
	Search     string
	Posts      bool
	Likes      bool
	Order      string
	Categories map[int]bool
}

func (f *Filters) PageFilters(q url.Values) {
	f.Search = q.Get("search")
	f.Posts = q.Get("posts") == "Posts"
	f.Likes = q.Get("likes") == "Likes"
	f.Order = q.Get("order")
	f.Categories = map[int]bool{}

	for _, c := range q["category"] {
		i, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		f.Categories[i] = true
	}
}

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
	Categories    []CategoryView
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

type CategoryView struct {
	ID       int
	Category string
	Type     string
}
