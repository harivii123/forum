package service

import (
	"database/sql"
	"forum/internal/database"
)
// takes value from main search bar and finds all matches using fts(fast text search)
func MainSearch(mainSearchValue string, db *sql.DB) (match []database.SearchResult, err error) {
	//if "" nothing happens
	if mainSearchValue == "" {
		// log.Println(4)
		return nil, nil
	}

	tx, err := db.Begin()
	if err != nil {
		// log.Println(5)
		return nil, err
	}
	defer tx.Rollback()

	match, err = database.FindAMatch(tx, mainSearchValue)
	if err != nil {
		// log.Println(6)
		return nil, nil
	}
	err = tx.Commit()
	if err != nil {
		// log.Println(7)
		return nil, err
	}

	// log.Println(match, "here")
	return match, nil
}
