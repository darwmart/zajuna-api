package request

// GetCoursesRequest representa los parámetros para listar cursos
type GetCoursesRequest struct {
	CategoryID int  `form:"categoryid" binding:"omitempty,min=1"`
	Page       int  `form:"page" binding:"omitempty,min=1"`
	Limit      int  `form:"limit" binding:"omitempty,min=1,max=100"`
	Visible    *int `form:"visible" binding:"omitempty,oneof=0 1"`
}

// SetDefaults establece valores por defecto
func (r *GetCoursesRequest) SetDefaults() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 20
	}
}

// HasCategoryFilter verifica si hay filtro de categoría
func (r *GetCoursesRequest) HasCategoryFilter() bool {
	return r.CategoryID > 0
}

// GetCourseDetailsRequest representa los parámetros para obtener detalles de un curso
type GetCourseDetailsRequest struct {
	IDNumber string `uri:"idnumber" binding:"required"`
}

// DeleteCoursesRequest representa la solicitud de eliminación de cursos
type DeleteCoursesRequest struct {
	CourseIDs []int `json:"courseid" binding:"required,min=1,dive,min=1"`
}

// Validate valida que los IDs sean únicos y no incluyan el curso site (ID=1)
func (r *DeleteCoursesRequest) Validate() error {
	seen := make(map[int]bool)
	for _, id := range r.CourseIDs {
		// Validar que no intenten eliminar el curso site (ID=1)
		if id == 1 {
			return &ValidationError{
				Field:   "courseid",
				Message: "No se puede eliminar el coursid",
			}
		}
		// Validar duplicados
		if seen[id] {
			return &ValidationError{
				Field:   "courseid",
				Message: "id duplicado detectado",
			}
		}
		seen[id] = true
	}
	return nil
}

// CreateCourseRequest representa la solicitud de creación de un curso
type CreateCourseRequest struct {
	CategoryID    int    `json:"category" binding:"required,min=1"`
	FullName      string `json:"fullname" binding:"required,min=3,max=254"`
	ShortName     string `json:"shortname" binding:"required,min=1,max=100"`
	IDNumber      string `json:"idnumber" binding:"omitempty,max=100"`
	Summary       string `json:"summary" binding:"omitempty"`
	SummaryFormat int    `json:"summaryformat" binding:"omitempty,oneof=0 1 2 4"`
	Format        string `json:"format" binding:"omitempty,oneof=weeks topics social"`
	StartDate     int64  `json:"startdate" binding:"omitempty,min=0"`
	EndDate       int64  `json:"enddate" binding:"omitempty,min=0"`
	Visible       int    `json:"visible" binding:"omitempty,oneof=0 1"`
}

// Validate valida reglas de negocio adicionales
func (r *CreateCourseRequest) Validate() error {
	// Validar que enddate sea mayor que startdate
	if r.EndDate > 0 && r.StartDate > 0 && r.EndDate <= r.StartDate {
		return &ValidationError{
			Field:   "enddate",
			Message: "La fecha de fin debe ser posterior a la fecha de inicio",
		}
	}
	return nil
}

// CourseFormatOption representa una opción de formato de curso
type CourseFormatOption struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// CustomField representa un campo personalizado de curso
type CustomField struct {
	ShortName string `json:"shortname" binding:"required"`
	Value     string `json:"value" binding:"required"`
}

// UpdateCourseRequest representa la solicitud de actualización de un curso
type UpdateCourseRequest struct {
	ID                  int                  `json:"id" binding:"required,min=1"`
	FullName            string               `json:"fullname" binding:"omitempty,min=3,max=254"`
	ShortName           string               `json:"shortname" binding:"omitempty,min=1,max=100"`
	CategoryID          *int                 `json:"categoryid" binding:"omitempty,min=1"`
	IDNumber            string               `json:"idnumber" binding:"omitempty,max=100"`
	Summary             string               `json:"summary" binding:"omitempty"`
	SummaryFormat       *int                 `json:"summaryformat" binding:"omitempty,oneof=0 1 2 4"`
	Format              string               `json:"format" binding:"omitempty,oneof=weeks topics social"`
	ShowGrades          *int                 `json:"showgrades" binding:"omitempty,oneof=0 1"`
	NewsItems           *int                 `json:"newsitems" binding:"omitempty,min=0"`
	StartDate           *int64               `json:"startdate" binding:"omitempty,min=0"`
	EndDate             *int64               `json:"enddate" binding:"omitempty,min=0"`
	NumSections         *int                 `json:"numsections" binding:"omitempty,min=0"`
	MaxBytes            *int                 `json:"maxbytes" binding:"omitempty,min=0"`
	ShowReports         *int                 `json:"showreports" binding:"omitempty,oneof=0 1"`
	Visible             *int                 `json:"visible" binding:"omitempty,oneof=0 1"`
	HiddenSections      *int                 `json:"hiddensections" binding:"omitempty,min=0"`
	GroupMode           *int                 `json:"groupmode" binding:"omitempty,oneof=0 1 2"`
	GroupModeForce      *int                 `json:"groupmodeforce" binding:"omitempty,oneof=0 1"`
	DefaultGroupingID   *int                 `json:"defaultgroupingid" binding:"omitempty,min=0"`
	EnableCompletion    *int                 `json:"enablecompletion" binding:"omitempty,oneof=0 1"`
	CompletionNotify    *int                 `json:"completionnotify" binding:"omitempty,oneof=0 1"`
	Lang                string               `json:"lang" binding:"omitempty,max=30"`
	ForceTheme          string               `json:"forcetheme" binding:"omitempty,max=50"`
	CourseFormatOptions []CourseFormatOption `json:"courseformatoptions" binding:"omitempty,dive"`
	CustomFields        []CustomField        `json:"customfields" binding:"omitempty,dive"`
}

// UpdateCoursesRequest representa la solicitud de actualización múltiple de cursos (Moodle format)
type UpdateCoursesRequest struct {
	Courses []UpdateCourseRequest `json:"courses" binding:"required,min=1,dive"`
}

// Validate valida las reglas de negocio de un curso individual
func (r *UpdateCourseRequest) Validate() error {
	// Validar que enddate sea mayor que startdate si ambos están presentes
	if r.EndDate != nil && r.StartDate != nil && *r.EndDate > 0 && *r.StartDate > 0 {
		if *r.EndDate <= *r.StartDate {
			return &ValidationError{
				Field:   "enddate",
				Message: "La fecha de fin debe ser posterior a la fecha de inicio",
			}
		}
	}

	// Validar que el formato esté dentro de los valores permitidos si está presente
	if r.Format != "" {
		validFormats := map[string]bool{"weeks": true, "topics": true, "social": true}
		if !validFormats[r.Format] {
			return &ValidationError{
				Field:   "format",
				Message: "Formato de curso inválido. Valores permitidos: weeks, topics, social",
			}
		}
	}

	return nil
}

// Validate valida reglas de negocio adicionales
func (r *UpdateCoursesRequest) Validate() error {
	// Validar que no se intente actualizar el curso site (ID=1)
	for i, course := range r.Courses {
		if course.ID == 1 {
			return &ValidationError{
				Field:   "courses",
				Message: "No se puede actualizar el curso site (ID=1)",
			}
		}

		// Validar cada curso individualmente
		if err := course.Validate(); err != nil {
			valErr := err.(*ValidationError)
			return &ValidationError{
				Field:   "courses[" + string(rune(i)) + "]." + valErr.Field,
				Message: valErr.Message,
			}
		}
	}
	return nil
}
