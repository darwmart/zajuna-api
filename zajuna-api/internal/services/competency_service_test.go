package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompetency_Success_NoParent(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	session := &models.Sessions{UserID: 10}
	competency := &models.Competency{
		CompetencyFrameworkID: 1,
		ParentID:              0,
	}

	createdCompetency := &models.Competency{
		ID:                    5,
		ParentID:              0,
		CompetencyFrameworkID: 1,
	}

	mockSession.On("FindBySID", "SID_OK").Return(session, nil)
	mockFramework.On("FindByID", uint(1)).Return(&models.CompetencyFramework{ID: 1}, nil)
	mockRepo.On("Create", competency).Return(createdCompetency, nil)

	// Para ParentID = 0 NO se llama a FindByID
	updatedCompetency := &models.Competency{
		ID:       5,
		Path:     "/5/",
		ParentID: 0,
	}
	mockRepo.On("Update", createdCompetency).Return(updatedCompetency, nil)

	result, err := service.CreateCompetency("SID_OK", competency)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "/5/", result.Path)
	mockRepo.AssertExpectations(t)
	mockSession.AssertExpectations(t)
	mockFramework.AssertExpectations(t)
}

func TestCreateCompetency_Success_WithParent(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	session := &models.Sessions{UserID: 10}
	competency := &models.Competency{
		CompetencyFrameworkID: 1,
		ParentID:              2,
	}

	createdCompetency := &models.Competency{
		ID:                    8,
		ParentID:              2,
		CompetencyFrameworkID: 1,
	}

	parentCompetency := &models.Competency{
		ID:   2,
		Path: "/1/2/",
	}

	mockSession.On("FindBySID", "SID_PARENT").Return(session, nil)
	mockFramework.On("FindByID", uint(1)).Return(&models.CompetencyFramework{ID: 1}, nil)
	mockRepo.On("Create", competency).Return(createdCompetency, nil)
	mockRepo.On("FindByID", uint(2)).Return(parentCompetency, nil)

	updatedCompetency := &models.Competency{
		ID:   8,
		Path: "/1/2/8/",
	}
	mockRepo.On("Update", createdCompetency).Return(updatedCompetency, nil)

	result, err := service.CreateCompetency("SID_PARENT", competency)

	assert.NoError(t, err)
	assert.Equal(t, "/1/2/8/", result.Path)
}

func TestCreateCompetency_SessionError(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	mockSession.On("FindBySID", "BAD").Return(nil, errors.New("session error"))

	result, err := service.CreateCompetency("BAD", &models.Competency{})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateCompetency_FrameworkError(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}
	session := &models.Sessions{UserID: 7}

	mockSession.On("FindBySID", "SID").Return(session, nil)
	mockFramework.On("FindByID", uint(33)).Return(nil, errors.New("framework error"))

	competency := &models.Competency{CompetencyFrameworkID: 33}

	result, err := service.CreateCompetency("SID", competency)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateCompetency_FrameworkNotFound(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	session := &models.Sessions{UserID: 7}

	mockSession.On("FindBySID", "SID").Return(session, nil)
	mockFramework.On("FindByID", uint(33)).Return(nil, nil)

	competency := &models.Competency{CompetencyFrameworkID: 33}

	result, err := service.CreateCompetency("SID", competency)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "competencyframeworkid no existe")
}

func TestCreateCompetency_CreateError(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	session := &models.Sessions{UserID: 12}
	competency := &models.Competency{CompetencyFrameworkID: 1}

	mockSession.On("FindBySID", "X").Return(session, nil)
	mockFramework.On("FindByID", uint(1)).Return(&models.CompetencyFramework{ID: 1}, nil)
	mockRepo.On("Create", competency).Return(nil, errors.New("create error"))

	result, err := service.CreateCompetency("X", competency)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateCompetency_FindParentError(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	session := &models.Sessions{UserID: 10}

	competency := &models.Competency{
		CompetencyFrameworkID: 1,
		ParentID:              99,
	}

	createdCompetency := &models.Competency{
		ID:                    50,
		ParentID:              99,
		CompetencyFrameworkID: 1,
	}

	mockSession.On("FindBySID", "SID_ERR").Return(session, nil)
	mockFramework.On("FindByID", uint(1)).Return(&models.CompetencyFramework{ID: 1}, nil)
	mockRepo.On("Create", competency).Return(createdCompetency, nil)
	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("parent error"))

	result, err := service.CreateCompetency("SID_ERR", competency)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateCompetency_UpdateError(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := &CompetencyService{
		repo:          mockRepo,
		sessionRepo:   mockSession,
		frameworkRepo: mockFramework,
	}

	session := &models.Sessions{UserID: 10}

	competency := &models.Competency{
		CompetencyFrameworkID: 1,
		ParentID:              0,
	}

	createdCompetency := &models.Competency{
		ID:                    3,
		ParentID:              0,
		CompetencyFrameworkID: 1,
	}

	mockSession.On("FindBySID", "SID_UPD").Return(session, nil)
	mockFramework.On("FindByID", uint(1)).Return(&models.CompetencyFramework{ID: 1}, nil)
	mockRepo.On("Create", competency).Return(createdCompetency, nil)
	mockRepo.On("Update", createdCompetency).Return(nil, errors.New("update error"))

	result, err := service.CreateCompetency("SID_UPD", competency)

	assert.Error(t, err)
	assert.Nil(t, result)
}
func TestNewCompetencyService(t *testing.T) {
	mockRepo := new(mocks.MockCompetencyRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockFramework := new(mocks.MockCompetencyFrameworkRepository)

	service := NewCompetencyService(mockRepo, mockSession, mockFramework)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
	assert.Equal(t, mockSession, service.sessionRepo)
	assert.Equal(t, mockFramework, service.frameworkRepo)
}
