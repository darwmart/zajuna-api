package mapper

import (
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
)

// CategoryToResponse convierte un modelo Category a CategoryResponse
func CategoryToResponse(category *models.Category) *response.CategoryResponse {
	if category == nil {
		return nil
	}

	return &response.CategoryResponse{
		ID:                category.ID,
		Name:              category.Name,
		IDNumber:          category.IDNumber,
		Description:       category.Description,
		DescriptionFormat: category.DescriptionFormat,
		Parent:            category.Parent,
		SortOrder:         category.SortOrder,
		CourseCount:       category.CourseCount,
		Visible:           category.Visible,
		Depth:             category.Depth,
		Path:              category.Path,
		Theme:             category.Theme,
	}
}

// CategoriesToResponse convierte un slice de Categories a slice de CategoryResponse
func CategoriesToResponse(categories []models.Category) []response.CategoryResponse {
	responses := make([]response.CategoryResponse, len(categories))
	for i, category := range categories {
		resp := CategoryToResponse(&category)
		if resp != nil {
			responses[i] = *resp
		}
	}
	return responses
}
