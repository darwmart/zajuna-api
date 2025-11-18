package request

// GetCategoriesRequest representa los parámetros para listar categorías
type GetCategoriesRequest struct {
	Parent  *int `form:"parent" binding:"omitempty,min=0"`
	Visible *int `form:"visible" binding:"omitempty,oneof=0 1"`
	Page    int  `form:"page" binding:"omitempty,min=1"`
	Limit   int  `form:"limit" binding:"omitempty,min=1,max=100"`
}

// SetDefaults establece valores por defecto
func (r *GetCategoriesRequest) SetDefaults() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 50
	}
}

// HasParentFilter verifica si hay filtro de categoría padre
func (r *GetCategoriesRequest) HasParentFilter() bool {
	return r.Parent != nil
}

// GetCategoryDetailsRequest representa los parámetros para obtener detalles de una categoría
type GetCategoryDetailsRequest struct {
	ID int `uri:"id" binding:"required,min=1"`
}

// CreateCategoryRequest representa la solicitud de creación de una categoría
type CreateCategoryRequest struct {
	Name              string `json:"name" binding:"required,min=1,max=255"`
	Parent            int    `json:"parent" binding:"omitempty,min=0"`
	IDNumber          string `json:"idnumber" binding:"omitempty,max=100"`
	Description       string `json:"description" binding:"omitempty"`
	DescriptionFormat int    `json:"descriptionformat" binding:"omitempty,oneof=0 1 2 4"`
	Visible           int    `json:"visible" binding:"omitempty,oneof=0 1"`
	Theme             string `json:"theme" binding:"omitempty,max=50"`
}

// UpdateCategoryRequest representa la solicitud de actualización de una categoría
type UpdateCategoryRequest struct {
	ID                int    `json:"id" binding:"required,min=1"`
	Name              string `json:"name" binding:"omitempty,min=1,max=255"`
	Parent            *int   `json:"parent" binding:"omitempty,min=0"`
	IDNumber          string `json:"idnumber" binding:"omitempty,max=100"`
	Description       string `json:"description" binding:"omitempty"`
	DescriptionFormat *int   `json:"descriptionformat" binding:"omitempty,oneof=0 1 2 4"`
	Visible           *int   `json:"visible" binding:"omitempty,oneof=0 1"`
	Theme             string `json:"theme" binding:"omitempty,max=50"`
}

// DeleteCategoriesRequest representa la solicitud de eliminación de categorías
type DeleteCategoriesRequest struct {
	CategoryIDs []int `json:"categoryids" binding:"required,min=1,dive,min=1"`
	MoveToID    *int  `json:"movetoid" binding:"omitempty,min=1"`
}

// Validate valida que los IDs sean únicos
func (r *DeleteCategoriesRequest) Validate() error {
	seen := make(map[int]bool)
	for _, id := range r.CategoryIDs {
		if seen[id] {
			return &ValidationError{
				Field:   "categoryids",
				Message: "IDs duplicados detectados",
			}
		}
		seen[id] = true
	}

	// Validar que movetoid no esté en la lista de eliminación
	if r.MoveToID != nil {
		if seen[*r.MoveToID] {
			return &ValidationError{
				Field:   "movetoid",
				Message: "No se puede mover los cursos a una categoría que está siendo eliminada",
			}
		}
	}

	return nil
}

// MoveCategoryRequest representa la solicitud de mover una categoría
type MoveCategoryRequest struct {
	ID       uint  `json:"id" binding:"required,min=1"`
	BeforeID uint  `json:"beforeid"`
	ParentID *uint `json:"parentid"` // Opcional: nuevo padre (nil = mantener padre actual)
}

// Validate valida la solicitud de mover categoría
func (r *MoveCategoryRequest) Validate() error {
	// No se puede mover una categoría antes de sí misma
	if r.ID == r.BeforeID && r.BeforeID != 0 {
		return &ValidationError{
			Field:   "beforeid",
			Message: "No se puede mover una categoría antes de sí misma",
		}
	}

	// No se puede ser su propio padre
	if r.ParentID != nil && r.ID == *r.ParentID {
		return &ValidationError{
			Field:   "parentid",
			Message: "Una categoría no puede ser su propio padre",
		}
	}

	return nil
}
