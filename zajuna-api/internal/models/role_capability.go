package models

// RoleCapability representa los permisos de un rol en un contexto
type RoleCapability struct {
	ID           int    `gorm:"column:id;primaryKey"`
	ContextID    int    `gorm:"column:contextid"`
	RoleID       int    `gorm:"column:roleid"`
	Capability   string `gorm:"column:capability"`
	Permission   int    `gorm:"column:permission"` // Valores: 0=INHERIT, 1=ALLOW, -1=PROHIBIT, -1000=PREVENT
	TimeModified int64  `gorm:"column:timemodified"`
	ModifierID   int    `gorm:"column:modifierid"`
}

func (RoleCapability) TableName() string {
	return "mdl_role_capabilities"
}

// Las constantes de permisos están definidas en context.go:
// - CapInherit  = 0     // CAP_INHERIT
// - CapAllow    = 1     // CAP_ALLOW
// - CapPrevent  = -1000 // CAP_PREVENT
// - CapProhibit = -1    // CAP_PROHIBIT
