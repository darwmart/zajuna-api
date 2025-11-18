package repository

import "zajunaApi/internal/models"

type CompetencyRepositoryInterface interface {
	Create(competency *models.Competency) (*models.Competency, error)
	Update(competency *models.Competency) (*models.Competency, error)
	FindByID(id uint) (*models.Competency, error)
}
