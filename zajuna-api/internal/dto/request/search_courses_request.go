package request

import "errors"

// SearchCoursesRequest representa la petición de búsqueda de cursos
type SearchCoursesRequest struct {
	CriteriaName  string `form:"criterianame" binding:"required"` // search, categoryid, etc.
	CriteriaValue string `form:"criteriavalue" binding:"required"` // valor a buscar
	Page          int    `form:"page"`                            // página (base 0)
	PerPage       int    `form:"perpage"`                         // elementos por página (0 = todos)
}

// SetDefaults establece valores por defecto
func (r *SearchCoursesRequest) SetDefaults() {
	if r.Page < 0 {
		r.Page = 0
	}
	if r.PerPage < 0 {
		r.PerPage = 0
	}
}

// Validate valida los parámetros de búsqueda
func (r *SearchCoursesRequest) Validate() error {
	if r.CriteriaName == "" {
		return errors.New("criterianame is required")
	}
	if r.CriteriaValue == "" {
		return errors.New("criteriavalue is required")
	}

	// Validar criterios soportados
	validCriteria := map[string]bool{
		"search":     true,
		"categoryid": true,
		"id":         true,
		"idnumber":   true,
	}

	if !validCriteria[r.CriteriaName] {
		return errors.New("invalid criterianame: supported values are search, categoryid, id, idnumber")
	}

	return nil
}
