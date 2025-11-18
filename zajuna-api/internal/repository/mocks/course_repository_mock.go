package mocks

import (
	"zajunaApi/internal/dto/request"
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

func (m *MockCourseRepository) GetCourseByIDNumber(idnumber string) (*models.Course, error) {
	args := m.Called(idnumber)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Course), args.Error(1)
}

func (m *MockCourseRepository) GetCourseDetails(idnumber string) (*repository.CourseDetails, error) {
	args := m.Called(idnumber)
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

func (m *MockCourseRepository) SearchCourses(criteriaName, criteriaValue string, page, perPage int) ([]models.Course, int64, error) {
	args := m.Called(criteriaName, criteriaValue, page, perPage)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]models.Course), args.Get(1).(int64), args.Error(2)
}

func (m *MockCourseRepository) UpdateCourseFormatOptions(courseID int, options []request.CourseFormatOption) error {
	args := m.Called(courseID, options)
	return args.Error(0)
}

func (m *MockCourseRepository) UpdateCourseCustomFields(courseID int, fields []request.CustomField) error {
	args := m.Called(courseID, fields)
	return args.Error(0)
}

func (m *MockCourseRepository) MoveCourse(id int, categoryID int, beforeID *int) error {
	args := m.Called(id, categoryID, beforeID)
	return args.Error(0)
}
