package models

type User struct {
	ID        uint   `gorm:"column:id;primaryKey" json:"id"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
	Username  string `gorm:"column:username" json:"username"`
	Email     string `gorm:"column:email" json:"email"`
	City      string `gorm:"column:city" json:"city"`
	Country   string `gorm:"column:country" json:"country"`
}

func (User) TableName() string {
	return "mdl_user"
}

type APIResponse struct {
	Users           []User `json:"users"`
	Total           int    `json:"total"`
	CurrentPage     int    `json:"currentPage"`
	PageSize        int    `json:"pageSize"`
	TotalPages      int    `json:"totalPages"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
}
