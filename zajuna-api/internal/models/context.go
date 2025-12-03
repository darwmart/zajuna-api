package models

// Context levels (niveles de contexto de Moodle)
const (
	ContextSystem    = 10
	ContextUser      = 30
	ContextCoursecat = 40
	ContextCourse    = 50
	ContextModule    = 70
	ContextBlock     = 80
)

// Permission levels (niveles de permiso de Moodle)
// Estos valores DEBEN coincidir EXACTAMENTE con Moodle (lib/accesslib.php)
const (
	CapInherit  = 0     // CAP_INHERIT - Heredar del contexto padre
	CapAllow    = 1     // CAP_ALLOW - Permitir explícitamente
	CapPrevent  = -1000 // CAP_PREVENT - Prevenir (sin bloquear)
	CapProhibit = -1    // CAP_PROHIBIT - Prohibir (bloquea todo, no puede ser sobrescrito)
)

type Context struct {
	ID           int    `gorm:"column:id;primaryKey"`
	ContextLevel int    `gorm:"column:contextlevel"`
	InstanceID   int    `gorm:"column:instanceid"`
	Path         string `gorm:"column:path"`
	Depth        int    `gorm:"column:depth"`
	Locked       int    `gorm:"column:locked"`
}

func (Context) TableName() string {
	return "mdl_context"
}

// AccessData representa los datos de acceso de un usuario (replica $USER->access de Moodle)
// Esta estructura se usa para cachear las asignaciones de roles y capabilities de un usuario
type AccessData struct {
	// RA (Role Assignments) - mapa de context path a roleids
	// Ejemplo: RA["/1/"] = map[5]5 significa que el usuario tiene el role 5 en el contexto /1/
	RA map[string]map[int]int

	// RSW (Role SWitch) - mapa de context path a roleid para role switching
	// Usado cuando un admin o teacher cambia temporalmente de rol
	RSW map[string]int

	// Time - timestamp de cuando se cargó este accessdata
	Time int64
}

// RoleDefinition representa las capabilities de un rol en todos los contextos
// Replica la estructura que retorna get_role_definitions() de Moodle
// Estructura: RoleDefinition[roleid][path][capability] = permission
type RoleDefinition map[string]map[string]int

// RoleDefinitions es un mapa de roleID a RoleDefinition
type RoleDefinitions map[int]RoleDefinition
