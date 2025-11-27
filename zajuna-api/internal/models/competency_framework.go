package models

type CompetencyFramework struct {
	ID                 int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShortName          string `gorm:"column:shortname" json:"shortname"`
	IDNumber           string `gorm:"column:idnumber" json:"idnumber"`
	Description        string `gorm:"column:description" json:"description"`
	DescriptionFormat  int16  `gorm:"column:descriptionformat" json:"descriptionformat"`
	Visible            int16  `gorm:"column:visible" json:"visible"`
	ScaleID            uint   `gorm:"column:scaleid" json:"scaleid"`
	ScaleConfiguration string `gorm:"column:scaleconfiguration" json:"scaleconfiguration"`
	ContextID          uint   `gorm:"column:contextid" json:"contextid"`
	Taxonomies         string `gorm:"column:taxonomies" json:"taxonomies"`
	TimeCreated        int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified       int64  `gorm:"column:timemodified" json:"timemodified"`
	UserModified       uint   `gorm:"column:usermodified" json:"usermodified"`
}
