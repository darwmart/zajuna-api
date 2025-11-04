package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockSessionsRepository struct {
	mock.Mock
}

func (m *MockSessionsRepository) FindBySID(sid string) (*models.Sessions, error) {
	args := m.Called(sid)
	if session, ok := args.Get(0).(*models.Sessions); ok {
		return session, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockSessionsRepository) InsertSession(session *models.Sessions) error {
	args := m.Called(session)
	return args.Error(0)
}

func (m *MockSessionsRepository) DeleteSession(sid string) error {
	args := m.Called(sid)
	return args.Error(0)
}
