package models

type Category struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	IDNumber          *string `json:"idnumber,omitempty"`
	Description       *string `json:"description,omitempty"`
	DescriptionFormat int     `json:"descriptionformat,omitempty"`
	Parent            int     `json:"parent"`
	SortOrder         int     `json:"sortorder"`
	CourseCount       int     `json:"coursecount"`
	Visible           int     `json:"visible"`
	Depth             int     `json:"depth"`
	Path              string  `json:"path"`
	Theme             *string `json:"theme,omitempty"`
}
