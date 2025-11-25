package services

import "zajunaApi/internal/models"

type CourseCompetencyServiceInterface interface {
	AddCompetencyToCourse(sid string, courseCompetency *models.CourseCompetency) error
}
