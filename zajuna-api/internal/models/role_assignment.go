package models

type RoleAssignment struct {
	ID        int `gorm:"column:id;primaryKey"`
	RoleID    int `gorm:"column:roleid"`
	UserID    int `gorm:"column:userid"`
	ContextID int `gorm:"column:contextid"`
}

func (RoleAssignment) TableName() string {
	return "mdl_role_assignments"
}
