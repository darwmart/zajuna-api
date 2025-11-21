package repository

import "zajunaApi/internal/models"

type CompetencyFrameworkRepositoryInterface interface {
	Create(competency *models.CompetencyFramework) (*models.CompetencyFramework, error)
	FindByID(id uint) (*models.CompetencyFramework, error)
}
