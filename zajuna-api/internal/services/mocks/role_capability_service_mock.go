package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockRoleCapabilityService struct {
	mock.Mock
}

func (m *MockRoleCapabilityService) HasCapabilities(capabilities []string, userID uint) (bool, error) {
	args := m.Called(capabilities, userID)

	return args.Bool(0), args.Error(1)
}
