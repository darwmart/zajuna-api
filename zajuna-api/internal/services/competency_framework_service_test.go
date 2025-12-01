package services

import (
	"errors"
	"testing"
	"time"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompetencyFramework_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyFrameworkRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	service := NewCompetencyFrameworkService(mockRepo, mockSessionRepo)

	session := &models.Sessions{UserID: 55}
	input := &models.CompetencyFramework{IDNumber: "CF01"}

	mockSessionRepo.On("FindBySID", "SID123").Return(session, nil)
	mockRepo.On("Create", input).Return(input, nil)

	// Act
	result, err := service.CreateCompetencyFramework("SID123", input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, uint(55), result.UserModified)
	assert.NotZero(t, result.TimeCreated)
	assert.NotZero(t, result.TimeModified)

	// Validamos que los timestamps sean recientes
	now := time.Now().Unix()
	assert.LessOrEqual(t, result.TimeCreated, now)
	assert.LessOrEqual(t, result.TimeModified, now)

	mockRepo.AssertExpectations(t)
	mockSessionRepo.AssertExpectations(t)
}

func TestCreateCompetencyFramework_SessionError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyFrameworkRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	service := NewCompetencyFrameworkService(mockRepo, mockSessionRepo)

	expectedErr := errors.New("session error")
	mockSessionRepo.On("FindBySID", "SID123").Return(nil, expectedErr)

	// Act
	result, err := service.CreateCompetencyFramework("SID123", &models.CompetencyFramework{})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockSessionRepo.AssertExpectations(t)
}

func TestCreateCompetencyFramework_CreateError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyFrameworkRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	service := NewCompetencyFrameworkService(mockRepo, mockSessionRepo)

	session := &models.Sessions{UserID: 22}
	input := &models.CompetencyFramework{IDNumber: "CF02"}

	mockSessionRepo.On("FindBySID", "SID123").Return(session, nil)
	mockRepo.On("Create", input).Return(nil, errors.New("insert error"))

	// Act
	result, err := service.CreateCompetencyFramework("SID123", input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "insert error")

	mockRepo.AssertExpectations(t)
	mockSessionRepo.AssertExpectations(t)
}

func TestNewCompetencyFrameworkService(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyFrameworkRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)

	// Act
	service := NewCompetencyFrameworkService(mockRepo, mockSessionRepo)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
	assert.Equal(t, mockSessionRepo, service.sessionRepo)
}
