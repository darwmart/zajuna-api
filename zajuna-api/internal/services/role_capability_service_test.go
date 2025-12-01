package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestHasCapabilities_DefaultUserRoleID_Error(t *testing.T) {
	mockRepo := new(mocks.MockRoleCapabilityRepository)
	mockConfig := new(mocks.MockConfigRepository)

	service := NewRoleCapabilityService(mockRepo, mockConfig)
	service.configRepo = mockConfig

	mockConfig.
		On("FindByName", "defaultuserroleid").
		Return(nil, errors.New("config error"))

	ok, err := service.HasCapabilities([]string{"cap1"}, 5)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, "config error", err.Error())
}

func TestHasCapabilities_DefaultFrontPageRoleID_Error(t *testing.T) {
	mockRepo := new(mocks.MockRoleCapabilityRepository)
	mockConfig := new(mocks.MockConfigRepository)

	service := NewRoleCapabilityService(mockRepo, mockConfig)
	service.configRepo = mockConfig

	mockConfig.
		On("FindByName", "defaultuserroleid").
		Return(&models.Config{Value: "role1"}, nil)

	mockConfig.
		On("FindByName", "defaultfrontpageroleid").
		Return(nil, errors.New("frontpage error"))

	ok, err := service.HasCapabilities([]string{"cap1"}, 10)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, "frontpage error", err.Error())
}

func TestHasCapabilities_RepoError(t *testing.T) {
	mockRepo := new(mocks.MockRoleCapabilityRepository)
	mockConfig := new(mocks.MockConfigRepository)

	service := NewRoleCapabilityService(mockRepo, mockConfig)
	service.configRepo = mockConfig

	mockConfig.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "u"}, nil)
	mockConfig.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "f"}, nil)

	mockRepo.
		On("FindByUserID", int64(22), []string{"u", "f"}, "cap1").
		Return(nil, errors.New("repo error"))

	ok, err := service.HasCapabilities([]string{"cap1"}, 22)

	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, "repo error", err.Error())
}

func TestHasCapabilities_Prohibited(t *testing.T) {
	mockRepo := new(mocks.MockRoleCapabilityRepository)
	mockConfig := new(mocks.MockConfigRepository)

	service := NewRoleCapabilityService(mockRepo, mockConfig)
	service.configRepo = mockConfig

	mockConfig.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "u"}, nil)
	mockConfig.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "f"}, nil)

	caps := []models.RoleCapability{
		{Permission: CAP_PROHIBIT},
	}

	mockRepo.
		On("FindByUserID", int64(7), []string{"u", "f"}, "cap1").
		Return(&caps, nil)

	ok, err := service.HasCapabilities([]string{"cap1"}, 7)

	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestHasCapabilities_NotAllowed(t *testing.T) {
	mockRepo := new(mocks.MockRoleCapabilityRepository)
	mockConfig := new(mocks.MockConfigRepository)

	service := NewRoleCapabilityService(mockRepo, mockConfig)
	service.configRepo = mockConfig

	mockConfig.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "u"}, nil)
	mockConfig.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "f"}, nil)

	// Empty permissions → not allowed
	caps := []models.RoleCapability{}

	mockRepo.
		On("FindByUserID", int64(33), []string{"u", "f"}, "cap1").
		Return(&caps, nil)

	ok, err := service.HasCapabilities([]string{"cap1"}, 33)

	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestHasCapabilities_AllAllowed(t *testing.T) {
	mockRepo := new(mocks.MockRoleCapabilityRepository)
	mockConfig := new(mocks.MockConfigRepository)

	service := NewRoleCapabilityService(mockRepo, mockConfig)
	service.configRepo = mockConfig

	mockConfig.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "u"}, nil)
	mockConfig.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "f"}, nil)

	allow := []models.RoleCapability{
		{Permission: CAP_ALLOW},
	}

	mockRepo.
		On("FindByUserID", int64(50), []string{"u", "f"}, "cap1").
		Return(&allow, nil)

	mockRepo.
		On("FindByUserID", int64(50), []string{"u", "f"}, "cap2").
		Return(&allow, nil)

	ok, err := service.HasCapabilities([]string{"cap1", "cap2"}, 50)

	assert.True(t, ok)
	assert.NoError(t, err)
}
