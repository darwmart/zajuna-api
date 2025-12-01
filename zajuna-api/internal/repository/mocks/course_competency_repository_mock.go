package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockCourseCompetencyRepository struct {
	mock.Mock
}

func (m *MockCourseCompetencyRepository) AddCompetencyToCourse(cc *models.CourseCompetency) error {
	args := m.Called(cc)
	return args.Error(0)
}

func (m *MockCourseCompetencyRepository) Exists(courseID uint, competencyID uint) (bool, error) {
	args := m.Called(courseID, competencyID)
	return args.Bool(0), args.Error(1)
}
