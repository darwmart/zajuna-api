package services

import (
	"zajunaApi/internal/models"
)

type CompetencyServiceInterface interface {
	CreateCompetency(sid string, competency *models.Competency) (*models.Competency, error)
}
