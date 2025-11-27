package services

import (
	"zajunaApi/internal/repository"
)

const (
	CAP_INHERIT  = 0
	CAP_ALLOW    = 1
	CAP_PREVENT  = -1
	CAP_PROHIBIT = -1000
)

type RoleCapabilityService struct {
	repo       repository.RoleCapabilityRepositoryInterface
	configRepo repository.ConfigRepositoryInterface
}

func NewRoleCapabilityService(repo repository.RoleCapabilityRepositoryInterface) *RoleCapabilityService {
	return &RoleCapabilityService{repo: repo}
}

func (s *RoleCapabilityService) HasCapabilities(capabilities []string, userID uint) (bool, error) {

	roles := []string{}
	allowed := false

	// Obtener ids por defecto desde configs
	defaultUserRoleID, err := s.configRepo.FindByName("defaultuserroleid")
	if err != nil {
		return false, err
	}
	if defaultUserRoleID != nil {
		roles = append(roles, defaultUserRoleID.Value)
	}

	defaultFrontPageRoleID, err := s.configRepo.FindByName("defaultfrontpageroleid")
	if err != nil {
		return false, err
	}
	if defaultFrontPageRoleID != nil {
		roles = append(roles, defaultFrontPageRoleID.Value)
	}

	// Evaluar cada capability del arreglo
	for _, cap := range capabilities {

		allowed = false // reiniciar para cada capability

		capsFromRepo, err := s.repo.FindByUserID(int64(userID), roles, cap)
		if err != nil {
			return false, err
		}

		for _, c := range *capsFromRepo {
			switch c.Permission {
			case CAP_PROHIBIT:
				// si cualquier capability está prohibido return false inmediato
				return false, nil
			case CAP_ALLOW:
				allowed = true
			}
		}

		// si este capability específico no está permitido return false
		if !allowed {
			return false, nil
		}
	}

	// Si llegó aquí, TODOS los capabilities fueron permitidos
	return true, nil
}
