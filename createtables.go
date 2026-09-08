package main

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {

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

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		image BLOB
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS like (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		like_value INTEGER NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		body TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
		)`)
	if err != nil {
		log.Fatal(err)
	}

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

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS category_post (
		category_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		PRIMARY KEY (category_id, post_id),
		FOREIGN KEY (category_id) REFERENCE category(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCE post(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS user_like (
		user_id INTEGER NOT NULL,
		like_id INTEGER NOT NULL,
		PRIMARY KEY (user_id, like_id),
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (like_id) REFERENCES like(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS user_comment (
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		PRIMARY KEY (user_id, comment_id),
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS user_post (
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		PRIMARY KEY (user_id, post_id),
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS comment_like (
		comment_id INTEGER NOT NULL,
		like_id INTEGER NOT NULL,
		PRIMARY KEY (comment_id, like_id),
		FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE,
		FOREIGN KEY (like_id) REFERENCES like(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS post_comment (
		post_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		PRIMARY KEY (post_id, comment_id),
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS post_like (
		post_id INTEGER NOT NULL,
		like_id INTEGER NOT NULL,
		PRIMARY KEY (post_id, like_id),
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
		FOREIGN KEY (like_id) REFERENCES like(id) ON DELETE CASCADE
		);`)
	if err != nil {
		log.Fatal(err)
	}

}
