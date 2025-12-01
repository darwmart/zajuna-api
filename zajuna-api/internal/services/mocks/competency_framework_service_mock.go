package mocks

import "zajunaApi/internal/models"

type MockCompetencyFrameworkService struct {
	// valores configurables para el retorno
	FrameworkToReturn *models.CompetencyFramework
	ErrorToReturn     error

	// rastreo de la llamada
	Called            bool
	CalledSID         string
	CalledWithPayload *models.CompetencyFramework
}

func (m *MockCompetencyFrameworkService) CreateCompetencyFramework(
	sid string,
	cf *models.CompetencyFramework,
) (*models.CompetencyFramework, error) {

	m.Called = true
	m.CalledSID = sid
	m.CalledWithPayload = cf

	return m.FrameworkToReturn, m.ErrorToReturn
}
