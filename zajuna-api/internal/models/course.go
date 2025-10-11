package models

type Course struct {
	ID                uint   `gorm:"column:id;primaryKey" json:"id"`
	Category          uint   `gorm:"column:category" json:"category"`
	FullName          string `gorm:"column:fullname" json:"fullname"`
	ShortName         string `gorm:"column:shortname" json:"shortname"`
	IDNumber          string `gorm:"column:idnumber" json:"idnumber"`
	Summary           string `gorm:"column:summary" json:"summary"`
	SummaryFormat     int    `gorm:"column:summaryformat" json:"summaryformat"`
	Format            string `gorm:"column:format" json:"format"`
	ShowGrades        int    `gorm:"column:showgrades" json:"showgrades"`
	NewsItems         int    `gorm:"column:newsitems" json:"newsitems"`
	StartDate         int64  `gorm:"column:startdate" json:"startdate"`
	EndDate           int64  `gorm:"column:enddate" json:"enddate"`
	Visible           int    `gorm:"column:visible" json:"visible"`
	GroupMode         int    `gorm:"column:groupmode" json:"groupmode"`
	GroupModeForce    int    `gorm:"column:groupmodeforce" json:"groupmodeforce"`
	DefaultGroupingID int    `gorm:"column:defaultgroupingid" json:"defaultgroupingid"`
	SortOrder         int    `gorm:"column:sortorder" json:"sortorder"`
	TimeCreated       int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified      int64  `gorm:"column:timemodified" json:"timemodified"`
}
