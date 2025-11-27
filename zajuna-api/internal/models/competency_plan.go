package models

type CompetencyPlan struct {
	ID                uint   `gorm:"column:id;primaryKey" json:"id"`
	Name              string `gorm:"column:name" json:"name"`
	Description       string `gorm:"column:description" json:"description"`
	DescriptionFormat int16  `gorm:"column:descriptionformat" json:"descriptionformat"`
	UserID            uint   `gorm:"column:userid" json:"userid"`
	TemplateID        *uint  `gorm:"column:templateid" json:"templateid"`
	OrigTemplateID    *uint  `gorm:"column:origtemplateid" json:"origtemplateid"`
	Status            int16  `gorm:"column:status" json:"status"`
	DueDate           int64  `gorm:"column:duedate" json:"duedate"`
	ReviewerID        *uint  `gorm:"column:reviewerid" json:"reviewerid"`
	TimeCreated       int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified      int64  `gorm:"column:timemodified" json:"timemodified"`
	UserModified      uint   `gorm:"column:usermodified" json:"usermodified"`
}
