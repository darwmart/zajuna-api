package models

type User struct {
	ID                uint              `gorm:"column:id;primaryKey" json:"id"`
	Username          string            `gorm:"column:username" json:"username"`
	FirstName         string            `gorm:"column:firstname" json:"firstname"`
	LastName          string            `gorm:"column:lastname" json:"lastname"`
	FullName          string            `gorm:"-" json:"fullname"`
	Email             string            `gorm:"column:email" json:"email"`
	Address           string            `gorm:"column:address" json:"address"`
	Phone1            string            `gorm:"column:phone1" json:"phone1"`
	Phone2            string            `gorm:"column:phone2" json:"phone2"`
	Department        string            `gorm:"column:department" json:"department"`
	Institution       string            `gorm:"column:institution" json:"institution"`
	IDNumber          string            `gorm:"column:idnumber" json:"idnumber"`
	Interests         string            `gorm:"column:interests" json:"interests"`
	FirstAccess       int64             `gorm:"column:firstaccess" json:"firstaccess"`
	LastAccess        int64             `gorm:"column:lastaccess" json:"lastaccess"`
	Auth              string            `gorm:"column:auth" json:"auth"`
	Suspended         int               `gorm:"column:suspended" json:"suspended"`
	Confirmed         int               `gorm:"column:confirmed" json:"confirmed"`
	Lang              string            `gorm:"column:lang" json:"lang"`
	Theme             string            `gorm:"column:theme" json:"theme"`
	Timezone          string            `gorm:"column:timezone" json:"timezone"`
	MailFormat        int               `gorm:"column:mailformat" json:"mailformat"`
	Description       string            `gorm:"column:description" json:"description"`
	DescriptionFormat int               `gorm:"column:descriptionformat" json:"descriptionformat"`
	City              string            `gorm:"column:city" json:"city"`
	Country           string            `gorm:"column:country" json:"country"`
	ProfileImageSmall string            `gorm:"-" json:"profileimageurlsmall"`
	ProfileImage      string            `gorm:"-" json:"profileimageurl"`
	CustomFields      []UserCustomField `gorm:"-" json:"customfields,omitempty"`
	Preferences       []UserPreference  `gorm:"-" json:"preferences,omitempty"`
}

// Nombre de la tabla real en Moodle
func (User) TableName() string {
	return "mdl_user"
}

// Campos adicionales del usuario (no pertenecen directamente a mdl_user)
type UserCustomField struct {
	Type         string `json:"type"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayvalue,omitempty"`
	Name         string `json:"name"`
	ShortName    string `json:"shortname"`
}

type UserPreference struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type APIResponse struct {
	Users       []User `json:"users"`
	Total       int    `json:"total"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	TotalPages  int    `json:"totalPages"`
	HasNext     bool   `json:"hasNext"`
	HasPrevious bool   `json:"hasPrevious"`
}
