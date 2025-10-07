package repository

import (
	"database/sql"
)

type Course struct {
	ID         int    `json:"id"`
	ShortName  string `json:"shortname"`
	FullName   string `json:"fullname"`
	CategoryID int    `json:"categoryid"`
	Summary    string `json:"summary"`
	Format     string `json:"format"`
}

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) GetCoursesByCategory(categoryID int) ([]Course, error) {
	query := `
		SELECT 
			id, 
			shortname, 
			fullname, 
			category, 
			COALESCE(summary, '') AS summary, 
			COALESCE(format, '') AS format
		FROM mdl_course
		WHERE category = $1
		ORDER BY fullname
	`

	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var c Course
		if err := rows.Scan(
			&c.ID,
			&c.ShortName,
			&c.FullName,
			&c.CategoryID,
			&c.Summary,
			&c.Format,
		); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}
