package response

// UserResponse representa un usuario en la respuesta de la API
type UserResponse struct {
	ID                uint              `json:"id"`
	Username          string            `json:"username"`
	FirstName         string            `json:"firstname"`
	LastName          string            `json:"lastname"`
	FullName          string            `json:"fullname"`
	Email             string            `json:"email"`
	Address           string            `json:"address,omitempty"`
	Phone1            string            `json:"phone1,omitempty"`
	Phone2            string            `json:"phone2,omitempty"`
	Department        string            `json:"department,omitempty"`
	Institution       string            `json:"institution,omitempty"`
	IDNumber          string            `json:"idnumber,omitempty"`
	Interests         string            `json:"interests,omitempty"`
	FirstAccess       int64             `json:"firstaccess"`
	LastAccess        int64             `json:"lastaccess"`
	Auth              string            `json:"auth"`
	Suspended         int               `json:"suspended"`
	Confirmed         int               `json:"confirmed"`
	Lang              string            `json:"lang"`
	Theme             string            `json:"theme,omitempty"`
	Timezone          string            `json:"timezone"`
	MailFormat        int               `json:"mailformat"`
	Description       string            `json:"description,omitempty"`
	DescriptionFormat int               `json:"descriptionformat"`
	City              string            `json:"city"`
	Country           string            `json:"country"`
	ProfileImageSmall string            `json:"profileimageurlsmall,omitempty"`
	ProfileImage      string            `json:"profileimageurl,omitempty"`
	CustomFields      []UserCustomField `json:"customfields,omitempty"`
	Preferences       []UserPreference  `json:"preferences,omitempty"`
}

// UserCustomField representa un campo personalizado del usuario
type UserCustomField struct {
	Type         string `json:"type"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayvalue,omitempty"`
	Name         string `json:"name"`
	ShortName    string `json:"shortname"`
}

// UserPreference representa una preferencia del usuario
type UserPreference struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// UserListResponse representa la respuesta de listado de usuarios (paginado)
type UserListResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination PaginationMeta `json:"pagination"`
}

// UpdateUserResponse representa la respuesta de actualización de usuario
type UpdateUserResponse struct {
	Message  string   `json:"message"`
	Updated  int64    `json:"updated"`
	Warnings []string `json:"warnings"`
}

// DeleteUserResponse representa la respuesta de eliminación de usuarios
type DeleteUserResponse struct {
	Message string   `json:"message"`
	Deleted int      `json:"deleted"`
	Errors  []string `json:"errors,omitempty"`
}

// ToggleUserStatusResponse representa la respuesta al cambiar el estado de un usuario
type ToggleUserStatusResponse struct {
	Message   string `json:"message"`
	UserID    uint   `json:"user_id"`
	NewStatus int    `json:"new_status"`
	StatusText string `json:"status_text"`
}
