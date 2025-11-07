package response

// SearchCoursesResponse representa la respuesta de búsqueda de cursos
type SearchCoursesResponse struct {
	Total      int              `json:"total"`
	Courses    []CourseResponse `json:"courses"`
	Pagination *PaginationMeta  `json:"pagination,omitempty"`
}
