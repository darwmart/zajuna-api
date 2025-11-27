package services

import "zajunaApi/internal/models"

type CompetencyPlanServiceInterface interface {
	CreateCompetencyPlan(sid string, plan *models.CompetencyPlan) (*models.CompetencyPlan, error)
}
