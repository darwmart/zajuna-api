package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

//////////////////////////////////////////////////////
// 1. OBTENER CONTEXTOS Y HERENCIA (path)
//////////////////////////////////////////////////////

// GetContextAndParents devuelve el contexto y todos sus contextos padres usando el campo PATH de mdl_context
// Replica el comportamiento de Moodle que evalúa permisos en toda la jerarquía
func (r *PermissionRepository) GetContextAndParents(contextID int) ([]int, error) {
	var ctx models.Context

	err := r.db.Table("mdl_context").
		Where("id = ?", contextID).
		First(&ctx).Error
	if err != nil {
		return nil, err
	}

	if ctx.Path == "" {
		return []int{contextID}, nil
	}

	// El path tiene formato "/1/3/57/" - extraemos los IDs
	parts := strings.Split(strings.Trim(ctx.Path, "/"), "/")

	result := make([]int, 0, len(parts))
	for _, p := range parts {
		id, convErr := strconv.Atoi(p)
		if convErr == nil {
			result = append(result, id)
		}
	}

	log.Printf("GetContextAndParents: contextID=%d, path=%s, parents=%v", contextID, ctx.Path, result)
	return result, nil
}

//////////////////////////////////////////////////////
// 2. OBTENER ROLES EN VARIOS CONTEXTOS
//////////////////////////////////////////////////////

// GetUserRolesInContexts obtiene todos los roles asignados al usuario en los contextos especificados
// Esto permite evaluar permisos heredados de contextos padres
func (r *PermissionRepository) GetUserRolesInContexts(userID int, contextIDs []int) ([]models.RoleAssignment, error) {
	if len(contextIDs) == 0 {
		return []models.RoleAssignment{}, nil
	}

	var roles []models.RoleAssignment

	err := r.db.Table("mdl_role_assignments").
		Where("userid = ?", userID).
		Where("contextid IN ?", contextIDs).
		Find(&roles).Error

	log.Printf("GetUserRolesInContexts: userID=%d, contexts=%v, found %d roles", userID, contextIDs, len(roles))
	return roles, err
}

//////////////////////////////////////////////////////
// 3. OBTENER CAPACIDAD DE UN ROL
//////////////////////////////////////////////////////

// GetRoleCapability obtiene el permiso de un rol específico para una capability
// Retorna: CAP_INHERIT (0), CAP_ALLOW (1), CAP_PREVENT (-1000), CAP_PROHIBIT (-1)
func (r *PermissionRepository) GetRoleCapability(roleID int, capability string) (int, error) {
	var rc models.RoleCapability

	err := r.db.Table("mdl_role_capabilities").
		Where("roleid = ? AND capability = ?", roleID, capability).
		First(&rc).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Moodle: si no existe registro, es CAP_INHERIT = 0
		return models.CapInherit, nil
	}

	if err != nil {
		return 0, err
	}

	log.Printf("GetRoleCapability: roleID=%d, capability=%s, permission=%d", roleID, capability, rc.Permission)
	return rc.Permission, nil
}

//////////////////////////////////////////////////////
// 4. VER SI ES SITE ADMIN
//////////////////////////////////////////////////////

// IsSiteAdmin verifica si un usuario es administrador del sitio
// Replica exactamente el comportamiento de is_siteadmin() de Moodle (accesslib.php:702)
func (r *PermissionRepository) IsSiteAdmin(userID int) (bool, error) {
	// En Moodle, los site admins están listados en la tabla mdl_config
	// con name='siteadmins' y value es una lista de IDs separados por comas
	// Ejemplo: "2,3,5"

	var siteAdminsConfig struct {
		Value string `gorm:"column:value"`
	}

	err := r.db.
		Table("mdl_config").
		Select("value").
		Where("name = ?", "siteadmins").
		First(&siteAdminsConfig).Error

	if err != nil {
		log.Printf("WARNING: No se encontró configuración 'siteadmins' en mdl_config, usando método alternativo por rol\n")
		// Si no existe la configuración, verificar usando el método alternativo
		// (usuarios con rol manager en contexto sistema)
		return r.isSiteAdminByRole(userID)
	}

	log.Printf("IsSiteAdmin: UserID=%d, siteadmins config='%s'\n", userID, siteAdminsConfig.Value)

	// Convertir la lista de IDs a un array
	if siteAdminsConfig.Value == "" {
		log.Printf("WARNING: Lista de siteadmins está vacía, usando método alternativo por rol\n")
		return r.isSiteAdminByRole(userID)
	}

	// Verificar si el userID está en la lista
	// El formato es: "2,3,5" donde cada número es un user ID
	adminIDs := strings.Split(siteAdminsConfig.Value, ",")
	userIDStr := fmt.Sprintf("%d", userID)

	for _, adminID := range adminIDs {
		if strings.TrimSpace(adminID) == userIDStr {
			log.Printf("Usuario %d ES site admin (encontrado en mdl_config.siteadmins)\n", userID)
			return true, nil
		}
	}

	log.Printf("Usuario %d NO es site admin (no está en la lista: %s)\n", userID, siteAdminsConfig.Value)
	return false, nil
}

