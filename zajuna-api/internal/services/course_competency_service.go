package services

import (
	"errors"
	"fmt"
	"time"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CourseCompetencyService struct {
	repo           repository.CourseCompetencyRepositoryInterface
	competencyRepo repository.CompetencyRepositoryInterface
	sessionRepo    repository.SessionsRepositoryInterface
	frameworkRepo  repository.CompetencyFrameworkRepositoryInterface
	courseRepo     repository.CourseRepositoryInterface
}

// NewCategoryService crea un nuevo servicio de categorías
func NewCourseCompetencyService(repo repository.CourseCompetencyRepositoryInterface, competencyRepo repository.CompetencyRepositoryInterface, sessionRepo repository.SessionsRepositoryInterface, frameworkRepo repository.CompetencyFrameworkRepositoryInterface, courseRepo repository.CourseRepositoryInterface) *CourseCompetencyService {
	return &CourseCompetencyService{repo: repo, competencyRepo: competencyRepo, sessionRepo: sessionRepo, frameworkRepo: frameworkRepo, courseRepo: courseRepo}
}

func (s *CourseCompetencyService) AddCompetencyToCourse(sid string, courseCompetency *models.CourseCompetency) error {

	exists, err := s.repo.Exists(courseCompetency.CourseID, courseCompetency.CompetencyID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("esta competencia ya está vinculada a este curso")
	} else {
		session, err := s.sessionRepo.FindBySID(sid)
		if err != nil {
			return err
		}
		log.Info(session.UserID)
		courseCompetency.UserModified = session.UserID
		courseCompetency.TimeCreated = time.Now().Unix()  // timestamp actual (segundos)
		courseCompetency.TimeModified = time.Now().Unix() // timestamp actual (segundos)
		competency, err := s.competencyRepo.FindByID(courseCompetency.CompetencyID)
		if err != nil {
			return err
		}
		if competency == nil {
			return fmt.Errorf("no existe una competencia con id %d", courseCompetency.CompetencyID)
		}
		framework, err := s.frameworkRepo.FindByID(competency.CompetencyFrameworkID)
		if err != nil {
			return err
		}

		if framework == nil {
			return fmt.Errorf("el framework con id %d no existe", competency.CompetencyFrameworkID)
		}

		if framework.Visible == 0 {
			return fmt.Errorf("una competencia que pertenece a un framework escondido no puede ser vinculada a un curso")
		}

		_, err = s.courseRepo.GetCourseByID(courseCompetency.CourseID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("no existe un curso con id %d", courseCompetency.CourseID)
			}
			return err // otro error (DB caída, SQL incorrecto, etc.)
		}

		err = s.repo.AddCompetencyToCourse(courseCompetency)
		if err != nil {
			return err
		}
		return nil
	}

}
