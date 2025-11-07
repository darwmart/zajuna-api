package response

// CourseResponse representa un curso en la respuesta de la API
type CourseResponse struct {
	ID                int    `json:"id"`
	Category          int    `json:"category"`
	FullName          string `json:"fullname"`
	ShortName         string `json:"shortname"`
	IDNumber          string `json:"idnumber"`
	Summary           string `json:"summary,omitempty"`
	SummaryFormat     int    `json:"summaryformat"`
	Format            string `json:"format"`
	ShowGrades        int    `json:"showgrades"`
	NewsItems         int    `json:"newsitems"`
	StartDate         int64  `json:"startdate"`
	EndDate           int64  `json:"enddate"`
	Visible           int    `json:"visible"`
	GroupMode         int    `json:"groupmode"`
	GroupModeForce    int    `json:"groupmodeforce"`
	DefaultGroupingID int    `json:"defaultgroupingid"`
	SortOrder         int    `json:"sortorder"`
	TimeCreated       int64  `json:"timecreated"`
	TimeModified      int64  `json:"timemodified"`
}

// CourseDetailResponse representa los detalles completos de un curso
type CourseDetailResponse struct {
	ID              int64            `json:"id"`
	FullName        string           `json:"fullname"`
	ShortName       string           `json:"shortname"`
	IDNumber        string           `json:"idnumber"`
	Format          string           `json:"format"`
	Category        string           `json:"category"`
	Groupings       int64            `json:"groupings"`
	Groups          int64            `json:"groups"`
	RoleAssignments map[string]int64 `json:"role_assignments"`
	EnrolMethods    []string         `json:"enrol_methods"`
	Sections        []string         `json:"sections"`
	StartDate       int64            `json:"startdate"`
	EndDate         int64            `json:"enddate"`
	Visible         bool             `json:"visible"`
	TimeCreated     int64            `json:"timecreated"`
	TimeModified    int64            `json:"timemodified"`
}

// CourseListResponse representa la respuesta de listado de cursos
type CourseListResponse struct {
	Courses    []CourseResponse `json:"courses"`
	Pagination *PaginationMeta  `json:"pagination,omitempty"`
}

// DeleteCoursesResponse representa la respuesta de eliminación de cursos
type DeleteCoursesResponse struct {
	Message  string    `json:"message,omitempty"`
	Deleted  int       `json:"deleted"`
	Warnings []Warning `json:"warnings,omitempty"`
}

// Warning representa una advertencia en la operación
type Warning struct {
	Item        string `json:"item"`
	ItemID      int    `json:"itemid"`
	WarningCode string `json:"warningcode"`
	Message     string `json:"message"`
}

// CreateCourseResponse representa la respuesta de creación de curso
type CreateCourseResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Message  string `json:"message"`
}

// UpdateCourseResponse representa la respuesta de actualización de curso
type UpdateCourseResponse struct {
	Message string   `json:"message"`
	Updated bool     `json:"updated"`
	Errors  []string `json:"errors,omitempty"`
}

// UpdateCoursesResponse representa la respuesta de actualización múltiple (Moodle format)
type UpdateCoursesResponse struct {
	Warnings []Warning `json:"warnings,omitempty"`
	//WarningsCount int       `json:"warningscount"`
}

/*
1
1
1
1
1
1
11
11
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
1
*/
const (
	// CourseFormatWeeks representa el formato de curso "weeks"
	CourseFormatWeeks = "weeks"
	// CourseFormatTopics representa el formato de curso "topics"
	CourseFormatTopics = "topics"
	// CourseFormatSocial representa el formato de curso "social"
	CourseFormatSocial = "social"
)
