package mocks

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"

	"github.com/stretchr/testify/mock"
)

// MockCourseRepository es un mock del CourseRepository para testing
type MockCourseRepository struct {
	mock.Mock
}

func (m *MockCourseRepository) GetAllCourses() ([]models.Course, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *MockCourseRepository) GetCoursesByCategory(categoryID uint) ([]models.Course, error) {
	args := m.Called(categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *MockCourseRepository) GetCourseByID(id uint) (*models.Course, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Course), args.Error(1)
}

func (m *MockCourseRepository) GetCourseDetails(courseID int) (*repository.CourseDetails, error) {
	args := m.Called(courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repository.CourseDetails), args.Error(1)
}

func (m *MockCourseRepository) DeleteCourses(courseIDs []int) ([]models.Warning, error) {
	args := m.Called(courseIDs)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Warning), args.Error(1)
}

func (m *MockCourseRepository) UpdateCourse(id int, updates map[string]interface{}) error {
	args := m.Called(id, updates)
	return args.Error(0)
}