// isSiteAdminByRole es un método alternativo para verificar admin
// Verifica si el usuario tiene el rol con archetype 'manager' en el contexto del sistema
func (r *PermissionRepository) isSiteAdminByRole(userID int) (bool, error) {
	var count int64
	err := r.db.
		Table("mdl_role_assignments ra").
		Joins("JOIN mdl_role r ON ra.roleid = r.id").
		Joins("JOIN mdl_context c ON ra.contextid = c.id").
		Where("ra.userid = ? AND r.archetype = ? AND c.contextlevel = ?", userID, "manager", models.ContextSystem).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		log.Printf("Usuario %d ES site admin (tiene rol manager en contexto sistema)\n", userID)
	} else {
		log.Printf("Usuario %d NO es site admin (no tiene rol manager)\n", userID)
	}

	return count > 0, nil
}

//////////////////////////////////////////////////////
// 5. HAS CAPABILITY - IMPLEMENTACIÓN COMO MOODLE
//////////////////////////////////////////////////////

// GetUserCapabilityInContext calcula si un usuario tiene una capacidad en un contexto
// Esta función replica EXACTAMENTE has_capability de Moodle (accesslib.php:432)
//
// Algoritmo de Moodle:
// 1. Si es site admin → ALLOW
// 2. Obtener contextos padres (herencia)
// 3. Obtener roles del usuario en esos contextos
// 4. Evaluar capacidades:
//    - Si alguna es PROHIBIT → DENY (bloquea todo)
//    - Si alguna es ALLOW → ALLOW
//    - Si ninguna es ALLOW → DENY
func (r *PermissionRepository) GetUserCapabilityInContext(userID int, capability string, contextID int) (int, error) {
	log.Printf("\n=== GetUserCapabilityInContext ===")
	log.Printf("UserID=%d, Capability=%s, ContextID=%d", userID, capability, contextID)

	// 1. Site admin SIEMPRE tiene todos los permisos
	isAdmin, err := r.IsSiteAdmin(userID)
	if err == nil && isAdmin {
		log.Printf("Usuario %d es site admin -> CAP_ALLOW", userID)
		return models.CapAllow, nil
	}

	// 2. Obtener contextos padres (herencia)
	contextIDs, err := r.GetContextAndParents(contextID)
	if err != nil {
		log.Printf("ERROR: Error obteniendo contextos padres: %v", err)
		return models.CapInherit, err
	}

	allow := false
	hasProhibit := false

	// 3. Obtener roles del usuario en esos contextos
	roles, err := r.GetUserRolesInContexts(userID, contextIDs)
	if err != nil {
		log.Printf("ERROR: Error obteniendo roles: %v", err)
		return models.CapInherit, err
	}

	if len(roles) == 0 {
		log.Printf("WARNING: Usuario %d no tiene roles en los contextos %v", userID, contextIDs)
		return models.CapInherit, nil
	}

	// 4. Evaluar capacidades igual que Moodle
	for _, role := range roles {
		perm, err := r.GetRoleCapability(role.RoleID, capability)
		if err != nil {
			log.Printf("WARNING: Error obteniendo capability para roleID=%d: %v", role.RoleID, err)
			continue
		}

		log.Printf("  RoleID=%d, Permission=%d", role.RoleID, perm)

		switch perm {
		case models.CapProhibit:
			// REGLA MOODLE: si hay prohibición, bloquea todo
			hasProhibit = true
			log.Printf("  CAP_PROHIBIT encontrado -> bloqueara el permiso")

		case models.CapAllow:
			allow = true
			log.Printf("  CAP_ALLOW encontrado")

		case models.CapPrevent:
			// no hace nada, sólo evita permitir
			log.Printf("  CAP_PREVENT encontrado")
		}
	}

	// Aplicar regla de prohibición
	if hasProhibit {
		log.Printf("Resultado final: CAP_PROHIBIT (bloqueado)")
		return models.CapProhibit, nil
	}

	if allow {
		log.Printf("Resultado final: CAP_ALLOW")
		return models.CapAllow, nil
	}

	log.Printf("Resultado final: CAP_INHERIT (no permitido)")
	return models.CapInherit, nil
}

