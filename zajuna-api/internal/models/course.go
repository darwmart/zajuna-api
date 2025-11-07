package models

type Course struct {
	ID                int    `gorm:"column:id" json:"id"`
	Category          int    `gorm:"column:category" json:"category"`
	FullName          string `gorm:"column:fullname" json:"fullName"`
	ShortName         string `gorm:"column:shortname" json:"shortName"`
	IDNumber          string `gorm:"column:idnumber" json:"idNumber"`
	Summary           string `gorm:"column:summary" json:"summary"`
	SummaryFormat     int    `gorm:"column:summaryformat" json:"summaryformat"`
	Format            string `gorm:"column:format" json:"format"`
	ShowGrades        int    `gorm:"column:showgrades" json:"showgrades"`
	NewsItems         int    `gorm:"column:newsitems" json:"newsitems"`
	StartDate         int64  `gorm:"column:startdate" json:"startdate"`
	EndDate           int64  `gorm:"column:enddate" json:"enddate"`
	NumSections       int    `gorm:"column:numsections" json:"numsections"`
	MaxBytes          int    `gorm:"column:maxbytes" json:"maxbytes"`
	ShowReports       int    `gorm:"column:showreports" json:"showreports"`
	Visible           int    `gorm:"column:visible" json:"visible"`
	HiddenSections    int    `gorm:"column:hiddensections" json:"hiddensections"`
	GroupMode         int    `gorm:"column:groupmode" json:"groupmode"`
	GroupModeForce    int    `gorm:"column:groupmodeforce" json:"groupmodeforce"`
	DefaultGroupingID int    `gorm:"column:defaultgroupingid" json:"defaultgroupingid"`
	EnableCompletion  int    `gorm:"column:enablecompletion" json:"enablecompletion"`
	CompletionNotify  int    `gorm:"column:completionnotify" json:"completionnotify"`
	Lang              string `gorm:"column:lang" json:"lang"`
	ForceTheme        string `gorm:"column:theme" json:"forcetheme"`
	SortOrder         int    `gorm:"column:sortorder" json:"sortorder"`
	TimeCreated       int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified      int64  `gorm:"column:timemodified" json:"timemodified"`
}
