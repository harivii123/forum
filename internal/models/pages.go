package models

type FrontPage struct {
	Posts []Post
	Categories []Category
	Session Session
}