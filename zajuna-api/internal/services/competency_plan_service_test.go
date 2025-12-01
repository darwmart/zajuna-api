package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	serviceMocks "zajunaApi/internal/services/mocks"

	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompetencyPlan_SessionError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	mockSession.On("FindBySID", "SID123").Return(nil, errors.New("session error"))

	plan := &models.CompetencyPlan{}

	// Act
	result, err := service.CreateCompetencyPlan("SID123", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "session error", err.Error())
	mockSession.AssertExpectations(t)
}

func TestCreateCompetencyPlan_TemplateNotAllowed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	session := &models.Sessions{UserID: 10}
	mockSession.On("FindBySID", "SID123").Return(session, nil)

	templateID := uint(99)
	plan := &models.CompetencyPlan{TemplateID: &templateID}

	// Act
	result, err := service.CreateCompetencyPlan("SID123", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "para crear un plan basado en una plantilla use api::create_plan_from_template()", err.Error())
	mockSession.AssertExpectations(t)
}

func TestCreateCompetencyPlan_StatusCompleteNotAllowed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	session := &models.Sessions{UserID: 10}
	mockSession.On("FindBySID", "SID123").Return(session, nil)

	plan := &models.CompetencyPlan{
		Status: request.STATUS_COMPLETE,
	}

	// Act
	result, err := service.CreateCompetencyPlan("SID123", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "no se puede crear un plan con estado 'completo'", err.Error())
	mockSession.AssertExpectations(t)
}

func TestCreateCompetencyPlan_Draft_NoPermissions(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	session := &models.Sessions{UserID: 5}
	mockSession.On("FindBySID", "SID123").Return(session, nil)

	plan := &models.CompetencyPlan{
		UserID: 7,
		Status: request.STATUS_DRAFT,
	}

	mockRoles.
		On("HasCapabilities",
			[]string{
				"moodle/competency:planmanagedraft",
				"moodle/competency:planmanageowndraft",
			},
			uint(5)).
		Return(false, nil)

	// Act
	result, err := service.CreateCompetencyPlan("SID123", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "no tiene permisos para crear un plan de competencia en estado borrador", err.Error())

	mockSession.AssertExpectations(t)
	mockRoles.AssertExpectations(t)
}

func TestCreateCompetencyPlan_Draft_WithPermissions(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	session := &models.Sessions{UserID: 8}
	mockSession.On("FindBySID", "SID123").Return(session, nil)

	plan := &models.CompetencyPlan{
		UserID: 99,
		Status: request.STATUS_DRAFT,
	}

	// draft permissions → allowed
	mockRoles.
		On("HasCapabilities",
			[]string{
				"moodle/competency:planmanagedraft",
				"moodle/competency:planmanageowndraft",
			},
			uint(8)).
		Return(true, nil)

	// final permission check → allowed
	mockRoles.
		On("HasCapabilities",
			[]string{
				"moodle/competency:planmanage",
				"moodle/competency:planmanageown",
			},
			uint(8)).
		Return(true, nil)

	mockRepo.On("Create", plan).Return(plan, nil)

	// Act
	result, err := service.CreateCompetencyPlan("SID123", plan)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, plan, result)

	mockSession.AssertExpectations(t)
	mockRoles.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateCompetencyPlan_FinalPermissionDenied(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	session := &models.Sessions{UserID: 11}
	mockSession.On("FindBySID", "SID123").Return(session, nil)

	plan := &models.CompetencyPlan{
		UserID: 11,
		Status: request.STATUS_ACTIVE,
	}

	mockRoles.
		On("HasCapabilities",
			[]string{"moodle/competency:planmanage"},
			uint(11)).
		Return(false, nil)

	// Act
	result, err := service.CreateCompetencyPlan("SID123", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "no tiene permisos para crear un plan de competencia", err.Error())

	mockSession.AssertExpectations(t)
	mockRoles.AssertExpectations(t)
}

func TestNewCompetencyPlanService(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	// Act
	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
	assert.Equal(t, mockSession, service.sessionRepo)
	assert.Equal(t, mockRoles, service.roleCapabilityCheck)
}
func TestCreateCompetencyPlan_Draft_CapabilityError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	// usuario == plan.UserID -> solo se pide "planmanagedraft"
	session := &models.Sessions{UserID: 30}
	mockSession.On("FindBySID", "SID_ERR_DRAFT").Return(session, nil)

	plan := &models.CompetencyPlan{
		UserID: 30,
		Status: request.STATUS_DRAFT,
	}

	mockRoles.
		On("HasCapabilities",
			[]string{"moodle/competency:planmanagedraft"},
			uint(30),
		).
		Return(false, errors.New("cap draft error"))

	// Act
	result, err := service.CreateCompetencyPlan("SID_ERR_DRAFT", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "cap draft error")

	mockSession.AssertExpectations(t)
	mockRoles.AssertExpectations(t)
}

func TestCreateCompetencyPlan_FinalCapabilityError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCompetencyPlanRepository)
	mockSession := new(mocks.MockSessionsRepository)
	mockRoles := new(serviceMocks.MockRoleCapabilityService)

	service := NewCompetencyPlanService(mockRepo, mockSession, mockRoles)

	// usuario != plan.UserID -> se pide planmanage + planmanageown
	session := &models.Sessions{UserID: 20}
	mockSession.On("FindBySID", "SID_ERR_FINAL").Return(session, nil)

	plan := &models.CompetencyPlan{
		UserID: 10,
		Status: request.STATUS_ACTIVE,
	}

	expectedCaps := []string{"moodle/competency:planmanage", "moodle/competency:planmanageown"}

	mockRoles.
		On("HasCapabilities",
			expectedCaps,
			uint(20),
		).
		Return(false, errors.New("cap final error"))

	// Act
	result, err := service.CreateCompetencyPlan("SID_ERR_FINAL", plan)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "cap final error")

	mockSession.AssertExpectations(t)
	mockRoles.AssertExpectations(t)
}
