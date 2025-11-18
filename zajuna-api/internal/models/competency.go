package models

type Competency struct {
	ID                    int    `gorm:"column:id" json:"id"`
	ShortName             string `gorm:"column:shortname" json:"shortName"`
	Description           string `gorm:"column:description" json:"description"`
	DescriptionFormat     int16  `gorm:"column:descriptionformat" json:"descriptionFormat"`
	IDNumber              string `gorm:"column:idnumber" json:"idNumber"`
	CompetencyFrameworkID uint   `gorm:"column:competencyframeworkid" json:"competencyframeworkid"`
	ParentID              uint   `gorm:"column:parentid" json:"parentid"`
	Path                  string `gorm:"column:path" json:"path"`
	SortOrder             int64  `gorm:"column:sortorder" json:"sortorder"`
	RuleType              string `gorm:"column:ruletype" json:"ruletype"`
	RuleOutcome           int16  `gorm:"column:ruleoutcome" json:"ruleoutcome"`
	RuleConfig            string `gorm:"column:ruleconfig" json:"roleconfig"`
	ScaleID               uint   `gorm:"column:scaleid" json:"scaleid"`
	ScaleConfiguration    string `gorm:"column:scaleconfiguration" json:"scaleconfiguration"`
	TimeCreated           int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified          int64  `gorm:"column:timemodified" json:"timemodified"`
	UserModified          uint   `gorm:"column:usermodified" json:"usermodified"`
}
