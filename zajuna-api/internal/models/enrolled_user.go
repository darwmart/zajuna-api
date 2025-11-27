package models

// CustomField representa un campo personalizado del usuario
type CustomField struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
}

// Group representa un grupo al que pertenece el usuario
type Group struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DescriptionFormat int    `json:"descriptionformat"`
}

// Role representa un rol del usuario en el curso
type Role struct {
	RoleID    int    `json:"roleid"`
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
	Sortorder int    `json:"sortorder"`
}

// Preference representa una preferencia del usuario
type Preference struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// EnrolledCourse representa un curso en el que está matriculado el usuario
type EnrolledCourse struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Shortname string `json:"shortname"`
}

// EnrolledUser representa un usuario matriculado con toda su información
type EnrolledUser struct {
	ID                   int              `json:"id"`
	Username             string           `json:"username,omitempty"`
	Firstname            string           `json:"firstname,omitempty"`
	Lastname             string           `json:"lastname,omitempty"`
	Fullname             string           `json:"fullname"`
	Email                string           `json:"email,omitempty"`
	Address              string           `json:"address,omitempty"`
	Phone1               string           `json:"phone1,omitempty"`
	Phone2               string           `json:"phone2,omitempty"`
	Department           string           `json:"department,omitempty"`
	Institution          string           `json:"institution,omitempty"`
	Idnumber             string           `json:"idnumber,omitempty"`
	Interests            string           `json:"interests,omitempty"`
	Firstaccess          int64            `json:"firstaccess,omitempty"`
	Lastaccess           int64            `json:"lastaccess,omitempty"`
	LastCourseAccess     int64            `json:"lastcourseaccess,omitempty"`
	Description          string           `json:"description,omitempty"`
	DescriptionFormat    int              `json:"descriptionformat,omitempty"`
	City                 string           `json:"city,omitempty"`
	Country              string           `json:"country,omitempty"`
	ProfileImageURLSmall string           `json:"profileimageurlsmall,omitempty"`
	ProfileImageURL      string           `json:"profileimageurl,omitempty"`
	CustomFields         []CustomField    `json:"customfields,omitempty"`
	Groups               []Group          `json:"groups,omitempty"`
	Roles                []Role           `json:"roles,omitempty"`
	Preferences          []Preference     `json:"preferences,omitempty"`
	EnrolledCourses      []EnrolledCourse `json:"enrolledcourses,omitempty"`
}
