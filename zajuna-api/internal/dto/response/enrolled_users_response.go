package response

// EnrolledUserResponse representa un usuario matriculado con toda su información
type EnrolledUserResponse struct {
	ID                   int64                `json:"id"`
	Username             string               `json:"username,omitempty"`
	FirstName            string               `json:"firstname,omitempty"`
	LastName             string               `json:"lastname,omitempty"`
	FullName             string               `json:"fullname"`
	Email                string               `json:"email,omitempty"`
	Address              string               `json:"address,omitempty"`
	Phone1               string               `json:"phone1,omitempty"`
	Phone2               string               `json:"phone2,omitempty"`
	Department           string               `json:"department,omitempty"`
	Institution          string               `json:"institution,omitempty"`
	IDNumber             string               `json:"idnumber,omitempty"`
	Interests            string               `json:"interests,omitempty"`
	FirstAccess          int64                `json:"firstaccess,omitempty"`
	LastAccess           int64                `json:"lastaccess,omitempty"`
	LastCourseAccess     int64                `json:"lastcourseaccess,omitempty"`
	Description          string               `json:"description,omitempty"`
	DescriptionFormat    int                  `json:"descriptionformat,omitempty"`
	City                 string               `json:"city,omitempty"`
	Country              string               `json:"country,omitempty"`
	ProfileImageURLSmall string               `json:"profileimageurlsmall,omitempty"`
	ProfileImageURL      string               `json:"profileimageurl,omitempty"`
	CustomFields         []UserCustomField    `json:"customfields,omitempty"`
	Groups               []EnrolledUserGroup  `json:"groups,omitempty"`
	Roles                []EnrolledUserRole   `json:"roles,omitempty"`
	Preferences          []UserPreference     `json:"preferences,omitempty"`
	EnrolledCourses      []UserEnrolledCourse `json:"enrolledcourses,omitempty"`
}

// EnrolledUserGroup representa un grupo al que pertenece el usuario
type EnrolledUserGroup struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DescriptionFormat int    `json:"descriptionformat"`
}

// EnrolledUserRole representa un rol del usuario en el curso
type EnrolledUserRole struct {
	RoleID    int64  `json:"roleid"`
	Name      string `json:"name"`
	ShortName string `json:"shortname"`
	SortOrder int    `json:"sortorder"`
}

// UserEnrolledCourse representa un curso en el que está matriculado el usuario
type UserEnrolledCourse struct {
	ID        int64  `json:"id"`
	FullName  string `json:"fullname"`
	ShortName string `json:"shortname"`
}

// EnrolledUsersListResponse representa la respuesta de la lista de usuarios matriculados
type EnrolledUsersListResponse struct {
	Users []EnrolledUserResponse `json:"users"`
	Total int                    `json:"total"`
}
