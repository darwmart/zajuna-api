package models

type CourseCompetency struct {
	ID           uint  `gorm:"column:id;primaryKey" json:"id"`
	CourseID     uint  `gorm:"column:courseid" json:"courseid"`
	CompetencyID uint  `gorm:"column:competencyid" json:"competencyid"`
	RuleOutcome  int16 `gorm:"column:ruleoutcome" json:"ruleoutcome"`
	TimeCreated  int64 `gorm:"column:timecreated" json:"timecreated"`
	TimeModified int64 `gorm:"column:timemodified" json:"timemodified"`
	UserModified uint  `gorm:"column:usermodified" json:"usermodified"`
	SortOrder    int64 `gorm:"column:sortorder" json:"sortorder"`
}
