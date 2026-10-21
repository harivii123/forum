package service

import (
	"context"
	"database/sql"
	"forum/internal/database"
	"log"
	"os"
	"strings"
)

func DeleteUser(ctx context.Context, db *sql.DB, userID int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	images, err := database.GetUserPostImages(ctx, tx, userID)
	if err != nil {
		return err
	}

	if err = database.DeleteUser(ctx, tx, userID); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	for _, image := range images {
		err = os.Remove(strings.TrimPrefix(image, "/"))
		log.Println(err)
	}

	return nil
}
