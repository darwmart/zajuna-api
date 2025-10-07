package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	FullName  string `json:"fullname"`
	Email     string `json:"email"`
	Auth      string `json:"auth"`
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
