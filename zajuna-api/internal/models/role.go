package models

type Role struct {
	ID        int    `gorm:"column:id;primaryKey"`
	ShortName string `gorm:"column:shortname"`
	Name      string `gorm:"column:name"`
}

func (Role) TableName() string {
	return "mdl_role"
}
