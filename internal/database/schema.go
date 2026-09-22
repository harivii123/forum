package database

import (
	"database/sql"
	"fmt"
)

func createTables(db *sql.DB) error {
	var schema = []struct {
		name string
		stmt string
	}{
		{"user", `
			CREATE TABLE IF NOT EXISTS user (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username TEXT NOT NULL UNIQUE,
				email TEXT NOT NULL UNIQUE,
				password_hash TEXT NOT NULL,
				profile_picture TEXT
			)
		`},

		{"user_fts", `
			CREATE VIRTUAL TABLE IF NOT EXISTS user_fts USING fts5(
				username,
				content='user',
				content_rowid='id',
				tokenize='trigram'
			)
		`},

		{"insert_trigger_user", `
			CREATE TRIGGER IF NOT EXISTS user_ai
			AFTER INSERT ON user
			BEGIN
				INSERT INTO user_fts(rowid, username)
				VALUES (new.id, new.username);
			END;
		`},

		{"delete_trigger_user", `
			CREATE TRIGGER IF NOT EXISTS user_ad
			AFTER DELETE ON user
			BEGIN
				INSERT INTO user_fts(user_fts, rowid, username)
				VALUES ('delete', old.id, old.username);
			END;
		`},

		{"update_trigger_user", `
			CREATE TRIGGER IF NOT EXISTS user_au
			AFTER UPDATE OF username
			ON user
			BEGIN
				INSERT INTO user_fts(user_fts, rowid, username)
				VALUES ('delete', old.id, old.username);
				INSERT INTO user_fts(rowid, username) VALUES (new.id, new.username);
			END;
		`},

		{"post", `
			CREATE TABLE IF NOT EXISTS post (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				title TEXT NOT NULL,
				body TEXT NOT NULL,
				image TEXT,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
			)
		`},

		{"post_fts", `
			CREATE VIRTUAL TABLE IF NOT EXISTS post_fts USING fts5(
				title,
				body,
				author,
				comments,
				commenters,
				categories,
				tokenize='trigram'
			)
		`},

		{"comment", `
			CREATE TABLE IF NOT EXISTS comment (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				post_id INTEGER NOT NULL,
				body TEXT NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
				FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
			)
		`},

		{"category", `
			CREATE TABLE IF NOT EXISTS category (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL, 
				type TEXT NOT NULL,
				UNIQUE (name, type)
			)
		`}, // deleted type TEXT NOT NULL,UNIQUE (name, type) due to seed data doesn't have type

		{"category_post", `
			CREATE TABLE IF NOT EXISTS category_post (
				category_id INTEGER NOT NULL,
				post_id INTEGER NOT NULL,
				PRIMARY KEY (category_id, post_id),
				FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE,
				FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
			)
		`},

		// vote: 1 = like, -1 = dislike
		{"post_vote", `
			CREATE TABLE IF NOT EXISTS post_vote (
				user_id INTEGER NOT NULL,
				post_id INTEGER NOT NULL,
				vote INTEGER NOT NULL CHECK (vote IN (-1, 1)),
				PRIMARY KEY (user_id, post_id),
				FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
				FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
			)
		`},

		{"comment_vote", `
			CREATE TABLE IF NOT EXISTS comment_vote (
				user_id INTEGER NOT NULL,
				comment_id INTEGER NOT NULL,
				vote INTEGER NOT NULL CHECK (vote IN (-1, 1)),
				PRIMARY KEY (user_id, comment_id),
				FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
				FOREIGN KEY (comment_id) REFERENCES comment(id) ON DELETE CASCADE
			)
		`},

		{"session", `
			CREATE TABLE IF NOT EXISTS session (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				token TEXT NOT NULL UNIQUE,
				user_id INTEGER NOT NULL,
				created_at DATETIME NOT NULL,
				expires_at DATETIME NOT NULL,
				last_active DATETIME NOT NULL,
				FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
			)
		`},

		{"insert_trigger_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS post_ai
			AFTER INSERT ON post
			BEGIN %s END;
		`, postFtsRefresh("NEW.id"))},

		{"update_trigger_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS post_au
			AFTER UPDATE OF title, body, user_id ON post
			BEGIN %s END;
		`, postFtsRefresh("OLD.id, NEW.id"))},

		{"delete_trigger_post", `
			CREATE TRIGGER IF NOT EXISTS post_ad
			AFTER DELETE ON post
			BEGIN
				DELETE FROM post_fts WHERE rowid = OLD.id;
			END;
		`},

		{"insert_trigger_comment_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS comment_post_ai
			AFTER INSERT ON comment
			BEGIN %s END;
		`, postFtsRefresh("NEW.post_id"))},

		{"update_trigger_comment_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS comment_post_au
			AFTER UPDATE ON comment
			BEGIN %s END;
		`, postFtsRefresh("OLD.post_id, NEW.post_id"))},

		{"delete_trigger_comment_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS comment_post_ad
			AFTER DELETE ON comment
			BEGIN %s END;
		`, postFtsRefresh("OLD.post_id"))},

		{"insert_trigger_category_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS category_post_ai
			AFTER INSERT ON category_post
			BEGIN %s END;
		`, postFtsRefresh("NEW.post_id"))},

		{"delete_trigger_category_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS category_post_ad
			AFTER DELETE ON category_post
			BEGIN %s END;
		`, postFtsRefresh("OLD.post_id"))},

		{"update_trigger_user_post", fmt.Sprintf(`
			CREATE TRIGGER IF NOT EXISTS user_post_au
			AFTER UPDATE OF username ON user
			BEGIN %s END;
		`, postFtsRefresh(`
				SELECT id FROM post WHERE user_id = NEW.id
				UNION
				SELECT post_id FROM comment WHERE user_id = NEW.id`))},

		{"idx_post_user", `
			CREATE INDEX IF NOT EXISTS idx_post_user
			ON post(user_id)
		`},

		{"idx_cp_post", `
			CREATE INDEX IF NOT EXISTS idx_cp_post
			ON category_post(post_id)
		`},
	}

	for _, t := range schema {
		if _, err := db.Exec(t.stmt); err != nil {
			return fmt.Errorf("creating %s: %w", t.name, err)
		}
		// sb.WriteString("\n" + t.name)
	}
	return nil
}

func postFtsRefresh(ids string) string {
	return fmt.Sprintf(`
		DELETE FROM post_fts WHERE rowid IN (%[1]s);
		INSERT INTO post_fts(rowid, title, body, author, comments, commenters, categories)
		SELECT
			p.id,
			p.title,
			p.body,
			u.username,
			(SELECT group_concat(c.body, ' ')
			   FROM comment c
			  WHERE c.post_id = p.id),
			(SELECT group_concat(DISTINCT cu.username)
			   FROM comment c JOIN user cu ON cu.id = c.user_id
			  WHERE c.post_id = p.id),
			(SELECT group_concat(cat.name, ' ')
			   FROM category_post cp JOIN category cat ON cat.id = cp.category_id
			  WHERE cp.post_id = p.id)
		FROM post p
		JOIN user u ON u.id = p.user_id
		WHERE p.id IN (%[1]s);`, ids)
}

func CreateSchemas(db *sql.DB) error {

	if err := createTables(db); err != nil {
		return err
	}

	return nil
}
