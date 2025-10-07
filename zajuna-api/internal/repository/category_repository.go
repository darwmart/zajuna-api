package repository

import (
	"database/sql"
	"zajunaApi/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	rows, err := r.db.Query(`
		SELECT id, name, idnumber, description, descriptionformat, parent, 
		       sortorder, coursecount, visible, depth, path, theme
		FROM mdl_course_categories
		ORDER BY sortorder
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.IDNumber,
			&c.Description,
			&c.DescriptionFormat,
			&c.Parent,
			&c.SortOrder,
			&c.CourseCount,
			&c.Visible,
			&c.Depth,
			&c.Path,
			&c.Theme, // 👈 ahora soporta NULL
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
