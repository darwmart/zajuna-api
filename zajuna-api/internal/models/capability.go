package models

// Capability representa una capacidad/permiso específico en Moodle
type Capability struct {
	Name         string `gorm:"column:name;primaryKey"`
	CapType      string `gorm:"column:captype"` // 'read' o 'write'
	ContextLevel int    `gorm:"column:contextlevel"`
	Component    string `gorm:"column:component"`
	RiskBitmask  int    `gorm:"column:riskbitmask"`
}

func (Capability) TableName() string {
	return "mdl_capabilities"
}
