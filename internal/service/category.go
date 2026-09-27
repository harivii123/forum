package service

import (
	"context"
	"database/sql"
	"forum/internal/database"
	"forum/internal/models"
)

func GetCategories(ctx context.Context, db *sql.DB) ([]models.CategoryView, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	categories, err := database.GetCategories(ctx, tx)
	if err != nil {
		return nil, err
	}

	return categories, tx.Commit()
}
