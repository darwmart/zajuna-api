package request

// GetUsersRequest representa los parámetros de búsqueda de usuarios
type GetUsersRequest struct {
	Firstname string `form:"firstname" binding:"omitempty,min=2,max=100"`
	Lastname  string `form:"lastname" binding:"omitempty,min=2,max=100"`
	Username  string `form:"username" binding:"omitempty,min=2,max=100"`
	Email     string `form:"email" binding:"omitempty,email"`
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
}

// SetDefaults establece valores por defecto para la paginación
func (r *GetUsersRequest) SetDefaults() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 15
	}
}

// ToFilterMap convierte la request a un mapa de filtros
func (r *GetUsersRequest) ToFilterMap() map[string]string {
	filters := make(map[string]string)

	if r.Firstname != "" {
		filters["firstname"] = r.Firstname
	}
	if r.Lastname != "" {
		filters["lastname"] = r.Lastname
	}
	if r.Username != "" {
		filters["username"] = r.Username
	}
	if r.Email != "" {
		filters["email"] = r.Email
	}

	return filters
}

// UpdateUserRequest representa la solicitud de actualización de un usuario
type UpdateUserRequest struct {
	ID        uint   `json:"id" binding:"required,min=1"`
	FirstName string `json:"firstname" binding:"required,min=2,max=100"`
	LastName  string `json:"lastname" binding:"required,min=2,max=100"`
	Email     string `json:"email" binding:"required,email"`
	City      string `json:"city" binding:"omitempty,max=120"`
	Country   string `json:"country" binding:"omitempty,len=2"`
	Lang      string `json:"lang" binding:"omitempty,min=2,max=30"`
	Timezone  string `json:"timezone" binding:"omitempty,max=100"`
	Phone1    string `json:"phone1" binding:"omitempty,max=20"`
}

// UpdateUsersRequest representa la solicitud de actualización de múltiples usuarios
type UpdateUsersRequest struct {
	Users []UpdateUserRequest `json:"users" binding:"required,min=1,dive"`
}

// DeleteUsersRequest representa la solicitud de eliminación de usuarios
type DeleteUsersRequest struct {
	UserIDs []int `json:"userids" binding:"required,min=1,dive,min=1"`
}

// Validate valida que los IDs sean únicos
func (r *DeleteUsersRequest) Validate() error {
	seen := make(map[int]bool)
	for _, id := range r.UserIDs {
		if seen[id] {
			return ErrDuplicateIDs
		}
		seen[id] = true
	}
	return nil
}

// Errores personalizados para validación
var (
	ErrDuplicateIDs = &ValidationError{Field: "userids", Message: "IDs duplicados detectados"}
)

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
