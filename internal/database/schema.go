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
	{"user", `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, profile_picture BLOB)`},
	{"post", `CREATE TABLE IF NOT EXISTS post (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL, title TEXT NOT NULL,	body TEXT NOT NULL,	image BLOB,	created  DATETIME DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE)`},
	{"like", `CREATE TABLE IF NOT EXISTS like (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL, comment_id INTEGER, post_id INTEGER, like_value INTEGER NOT NULL, FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE, FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE, FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE)`},
	{"comment", `CREATE TABLE IF NOT EXISTS comment (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL, post_id INTEGER NOT NULL, body TEXT NOT NULL, FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE)`},
	{"category", `CREATE TABLE IF NOT EXISTS category (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE)`},
	{"category_post", `CREATE TABLE IF NOT EXISTS category_post (category_id INTEGER NOT NULL, post_id INTEGER NOT NULL, PRIMARY KEY (category_id, post_id), FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE, FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE)`},
	{"cookies", `CREATE TABLE IF NOT EXISTS cookies (id INTEGER PRIMARY KEY AUTOINCREMENT, cookie_key TEXT NOT NULL UNIQUE, user_id INTEGER, created_at TEXT NOT NULL, expires_at TEXT NOT NULL, last_active TEXT NOT NULL, FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE)`},
	{"idx_post_user", `CREATE INDEX IF NOT EXISTS idx_post_user ON post(user_id)`},
	{"idx_cp_post", `CREATE INDEX IF NOT EXISTS idx_cp_post ON category_post(post_id)`},
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
