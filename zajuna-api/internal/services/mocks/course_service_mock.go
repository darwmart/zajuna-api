package mocks

import (
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"

	"github.com/stretchr/testify/mock"
)

// MockCourseService es un mock del CourseService para testing
type MockCourseService struct {
	mock.Mock
}

// GetAllCourses mockea el método GetAllCourses
func (m *MockCourseService) GetAllCourses() ([]models.Course, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

// GetCoursesByCategory mockea el método GetCoursesByCategory
func (m *MockCourseService) GetCoursesByCategory(categoryID uint) ([]models.Course, error) {
	args := m.Called(categoryID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

// GetCourseDetails mockea el método GetCourseDetails
func (m *MockCourseService) GetCourseDetails(courseID int) (*repository.CourseDetails, error) {
	args := m.Called(courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repository.CourseDetails), args.Error(1)
}

// DeleteCourses mockea el método DeleteCourses
func (m *MockCourseService) DeleteCourses(courseIDs []int) (*models.DeleteCoursesResponse, error) {
	args := m.Called(courseIDs)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.DeleteCoursesResponse), args.Error(1)
}

// UpdateCourses mockea el método UpdateCourses
func (m *MockCourseService) UpdateCourses(courses []request.UpdateCourseRequest) (*models.UpdateCoursesResponse, error) {
	args := m.Called(courses)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UpdateCoursesResponse), args.Error(1)
}
