package database

import "database/sql"

func UpdateUserProfile(db *sql.DB, userID int, bio string) error {
	_, err := db.Exec(`UPDATE user SET bio = ? WHERE id = ? `, bio, userID)
	return err
}
