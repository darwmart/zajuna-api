package models

type Config struct {
	ID    int64  `gorm:"column:id;primaryKey"`
	Name  string `gorm:"column:name"`
	Value string `gorm:"column:value"`
}

func (Config) TableName() string {
	return "mdl_config"
}
