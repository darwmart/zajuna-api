package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockCompetencyFrameworkRepository struct {
	mock.Mock
}

func (m *MockCompetencyFrameworkRepository) Create(cf *models.CompetencyFramework) (*models.CompetencyFramework, error) {
	args := m.Called(cf)

	if res, ok := args.Get(0).(*models.CompetencyFramework); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompetencyFrameworkRepository) FindByID(id uint) (*models.CompetencyFramework, error) {
	args := m.Called(id)

	if res, ok := args.Get(0).(*models.CompetencyFramework); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}
