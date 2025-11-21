package request

type CreateCompetencyFrameworkRequest struct {
	ShortName          string `json:"shortname" binding:"omitempty,min=1"`
	IDNumber           string `json:"idnumber" binding:"omitempty,min=1"`
	Description        string `json:"description" binding:"omitempty,min=1"`
	DescriptionFormat  int16  `json:"descriptionformat" binding:"omitempty,min=1,oneof=0 1 2 4"`
	Visible            int16  `json:"visible" binding:"omitempty,min=1"`
	ScaleID            uint   `json:"scaleid" binding:"omitempty,min=1"`
	ScaleConfiguration string `json:"scaleconfiguration" binding:"omitempty,min=1"`
	ContextID          uint   `json:"contextid" binding:"omitempty,min=1"`
	Taxonomies         string `json:"taxonomies" binding:"omitempty,min=1"`
}