//////////////////////////////////////////////////////
// 6. OBTENER CONTEXTOS
//////////////////////////////////////////////////////

func (r *PermissionRepository) GetSystemContext() (*models.Context, error) {
	var ctx models.Context
	err := r.db.Table("mdl_context").
		Where("contextlevel = ? AND instanceid = ?", models.ContextSystem, 0).
		First(&ctx).Error
	return &ctx, err
}

func (r *PermissionRepository) GetContextByID(contextID int) (*models.Context, error) {
	var ctx models.Context
	err := r.db.Table("mdl_context").
		Where("id = ?", contextID).
		First(&ctx).Error
	return &ctx, err
}

func (r *PermissionRepository) GetContextByLevelAndInstance(level int, instanceID int) (*models.Context, error) {
	var ctx models.Context
	err := r.db.Table("mdl_context").
		Where("contextlevel = ? AND instanceid = ?", level, instanceID).
		First(&ctx).Error
	return &ctx, err
}

//////////////////////////////////////////////////////
// 7. ASIGNAR / QUITAR ROLES
//////////////////////////////////////////////////////

func (r *PermissionRepository) AssignRole(roleID, userID, contextID int, component string, itemID int) error {
	return r.db.Exec(`
		INSERT INTO mdl_role_assignments (roleid, userid, contextid, component, itemid, timemodified, modifierid)
		VALUES (?, ?, ?, ?, ?, UNIX_TIMESTAMP(), 2)
	`, roleID, userID, contextID, component, itemID).Error
}

func (r *PermissionRepository) UnassignRole(roleID, userID, contextID int) error {
	return r.db.Exec(`
		DELETE FROM mdl_role_assignments
		WHERE roleid = ? AND userid = ? AND contextid = ?
	`, roleID, userID, contextID).Error
}

func (r *PermissionRepository) GetUserRoleAssignments(userID int) ([]models.RoleAssignment, error) {
	var assignments []models.RoleAssignment
	err := r.db.Table("mdl_role_assignments").
		Where("userid = ?", userID).
		Find(&assignments).Error
	return assignments, err
}

func (r *PermissionRepository) GetUserRoleAssignmentsInContext(userID, contextID int) ([]models.RoleAssignment, error) {
	// Obtener el contexto para su path
	ctx, err := r.GetContextByID(contextID)
	if err != nil {
		return nil, err
	}

	// Buscar asignaciones en este contexto y en sus contextos padre (usando el path)
	var assignments []models.RoleAssignment
	err = r.db.
		Table("mdl_role_assignments ra").
		Joins("JOIN mdl_context c ON ra.contextid = c.id").
		Where("ra.userid = ?", userID).
		Where("? LIKE CONCAT(c.path, '%')", ctx.Path).
		Find(&assignments).Error

	return assignments, err
}

//////////////////////////////////////////////////////
// 8. ROLES
//////////////////////////////////////////////////////

func (r *PermissionRepository) GetRole(roleID int) (*models.MoodleRole, error) {
	var role models.MoodleRole
	err := r.db.Table("mdl_role").
		Where("id = ?", roleID).
		First(&role).Error
	return &role, err
}
