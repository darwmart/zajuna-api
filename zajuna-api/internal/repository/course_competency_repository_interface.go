package repository

import "zajunaApi/internal/models"

type CourseCompetencyRepositoryInterface interface {
	AddCompetencyToCourse(courseCompetency *models.CourseCompetency) error
	Exists(courseID uint, competencyID uint) (bool, error)
}
