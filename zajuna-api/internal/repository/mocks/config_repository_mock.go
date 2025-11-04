package mocks

import (
	"zajunaApi/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockConfigRepository simula el comportamiento del repositorio de configuración.
type MockConfigRepository struct {
	mock.Mock
}

func (m *MockConfigRepository) FindByName(name string) (*models.Config, error) {
	args := m.Called(name)
	if config, ok := args.Get(0).(*models.Config); ok {
		return config, args.Error(1)
	}
	return nil, args.Error(1)
}
