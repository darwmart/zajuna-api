package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository es un mock del UserRepository para testing
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByFilters(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	args := m.Called(filters, page, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]models.User), args.Get(1).(int64), args.Error(2)
}

func (m *MockUserRepository) DeleteUsers(userIDs []int) error {
	args := m.Called(userIDs)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUsers(users []models.User) (int64, error) {
	args := m.Called(users)
	return args.Get(0).(int64), args.Error(1)
}
