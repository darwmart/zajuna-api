package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockCompetencyPlanRepository struct {
	mock.Mock
}

func (m *MockCompetencyPlanRepository) Create(plan *models.CompetencyPlan) (*models.CompetencyPlan, error) {
	args := m.Called(plan)

	if res, ok := args.Get(0).(*models.CompetencyPlan); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}
