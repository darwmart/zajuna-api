package repository

import "zajunaApi/internal/models"

type CompetencyPlanRepositoryInterface interface {
	Create(plan *models.CompetencyPlan) (*models.CompetencyPlan, error)
}
