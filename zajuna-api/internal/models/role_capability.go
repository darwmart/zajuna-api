package models

type RoleCapability struct {
	ID           int    `gorm:"column:id;primaryKey"`
	ContextID    int64  `gorm:"column:contextid"`
	RoleID       int64  `gorm:"column:roleid"`
	Capability   string `gorm:"column:capability"`
	Permission   int64  `gorm:"column:permission"`
	TimeModified int64  `gorm:"column:timemodified"`
	ModifierID   int64  `gorm:"column:modifierid"`
}

func (RoleCapability) TableName() string {
	return "mdl_role_capabilities"
}
