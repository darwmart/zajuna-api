package mocks

import (
	"net/http"
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockUserService es un mock del UserService para testing
type MockUserService struct {
	mock.Mock
}

// GetUsers mockea el método GetUsers
func (m *MockUserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	args := m.Called(filters, page, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]models.User), args.Get(1).(int64), args.Error(2)
}

// DeleteUsers mockea el método DeleteUsers
func (m *MockUserService) DeleteUsers(userIDs []int) error {
	args := m.Called(userIDs)
	return args.Error(0)
}

// UpdateUsers mockea el método UpdateUsers
func (m *MockUserService) UpdateUsers(users []models.User) (int64, error) {
	args := m.Called(users)
	return args.Get(0).(int64), args.Error(1)
}
func (m *MockUserService) Login(r *http.Request, username, password string) (string, error) {
	args := m.Called(r, username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) Logout(sid string) (string, error) {
	args := m.Called(sid)
	return args.String(0), args.Error(1)
}
