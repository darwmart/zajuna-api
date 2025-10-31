package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository es un mock del CategoryRepository para testing
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetAllCategories() ([]models.Category, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Category), args.Error(1)
}
