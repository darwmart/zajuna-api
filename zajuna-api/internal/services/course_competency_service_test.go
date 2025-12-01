package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAddCompetencyToCourse_AlreadyExists(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 1, CompetencyID: 2}

	mockRepo.
		On("Exists", uint(1), uint(2)).
		Return(true, nil)

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ya está vinculada")
}

func TestAddCompetencyToCourse_ExistsError(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	mockRepo.
		On("Exists", uint(1), uint(2)).
		Return(false, errors.New("db error"))

	err := service.AddCompetencyToCourse("SID", &models.CourseCompetency{CourseID: 1, CompetencyID: 2})

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}

func TestAddCompetencyToCourse_SessionError(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	mockRepo.
		On("Exists", uint(1), uint(2)).
		Return(false, nil)

	mockSession.
		On("FindBySID", "BAD").
		Return(nil, errors.New("session error"))

	err := service.AddCompetencyToCourse("BAD", &models.CourseCompetency{CourseID: 1, CompetencyID: 2})

	assert.Error(t, err)
	assert.Equal(t, "session error", err.Error())
}

func TestAddCompetencyToCourse_CompetencyNotFound(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 1, CompetencyID: 2}

	mockRepo.On("Exists", uint(1), uint(2)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)

	// Competencia no encontrada
	mockCompetency.On("FindByID", uint(2)).Return(nil, nil)

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no existe una competencia")
}

func TestAddCompetencyToCourse_CompetencyFindError(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	mockRepo.On("Exists", uint(1), uint(2)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)
	mockCompetency.On("FindByID", uint(2)).Return(nil, errors.New("competency error"))

	err := service.AddCompetencyToCourse("SID", &models.CourseCompetency{CourseID: 1, CompetencyID: 2})

	assert.Error(t, err)
	assert.Equal(t, "competency error", err.Error())
}

func TestAddCompetencyToCourse_FrameworkNotFound(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 1, CompetencyID: 2}

	mockRepo.On("Exists", uint(1), uint(2)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)

	mockCompetency.On("FindByID", uint(2)).
		Return(&models.Competency{ID: 2, CompetencyFrameworkID: 9}, nil)

	mockFramework.On("FindByID", uint(9)).Return(nil, nil)

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "framework con id 9 no existe")
}

func TestAddCompetencyToCourse_FrameworkFindError(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	mockRepo.On("Exists", uint(1), uint(2)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)
	mockCompetency.On("FindByID", uint(2)).
		Return(&models.Competency{ID: 2, CompetencyFrameworkID: 5}, nil)

	mockFramework.On("FindByID", uint(5)).
		Return(nil, errors.New("framework error"))

	err := service.AddCompetencyToCourse("SID", &models.CourseCompetency{CourseID: 1, CompetencyID: 2})

	assert.Error(t, err)
	assert.Equal(t, "framework error", err.Error())
}

func TestAddCompetencyToCourse_FrameworkHidden(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 1, CompetencyID: 2}

	mockRepo.On("Exists", uint(1), uint(2)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)

	mockCompetency.On("FindByID", uint(2)).
		Return(&models.Competency{ID: 2, CompetencyFrameworkID: 5}, nil)

	mockFramework.On("FindByID", uint(5)).
		Return(&models.CompetencyFramework{ID: 5, Visible: 0}, nil)

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "framework escondido")
}

func TestAddCompetencyToCourse_CourseNotFound(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 20, CompetencyID: 3}

	mockRepo.On("Exists", uint(20), uint(3)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)
	mockCompetency.On("FindByID", uint(3)).
		Return(&models.Competency{ID: 3, CompetencyFrameworkID: 7}, nil)
	mockFramework.On("FindByID", uint(7)).
		Return(&models.CompetencyFramework{ID: 7, Visible: 1}, nil)

	// Curso NO existe
	mockCourse.On("GetCourseByID", uint(20)).
		Return(nil, gorm.ErrRecordNotFound)

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no existe un curso con id 20")
}

func TestAddCompetencyToCourse_CourseQueryError(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 20, CompetencyID: 3}

	mockRepo.On("Exists", uint(20), uint(3)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)
	mockCompetency.On("FindByID", uint(3)).
		Return(&models.Competency{ID: 3, CompetencyFrameworkID: 7}, nil)
	mockFramework.On("FindByID", uint(7)).
		Return(&models.CompetencyFramework{ID: 7, Visible: 1}, nil)

	// Error inesperado
	mockCourse.On("GetCourseByID", uint(20)).
		Return(nil, errors.New("db error"))

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}

func TestAddCompetencyToCourse_AddError(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 20, CompetencyID: 3}

	mockRepo.On("Exists", uint(20), uint(3)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)
	mockCompetency.On("FindByID", uint(3)).
		Return(&models.Competency{ID: 3, CompetencyFrameworkID: 7}, nil)
	mockFramework.On("FindByID", uint(7)).
		Return(&models.CompetencyFramework{ID: 7, Visible: 1}, nil)
	mockCourse.On("GetCourseByID", uint(20)).
		Return(&models.Course{ID: 20}, nil)

	mockRepo.On("AddCompetencyToCourse", cc).Return(errors.New("insert error"))

	err := service.AddCompetencyToCourse("SID", cc)

	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())
}

func TestAddCompetencyToCourse_Success(t *testing.T) {
	mockRepo := new(mocks.MockCourseCompetencyRepository)
	mockCompetency := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)
	mockCourse := new(mocks.MockCourseRepository)

	service := NewCourseCompetencyService(mockRepo, mockCompetency, mockSession, mockFramework, mockCourse)

	cc := &models.CourseCompetency{CourseID: 20, CompetencyID: 3}

	mockRepo.On("Exists", uint(20), uint(3)).Return(false, nil)
	mockSession.On("FindBySID", "SID").Return(&models.Sessions{UserID: 10}, nil)

	mockCompetency.On("FindByID", uint(3)).
		Return(&models.Competency{ID: 3, CompetencyFrameworkID: 7}, nil)

	mockFramework.On("FindByID", uint(7)).
		Return(&models.CompetencyFramework{ID: 7, Visible: 1}, nil)

	mockCourse.On("GetCourseByID", uint(20)).
		Return(&models.Course{ID: 20}, nil)

	mockRepo.On("AddCompetencyToCourse", cc).Return(nil)

	err := service.AddCompetencyToCourse("SID", cc)

	assert.NoError(t, err)
}
