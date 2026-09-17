package database

import (
	"database/sql"
	"time"
)

func GetPostsByCategory(db *sql.DB, categoryName string) ([]Post, error) {
	rows, err := db.Query(`
        SELECT p.id, p.user_id, p.title, p.body, p.created
        FROM post p
        JOIN category_post cp ON cp.post_id = p.id
        JOIN category c       ON c.id = cp.category_id
        WHERE c.name = ?
        ORDER BY p.created DESC`, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.Created); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}
