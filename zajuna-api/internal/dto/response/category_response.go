package response

// CategoryResponse representa una categoría en la respuesta de la API
type CategoryResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	IDNumber          string `json:"idnumber"`
	Description       string `json:"description,omitempty"`
	DescriptionFormat int    `json:"descriptionformat"`
	Parent            int    `json:"parent"`
	SortOrder         int    `json:"sortorder"`
	CourseCount       int    `json:"coursecount"`
	Visible           int    `json:"visible"`
	Depth             int    `json:"depth"`
	Path              string `json:"path"`
	Theme             string `json:"theme,omitempty"`
}

// CategoryDetailResponse representa los detalles completos de una categoría
type CategoryDetailResponse struct {
	ID                uint     `json:"id"`
	Name              string   `json:"name"`
	IDNumber          string   `json:"idnumber"`
	Description       string   `json:"description,omitempty"`
	DescriptionFormat int      `json:"descriptionformat"`
	Parent            int      `json:"parent"`
	ParentName        string   `json:"parentname,omitempty"`
	SortOrder         int      `json:"sortorder"`
	CourseCount       int      `json:"coursecount"`
	Visible           int      `json:"visible"`
	Depth             int      `json:"depth"`
	Path              string   `json:"path"`
	Theme             string   `json:"theme,omitempty"`
	Subcategories     []string `json:"subcategories,omitempty"`
}

// CategoryListResponse representa la respuesta de listado de categorías
type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Pagination *PaginationMeta    `json:"pagination,omitempty"`
}

// CreateCategoryResponse representa la respuesta de creación de categoría
type CreateCategoryResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// UpdateCategoryResponse representa la respuesta de actualización de categoría
type UpdateCategoryResponse struct {
	Message string   `json:"message"`
	Updated bool     `json:"updated"`
	Errors  []string `json:"errors,omitempty"`
}

// DeleteCategoriesResponse representa la respuesta de eliminación de categorías
type DeleteCategoriesResponse struct {
	Message  string    `json:"message,omitempty"`
	Deleted  int       `json:"deleted"`
	Warnings []Warning `json:"warnings,omitempty"`
}

// CategoryTreeResponse representa una categoría en estructura de árbol
type CategoryTreeResponse struct {
	ID          uint                    `json:"id"`
	Name        string                  `json:"name"`
	CourseCount int                     `json:"coursecount"`
	Visible     int                     `json:"visible"`
	Depth       int                     `json:"depth"`
	Children    []CategoryTreeResponse  `json:"children,omitempty"`
}

// MoveCategoryResponse representa la respuesta de mover una categoría
type MoveCategoryResponse struct {
	Message    string `json:"message"`
	CategoryID uint   `json:"category_id"`
	NewParent  uint   `json:"new_parent"`
}
