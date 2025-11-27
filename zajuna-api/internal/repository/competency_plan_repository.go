package repository

import (
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CompetencyPlanRepository struct {
	DB *gorm.DB
}

func NewCompetencyPlanRepository(db *gorm.DB) *CompetencyPlanRepository {
	return &CompetencyPlanRepository{DB: db}
}

func (r *CompetencyPlanRepository) Create(plan *models.CompetencyPlan) (*models.CompetencyPlan, error) {
	if err := r.DB.Table("mdl_competency_plan").Create(plan).Error; err != nil {
		return nil, err
	}
	return plan, nil
}
