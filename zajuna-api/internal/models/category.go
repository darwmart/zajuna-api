package models

type Category struct {
	ID                uint   `gorm:"column:id;primaryKey" json:"id"`
	Name              string `gorm:"column:name" json:"name"`
	IDNumber          string `gorm:"column:idnumber" json:"idnumber"`
	Description       string `gorm:"column:description" json:"description"`
	DescriptionFormat int    `gorm:"column:descriptionformat" json:"descriptionformat"`
	Parent            int    `gorm:"column:parent" json:"parent"`
	SortOrder         int    `gorm:"column:sortorder" json:"sortorder"`
	CourseCount       int    `gorm:"column:coursecount" json:"coursecount"`
	Visible           int    `gorm:"column:visible" json:"visible"`
	Depth             int    `gorm:"column:depth" json:"depth"`
	Path              string `gorm:"column:path" json:"path"`
	Theme             string `gorm:"column:theme" json:"theme"`
}
