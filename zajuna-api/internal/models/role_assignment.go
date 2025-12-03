package models

type RoleAssignment struct {
	ID           int    `gorm:"column:id;primaryKey"`
	RoleID       int    `gorm:"column:roleid"`
	UserID       int    `gorm:"column:userid"`
	ContextID    int    `gorm:"column:contextid"`
	Component    string `gorm:"column:component"`
	ItemID       int    `gorm:"column:itemid"`
	TimeModified int64  `gorm:"column:timemodified"`
	ModifierID   int    `gorm:"column:modifierid"`
	SortOrder    int    `gorm:"column:sortorder"`
}

func (RoleAssignment) TableName() string {
	return "mdl_role_assignments"
}
