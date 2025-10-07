package models

type Course struct {
	ID            int    `json:"id"`
	ShortName     string `json:"shortname"`
	CategoryID    int    `json:"categoryid"`
	FullName      string `json:"fullname"`
	DisplayName   string `json:"displayname"`
	Summary       string `json:"summary"`
	SummaryFormat int    `json:"summaryformat"`
	Format        string `json:"format"`
	Visible       int    `json:"visible"`
	StartDate     int64  `json:"startdate"`
	EndDate       int64  `json:"enddate"`
	TimeCreated   int64  `json:"timecreated"`
	TimeModified  int64  `json:"timemodified"`
}
