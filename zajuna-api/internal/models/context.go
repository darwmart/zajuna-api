package models

type Context struct {
	ID           int `gorm:"column:id;primaryKey"`
	ContextLevel int `gorm:"column:contextlevel"`
	InstanceID   int `gorm:"column:instanceid"`
}

func (Context) TableName() string {
	return "mdl_context"
}
