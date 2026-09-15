package database

import (
	"database/sql"
	"fmt"
	"strings"
)

var schema = []struct {
	name string
	stmt string
}{
	{"user", `CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		profile_picture BLOB
	)`},

	{"post", `CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		image BLOB,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
	)`},

	{"comment", `CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
	)`},

	{"category", `CREATE TABLE IF NOT EXISTS category (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	)`},

	{"category_post", `CREATE TABLE IF NOT EXISTS category_post (
		category_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		PRIMARY KEY (category_id, post_id),
		FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
	)`},
	// 1=like, -1 = dislike
	{"post_vote", `CREATE TABLE IF NOT EXISTS post_vote (
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		vote INTEGER NOT NULL CHECK (vote IN (-1, 1)),
		PRIMARY KEY (user_id, post_id),
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
	)`},

	{"comment_vote", `CREATE TABLE IF NOT EXISTS comment_vote (
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		vote INTEGER NOT NULL CHECK (vote IN (-1, 1)),
		PRIMARY KEY (user_id, comment_id),
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE
	)`},

	{"session", `CREATE TABLE IF NOT EXISTS session (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		token TEXT NOT NULL UNIQUE,
		user_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL,
		expires_at DATETIME NOT NULL,
		last_active DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
	)`},

	{"idx_post_user",
		`CREATE INDEX IF NOT EXISTS idx_post_user ON post(user_id)`},

	{"idx_cp_post",
		`CREATE INDEX IF NOT EXISTS idx_cp_post
		ON category_post(post_id)`},
}

func CreateSchema(db *sql.DB) (string, error) {
	var sb strings.Builder
	for _, t := range schema {
		if _, err := db.Exec(t.stmt); err != nil {
			return sb.String(), fmt.Errorf("creating %s: %w", t.name, err)
		}
		sb.WriteString("\n" + t.name)
	}
	return sb.String(), nil
}
