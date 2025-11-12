package mocks

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"

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

func (m *MockUserRepository) ToggleUserStatus(userID uint) (int, error) {
	args := m.Called(userID)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockUserRepository) GetEnrolledUsers(courseID int, options map[string]interface{}) ([]repository.EnrolledUserDetail, int, error) {
	args := m.Called(courseID, options)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int), args.Error(2)
	}
	return args.Get(0).([]repository.EnrolledUserDetail), args.Get(1).(int), args.Error(2)
}

func (m *MockUserRepository) GetUserGroupsInCourse(userID, courseID int) ([]map[string]interface{}, error) {
	args := m.Called(userID, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockUserRepository) GetUserRolesInCourse(userID, courseID int) ([]map[string]interface{}, error) {
	args := m.Called(userID, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockUserRepository) GetUserCustomFields(userID int) ([]map[string]interface{}, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockUserRepository) GetUserPreferences(userID int) ([]map[string]interface{}, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockUserRepository) GetUserEnrolledCourses(userID int) ([]map[string]interface{}, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	args := m.Called(username)

	// Si el primer argumento es nil, devolvemos nil directamente
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.User), args.Error(1)
}
