package services

import (
	"fmt"
	"strconv"
	"time"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"

	log "github.com/sirupsen/logrus"
)

type CompetencyService struct {
	repo          repository.CompetencyRepositoryInterface
	sessionRepo   repository.SessionsRepositoryInterface
	frameworkRepo repository.CompetencyFrameworkRepositoryInterface
	courseRepo    repository.CourseRepositoryInterface
}

// NewCategoryService crea un nuevo servicio de categorías
func NewCompetencyService(repo repository.CompetencyRepositoryInterface, sessionRepo repository.SessionsRepositoryInterface) *CompetencyService {
	return &CompetencyService{repo: repo, sessionRepo: sessionRepo}
}

func (s *CompetencyService) CreateCompetency(sid string, competency *models.Competency) (*models.Competency, error) {

	session, err := s.sessionRepo.FindBySID(sid)
	if err != nil {
		return nil, err
	}

	framework, err := s.frameworkRepo.FindByID(competency.CompetencyFrameworkID)
	if err != nil {
		return nil, err
	}
	if framework == nil {
		return nil, fmt.Errorf("competencyframeworkid no existe")
	}

	competency.UserModified = session.UserID
	competency.TimeCreated = time.Now().Unix()  // timestamp actual (segundos)
	competency.TimeModified = time.Now().Unix() // timestamp actual (segundos)

	competency, err = s.repo.Create(competency)
	if err != nil {
		return nil, err
	}
	log.Info("aaa  ", strconv.Itoa(competency.ID))
	// Actualizar el Path con el ID asignado
	if competency.ParentID == 0 {
		competency.Path = "/" + strconv.Itoa(competency.ID) + "/"
	} else {
		parentCompetency, err := s.repo.FindByID(competency.ParentID)
		if err != nil {
			return nil, err
		}
		competency.Path = parentCompetency.Path + strconv.Itoa(competency.ID) + "/"
	}

	competencyFinal, err := s.repo.Update(competency)
	if err != nil {
		return nil, err
	}

	return competencyFinal, nil
}

func (s *CompetencyService) AddCompetencyToCourse(idCompetency int, idCourse int) error {

	s.repo.FindByID(uint(idCompetency))

	return nil
}
