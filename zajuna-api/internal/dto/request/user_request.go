package request

// GetUsersRequest representa los parámetros de búsqueda de usuarios
type GetUsersRequest struct {
	Q         string `form:"q" binding:"omitempty,min=1,max=100"` // Búsqueda global (firstname, lastname, username, email)
	Firstname string `form:"firstname" binding:"omitempty,min=2,max=100"`
	Lastname  string `form:"lastname" binding:"omitempty,min=2,max=100"`
	Username  string `form:"username" binding:"omitempty,min=2,max=100"`
	Email     string `form:"email" binding:"omitempty,email"`
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=1000"`
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

	// Si hay búsqueda global (q), ignorar los demás filtros
	// La búsqueda global tiene prioridad sobre filtros individuales
	if r.Q != "" {
		filters["q"] = r.Q
		return filters
	}

	// Si no hay búsqueda global, aplicar filtros individuales
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
	ID        uint    `json:"id" binding:"required,min=1"`
	FirstName *string `json:"firstname" binding:"omitempty,min=2,max=100"`
	LastName  *string `json:"lastname" binding:"omitempty,min=2,max=100"`
	Email     *string `json:"email" binding:"omitempty,email"`
	City      *string `json:"city" binding:"omitempty,max=120"`
	Country   *string `json:"country" binding:"omitempty,len=2"`
	Lang      *string `json:"lang" binding:"omitempty,min=2,max=30"`
	Timezone  *string `json:"timezone" binding:"omitempty,max=100"`
	Phone1    *string `json:"phone1" binding:"omitempty,max=20"`
	Suspended *int    `json:"suspended" binding:"omitempty,min=0,max=1"`
	Deleted   *int    `json:"deleted" binding:"omitempty,min=0,max=1"`
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

// CreateUserRequest representa la solicitud de creación de un usuario (compatible con Moodle 4.3 core_user_create_users)
type CreateUserRequest struct {
	// Campos requeridos
	Username  string `json:"username" binding:"required,min=2,max=100"`
	FirstName string `json:"firstname" binding:"required,min=2,max=100"`
	LastName  string `json:"lastname" binding:"required,min=2,max=100"`
	Email     string `json:"email" binding:"required,email"`

	// Password (opcional si createpassword es true)
	CreatePassword *int    `json:"createpassword" binding:"omitempty,min=0,max=1"` // 1 = crear y enviar por correo
	Password       *string `json:"password" binding:"omitempty,min=8"`              // Requerido si createpassword no está presente o es 0

	// Auth
	Auth *string `json:"auth" binding:"omitempty"` // Valor por defecto: "manual"

	// Campos opcionales de perfil básico
	MailDisplay *int    `json:"maildisplay" binding:"omitempty"`
	City        *string `json:"city" binding:"omitempty,max=120"`
	Country     *string `json:"country" binding:"omitempty,len=2"`
	Timezone    *string `json:"timezone" binding:"omitempty,max=100"`
	Description *string `json:"description" binding:"omitempty"`

	// Campos fonéticos y nombres alternativos
	FirstNamePhonetic *string `json:"firstnamephonetic" binding:"omitempty,max=255"`
	LastNamePhonetic  *string `json:"lastnamephonetic" binding:"omitempty,max=255"`
	MiddleName        *string `json:"middlename" binding:"omitempty,max=255"`
	AlternateName     *string `json:"alternatename" binding:"omitempty,max=255"`

	// Información adicional
	Interests   *string `json:"interests" binding:"omitempty"` // Separados por comas
	IDNumber    *string `json:"idnumber" binding:"omitempty,max=255"`
	Institution *string `json:"institution" binding:"omitempty,max=255"`
	Department  *string `json:"department" binding:"omitempty,max=255"`
	Phone1      *string `json:"phone1" binding:"omitempty,max=20"`
	Phone2      *string `json:"phone2" binding:"omitempty,max=20"`
	Address     *string `json:"address" binding:"omitempty,max=255"`

	// Configuración de idioma y calendario
	Lang         *string `json:"lang" binding:"omitempty,min=2,max=30"`         // Valor por defecto: "es"
	CalendarType *string `json:"calendartype" binding:"omitempty,max=30"`       // Valor por defecto: "gregorian"
	Theme        *string `json:"theme" binding:"omitempty,max=50"`              // Ej: "standard"
	MailFormat   *int    `json:"mailformat" binding:"omitempty,min=0,max=1"`    // 0 = texto plano, 1 = HTML

	// Campos personalizados y preferencias
	CustomFields []UserCustomFieldRequest `json:"customfields" binding:"omitempty,dive"`
	Preferences  []UserPreferenceRequest  `json:"preferences" binding:"omitempty,dive"`
}

// UserCustomFieldRequest representa un campo personalizado del usuario
type UserCustomFieldRequest struct {
	Type  string `json:"type" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// UserPreferenceRequest representa una preferencia del usuario
type UserPreferenceRequest struct {
	Type  string `json:"type" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// CreateUsersRequest representa la solicitud de creación de múltiples usuarios
type CreateUsersRequest struct {
	Users []CreateUserRequest `json:"users" binding:"required,min=1,dive"`
}

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
