package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockCompetencyRepository struct {
	mock.Mock
}

func (m *MockCompetencyRepository) Create(c *models.Competency) (*models.Competency, error) {
	args := m.Called(c)

	if res, ok := args.Get(0).(*models.Competency); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompetencyRepository) Update(c *models.Competency) (*models.Competency, error) {
	args := m.Called(c)

	if res, ok := args.Get(0).(*models.Competency); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompetencyRepository) FindByID(id uint) (*models.Competency, error) {
	args := m.Called(id)

	if res, ok := args.Get(0).(*models.Competency); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}
