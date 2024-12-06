package database

import (
	"database/sql"

	"github.com/jtieri/habbgo/models"
)

type NavigatorRepo struct {
	database *sql.DB
}

// NewNavigatorRepo returns a new instance of NavigatorRepo.
func NewNavigatorRepo(db *sql.DB) *NavigatorRepo {
	return &NavigatorRepo{database: db}
}

// Categories retrieves the navigator categories found in database table room_categories and returns them as a slice of
// Category structs.
func (navRepo *NavigatorRepo) Categories() ([]models.Category, error) {
	rows, err := navRepo.database.Query("SELECT id, parent_id, is_node, name, is_public, is_trading, min_rank_access, min_rank_setflatcat FROM room_categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		err = rows.Scan(
			&cat.ID,
			&cat.ParentID,
			&cat.IsNode,
			&cat.Name,
			&cat.IsPublic,
			&cat.IsTrading,
			&cat.MinRankAccess,
			&cat.MinRankSetFlat,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, cat)
	}

	return categories, nil
}
