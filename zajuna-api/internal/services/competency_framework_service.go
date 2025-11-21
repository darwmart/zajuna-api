package services

import (
	"time"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"

	log "github.com/sirupsen/logrus"
)

type CompetencyFrameworkService struct {
	repo        repository.CompetencyFrameworkRepositoryInterface
	sessionRepo repository.SessionsRepositoryInterface
}

// NewCategoryService crea un nuevo servicio de categorías
func NewCompetencyFrameworkService(repo repository.CompetencyFrameworkRepositoryInterface, sessionRepo repository.SessionsRepositoryInterface) *CompetencyFrameworkService {
	return &CompetencyFrameworkService{repo: repo, sessionRepo: sessionRepo}
}

func (s *CompetencyFrameworkService) CreateCompetencyFramework(sid string, competencyFramework *models.CompetencyFramework) (*models.CompetencyFramework, error) {

	session, err := s.sessionRepo.FindBySID(sid)
	if err != nil {
		return nil, err
	}
	log.Info(session.UserID)
	competencyFramework.UserModified = session.UserID
	competencyFramework.TimeCreated = time.Now().Unix()  // timestamp actual (segundos)
	competencyFramework.TimeModified = time.Now().Unix() // timestamp actual (segundos)

	competencyFramework, err = s.repo.Create(competencyFramework)

	if err != nil {
		return nil, err
	}

	return competencyFramework, nil
}
