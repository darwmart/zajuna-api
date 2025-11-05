package mocks

import "github.com/stretchr/testify/mock"

type MockAuthPlugin struct {
	mock.Mock
}

// PreventLocalPasswords simula el método PreventLocalPasswords.
func (m *MockAuthPlugin) PreventLocalPasswords() bool {
	args := m.Called()
	return args.Bool(0)
}

// Login simula el método Login(username, password string)
func (m *MockAuthPlugin) Login(username, password string) (bool, error) {
	args := m.Called(username, password)
	return args.Bool(0), args.Error(1)
}
