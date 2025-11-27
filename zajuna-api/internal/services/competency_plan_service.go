package services

import (
	"fmt"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type CompetencyPlanService struct {
	repo                repository.CompetencyPlanRepositoryInterface
	sessionRepo         repository.SessionsRepositoryInterface
	roleCapabilityCheck RoleCapabilityServiceInterface
}

// NewCompetencyService crea un nuevo servicio de categorías
func NewCompetencyPlanService(repo repository.CompetencyPlanRepositoryInterface, sessionRepo repository.SessionsRepositoryInterface, roleCapabilityCheck RoleCapabilityServiceInterface) *CompetencyPlanService {
	return &CompetencyPlanService{repo: repo, sessionRepo: sessionRepo, roleCapabilityCheck: roleCapabilityCheck}
}

func (s *CompetencyPlanService) CreateCompetencyPlan(sid string, plan *models.CompetencyPlan) (*models.CompetencyPlan, error) {

	session, err := s.sessionRepo.FindBySID(sid)
	if err != nil {
		return nil, err
	}
	if plan.TemplateID != nil {
		return nil, fmt.Errorf("para crear un plan basado en una plantilla use api::create_plan_from_template()")
	} else if plan.Status == request.STATUS_COMPLETE {
		return nil, fmt.Errorf("no se puede crear un plan con estado 'completo'")
	}

	is_draft := plan.Status == request.STATUS_DRAFT || plan.Status == request.STATUS_WAITING_FOR_REVIEW || plan.Status == request.STATUS_IN_REVIEW
	if is_draft {
		capabilities := []string{
			"moodle/competency:planmanagedraft",
		}
		if session.UserID != plan.UserID {
			capabilities = append(capabilities, "moodle/competency:planmanageowndraft")
		}
		hasCap, err := s.roleCapabilityCheck.HasCapabilities(capabilities, session.UserID)
		if err != nil {
			return nil, err
		}
		if !hasCap {
			return nil, fmt.Errorf("no tiene permisos para crear un plan de competencia en estado borrador")
		}
	}
	capabilities := []string{
		"moodle/competency:planmanage",
	}
	if session.UserID != plan.UserID {
		capabilities = append(capabilities, "moodle/competency:planmanageown")
	}
	hasCap, err := s.roleCapabilityCheck.HasCapabilities(capabilities, session.UserID)
	if err != nil {
		return nil, err
	}
	if !hasCap {
		return nil, fmt.Errorf("no tiene permisos para crear un plan de competencia")
	}
	return s.repo.Create(plan)
}
