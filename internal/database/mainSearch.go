package database

import (
	"database/sql"
)

type SearchResult struct {
	ID          int
	Type        string //tells us what came up(user, comment, post, etc...)
	Match       string //what matched the actual searchValue
	Description string //for post it has some actual data but for user and comment not used
}

func FindAMatch(tx *sql.Tx, searchValue string) ([]SearchResult, error) {
	// log.Println(searchValue, "rep")
	var sr []SearchResult

	rows, err := tx.Query(
		`SELECT u.id, 'user', u.username, 'User Account'
		 FROM user u
		 JOIN user_fts fts ON u.id = fts.rowid
		 WHERE user_fts MATCH ?

		UNION ALL

		SELECT p.id, 'post', p.title, p.body
		FROM post p
		JOIN post_fts fts ON p.id = fts.rowid
		WHERE post_fts MATCH ?

		UNION ALL

		SELECT c.id, 'comment', c.body, 'A Comment'
		FROM comment c
		JOIN comment_fts fts ON c.id = fts.rowid
		WHERE comment_fts MATCH ?`, searchValue, searchValue, searchValue) // searchValue 3 times for all 3 tables.
	if err != nil {
		// log.Println(1, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var oneSR SearchResult

		err := rows.Scan(&oneSR.ID, &oneSR.Type, &oneSR.Match, &oneSR.Description)
		if err != nil {
			// log.Println(2)
			return nil, err
		}
		sr = append(sr, oneSR)
	}
	if rows.Err() != nil {
		// log.Println(3)
		return nil, err
	}

	return sr, nil
}
