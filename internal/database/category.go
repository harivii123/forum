package database

import (
	"context"
	"database/sql"
	"forum/internal/models"
)

func GetCategories(ctx context.Context, tx *sql.Tx) ([]models.CategoryView, error) {
	query := `SELECT * FROM category;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.CategoryView
	for rows.Next() {
		var category models.CategoryView
		if err = rows.Scan(&category.ID, &category.Category, &category.Type); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return categories, nil
}
