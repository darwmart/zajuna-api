package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockRoleCapabilityRepository simula el comportamiento del repositorio de roles y capacidades.
type MockRoleCapabilityRepository struct {
	mock.Mock
}

func (m *MockRoleCapabilityRepository) FindByUserID(userID int64, roles []string, capability string) (*[]models.RoleCapability, error) {
	args := m.Called(userID, roles, capability)
	if rc, ok := args.Get(0).(*[]models.RoleCapability); ok {
		return rc, args.Error(1)
	}
	return nil, args.Error(1)
}
