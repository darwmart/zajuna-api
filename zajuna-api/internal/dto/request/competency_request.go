package request

type CreateCompetencyRequest struct {
	ShortName             string `json:"shortname" binding:"omitempty,min=1"`
	Description           string `json:"description" binding:"omitempty,min=1"`
	DescriptionFormat     int16  `json:"descriptionformat" binding:"omitempty,min=1,oneof=0 1 2 4"`
	IDNumber              string `json:"idnumber" binding:"omitempty,min=1"`
	CompetencyFrameworkID uint   `json:"competencyframeworkid" binding:"omitempty,min=1"`
	ParentID              uint   `json:"parentid" binding:"omitempty,min=1"`
	RuleType              string `json:"ruletype" binding:"omitempty,min=1"`
	RuleOutcome           int16  `json:"ruleoutcome" binding:"omitempty,min=1,oneof=0 1 2 3"`
	RuleConfig            string `json:"ruleconfig" binding:"omitempty,min=1"`
	ScaleID               uint   `json:"scaleid" binding:"omitempty,min=1"`
	ScaleConfiguration    string `json:"scaleconfiguration" binding:"omitempty,min=1"`
}
