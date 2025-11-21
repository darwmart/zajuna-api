package models

type CompetencyFramework struct {
	ID                 int    `gorm:"column:id;primaryKey;autoIncrement"`
	ShortName          string `gorm:"column:shortname"`
	IDNumber           string `gorm:"column:idnumber"`
	Description        string `gorm:"column:description"`
	DescriptionFormat  int16  `gorm:"column:descriptionformat"`
	Visible            int16  `gorm:"column:visible"`
	ScaleID            uint   `gorm:"column:scaleid"`
	ScaleConfiguration string `gorm:"column:scaleconfiguration"`
	ContextID          uint   `gorm:"column:contextid"`
	Taxonomies         string `gorm:"column:taxonomies"`
	TimeCreated        int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified       int64  `gorm:"column:timemodified" json:"timemodified"`
	UserModified       uint   `gorm:"column:usermodified" json:"usermodified"`
}
