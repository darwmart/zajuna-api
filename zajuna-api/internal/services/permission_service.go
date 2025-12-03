package services

import (
	"fmt"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repo *repository.PermissionRepository) *PermissionService {
	return &PermissionService{repo: repo}
}

// HasCapability verifica si un usuario tiene una capacidad en un contexto (replica has_capability de Moodle)
func (s *PermissionService) HasCapability(userID int, capability string, contextID int) (bool, error) {
	permission, err := s.repo.GetUserCapabilityInContext(userID, capability, contextID)
	if err != nil {
		return false, err
	}

	// Solo ALLOW (1) significa que tiene el permiso
	return permission == models.CapAllow, nil
}

// HasCapabilityByCourse verifica si un usuario tiene una capacidad en un curso
func (s *PermissionService) HasCapabilityByCourse(userID int, capability string, courseID int) (bool, error) {
	// Obtener contexto del curso
	ctx, err := s.repo.GetContextByLevelAndInstance(models.ContextCourse, courseID)
	if err != nil {
		return false, err
	}

	return s.HasCapability(userID, capability, ctx.ID)
}

// HasCapabilityByModule verifica si un usuario tiene una capacidad en un módulo
func (s *PermissionService) HasCapabilityByModule(userID int, capability string, moduleID int) (bool, error) {
	ctx, err := s.repo.GetContextByLevelAndInstance(models.ContextModule, moduleID)
	if err != nil {
		return false, err
	}

	return s.HasCapability(userID, capability, ctx.ID)
}

// RequireCapability lanza un error si el usuario no tiene la capacidad (replica require_capability)
func (s *PermissionService) RequireCapability(userID int, capability string, contextID int) error {
	hasCapability, err := s.HasCapability(userID, capability, contextID)
	if err != nil {
		return err
	}

	if !hasCapability {
		return fmt.Errorf("required capability '%s' not satisfied for user %d in context %d", capability, userID, contextID)
	}

	return nil
}

// RequireCapabilityByCourse lanza un error si el usuario no tiene la capacidad en el curso
func (s *PermissionService) RequireCapabilityByCourse(userID int, capability string, courseID int) error {
	ctx, err := s.repo.GetContextByLevelAndInstance(models.ContextCourse, courseID)
	if err != nil {
		return err
	}

	return s.RequireCapability(userID, capability, ctx.ID)
}

// IsSiteAdmin verifica si un usuario es administrador del sitio
func (s *PermissionService) IsSiteAdmin(userID int) (bool, error) {
	return s.repo.IsSiteAdmin(userID)
}

// HasAnyCapability verifica si el usuario tiene al menos una de las capacidades
func (s *PermissionService) HasAnyCapability(userID int, capabilities []string, contextID int) (bool, error) {
	for _, cap := range capabilities {
		hasCapability, err := s.HasCapability(userID, cap, contextID)
		if err != nil {
			return false, err
		}
		if hasCapability {
			return true, nil
		}
	}
	return false, nil
}

// HasAllCapabilities verifica si el usuario tiene todas las capacidades
func (s *PermissionService) HasAllCapabilities(userID int, capabilities []string, contextID int) (bool, error) {
	for _, cap := range capabilities {
		hasCapability, err := s.HasCapability(userID, cap, contextID)
		if err != nil {
			return false, err
		}
		if !hasCapability {
			return false, nil
		}
	}
	return true, nil
}

// GetSystemContext obtiene el contexto del sistema
func (s *PermissionService) GetSystemContext() (*models.Context, error) {
	return s.repo.GetSystemContext()
}

// GetContext obtiene un contexto por ID
func (s *PermissionService) GetContext(contextID int) (*models.Context, error) {
	return s.repo.GetContextByID(contextID)
}

// GetContextByLevelAndInstance obtiene un contexto por nivel e instanceID
func (s *PermissionService) GetContextByLevelAndInstance(level int, instanceID int) (*models.Context, error) {
	return s.repo.GetContextByLevelAndInstance(level, instanceID)
}

// GetUserRoles obtiene los roles de un usuario
func (s *PermissionService) GetUserRoles(userID int) ([]models.RoleAssignment, error) {
	return s.repo.GetUserRoleAssignments(userID)
}

// GetUserRolesInContext obtiene los roles de un usuario en un contexto específico
func (s *PermissionService) GetUserRolesInContext(userID, contextID int) ([]models.RoleAssignment, error) {
	return s.repo.GetUserRoleAssignmentsInContext(userID, contextID)
}

// AssignRole asigna un rol a un usuario en un contexto
func (s *PermissionService) AssignRole(roleID, userID, contextID int, component string, itemID int) error {
	return s.repo.AssignRole(roleID, userID, contextID, component, itemID)
}

// UnassignRole elimina una asignación de rol
func (s *PermissionService) UnassignRole(roleID, userID, contextID int) error {
	return s.repo.UnassignRole(roleID, userID, contextID)
}

// GetRole obtiene información de un rol
func (s *PermissionService) GetRole(roleID int) (*models.MoodleRole, error) {
	return s.repo.GetRole(roleID)
}

// --- Funciones de conveniencia para casos de uso comunes ---

// CanViewCourse verifica si un usuario puede ver un curso
func (s *PermissionService) CanViewCourse(userID, courseID int) (bool, error) {
	return s.HasCapabilityByCourse(userID, "moodle/course:view", courseID)
}

// CanEditCourse verifica si un usuario puede editar un curso
func (s *PermissionService) CanEditCourse(userID, courseID int) (bool, error) {
	return s.HasCapabilityByCourse(userID, "moodle/course:update", courseID)
}

// CanManageCategories verifica si un usuario puede gestionar categorías
func (s *PermissionService) CanManageCategories(userID int) (bool, error) {
	systemCtx, err := s.GetSystemContext()
	if err != nil {
		return false, err
	}

	return s.HasCapability(userID, "moodle/category:manage", systemCtx.ID)
}

// CanBackupCourse verifica si un usuario puede hacer backup de un curso
func (s *PermissionService) CanBackupCourse(userID, courseID int) (bool, error) {
	return s.HasCapabilityByCourse(userID, "moodle/backup:backupcourse", courseID)
}

// CanEnrolUsers verifica si un usuario puede inscribir usuarios en un curso
func (s *PermissionService) CanEnrolUsers(userID, courseID int) (bool, error) {
	return s.HasCapabilityByCourse(userID, "enrol/manual:enrol", courseID)
}

// CanViewGrades verifica si un usuario puede ver calificaciones
func (s *PermissionService) CanViewGrades(userID, courseID int) (bool, error) {
	return s.HasCapabilityByCourse(userID, "moodle/grade:viewall", courseID)
}

// CanDeleteCourse verifica si un usuario puede eliminar un curso
func (s *PermissionService) CanDeleteCourse(userID, courseID int) (bool, error) {
	return s.HasCapabilityByCourse(userID, "moodle/course:delete", courseID)
}

// CanCreateCourse verifica si un usuario puede crear cursos
func (s *PermissionService) CanCreateCourse(userID int) (bool, error) {
	systemCtx, err := s.GetSystemContext()
	if err != nil {
		return false, err
	}

	return s.HasCapability(userID, "moodle/course:create", systemCtx.ID)
}
