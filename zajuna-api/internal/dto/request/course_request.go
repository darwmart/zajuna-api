package request

// GetCoursesRequest representa los parámetros para listar cursos
type GetCoursesRequest struct {
	CategoryID int `form:"categoryid" binding:"omitempty,min=1"`
	Page       int `form:"page" binding:"omitempty,min=1"`
	Limit      int `form:"limit" binding:"omitempty,min=1,max=100"`
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
	ID int `uri:"id" binding:"required,min=1"`
}

// DeleteCoursesRequest representa la solicitud de eliminación de cursos
type DeleteCoursesRequest struct {
	CourseIDs []int `json:"courseids" binding:"required,min=1,dive,min=1"`
}

// Validate valida que los IDs sean únicos y no incluyan el curso site (ID=1)
func (r *DeleteCoursesRequest) Validate() error {
	seen := make(map[int]bool)
	for _, id := range r.CourseIDs {
		// Validar que no intenten eliminar el curso site (ID=1)
		if id == 1 {
			return &ValidationError{
				Field:   "courseids",
				Message: "No se puede eliminar el curso site (ID=1)",
			}
		}
		// Validar duplicados
		if seen[id] {
			return &ValidationError{
				Field:   "courseids",
				Message: "IDs duplicados detectados",
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

// UpdateCourseRequest representa la solicitud de actualización de un curso
type UpdateCourseRequest struct {
	ID            int    `json:"id" binding:"required,min=1"`
	CategoryID    *int   `json:"category" binding:"omitempty,min=1"`
	FullName      string `json:"fullname" binding:"omitempty,min=3,max=254"`
	ShortName     string `json:"shortname" binding:"omitempty,min=1,max=100"`
	IDNumber      string `json:"idnumber" binding:"omitempty,max=100"`
	Summary       string `json:"summary" binding:"omitempty"`
	SummaryFormat *int   `json:"summaryformat" binding:"omitempty,oneof=0 1 2 4"`
	Format        string `json:"format" binding:"omitempty,oneof=weeks topics social"`
	StartDate     *int64 `json:"startdate" binding:"omitempty,min=0"`
	EndDate       *int64 `json:"enddate" binding:"omitempty,min=0"`
	Visible       *int   `json:"visible" binding:"omitempty,oneof=0 1"`
}
