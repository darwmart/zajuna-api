package models

// MoodleRole representa un rol en la base de datos de Moodle (mdl_role)
// Se usa para el sistema de permisos, diferente del Role usado en EnrolledUser
type MoodleRole struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name"`
	ShortName   string `gorm:"column:shortname"`
	Description string `gorm:"column:description"`
	SortOrder   int    `gorm:"column:sortorder"`
	Archetype   string `gorm:"column:archetype"`
}

func (MoodleRole) TableName() string {
	return "mdl_role"
}
