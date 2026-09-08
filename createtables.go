package main

import (
	"database/sql"
	"log"
	"strings"
)

func CreateTables(db *sql.DB) string {

	var sb strings.Builder
	
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		profile_picture BLOB
		)`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\nuser")
	

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		comment_id INTEGER,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		image BLOB,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comment(id)
		)`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\npost")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS like (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		comment_id INTEGER,
		post_id INTEGER,
		like_value INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		)`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\nlike")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
		)`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\ncomment")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
		)`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\ncategory")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS cookies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		cookie_key TEXT NOT NULL UNIQUE,
		user_id INTEGER,
		created_at TEXT NOT NULL,
		expires_at TEXT NOT NULL,
		last_active TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(id)
		)`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\ncookies")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS category_post (
		category_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		PRIMARY KEY (category_id, post_id),
		FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}
	sb.WriteString("\ncategory_post")

	return sb.String()
}
