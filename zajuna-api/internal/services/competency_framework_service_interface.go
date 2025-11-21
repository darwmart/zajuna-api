package services

import "zajunaApi/internal/models"

type CompetencyFrameworkServiceInterface interface {
	CreateCompetencyFramework(sid string, competencyFramework *models.CompetencyFramework) (*models.CompetencyFramework, error)
}
