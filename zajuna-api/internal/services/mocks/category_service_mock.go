package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockCategoryService es un mock del CategoryService para testing
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetCategories() ([]models.Category, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) MoveCategory(id uint, beforeid uint, parentid *uint) error {
	args := m.Called(id, beforeid, parentid)
	return args.Error(0)
}
