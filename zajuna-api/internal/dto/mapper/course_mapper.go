package mapper

import (
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

// CourseToResponse convierte un modelo Course a CourseResponse
func CourseToResponse(course *models.Course) *response.CourseResponse {
	if course == nil {
		return nil
	}

	return &response.CourseResponse{
		ID:                course.ID,
		Category:          course.Category,
		FullName:          course.FullName,
		ShortName:         course.ShortName,
		IDNumber:          course.IDNumber,
		Summary:           course.Summary,
		SummaryFormat:     course.SummaryFormat,
		Format:            course.Format,
		ShowGrades:        course.ShowGrades,
		NewsItems:         course.NewsItems,
		StartDate:         course.StartDate,
		EndDate:           course.EndDate,
		Visible:           course.Visible,
		GroupMode:         course.GroupMode,
		GroupModeForce:    course.GroupModeForce,
		DefaultGroupingID: course.DefaultGroupingID,
		SortOrder:         course.SortOrder,
		TimeCreated:       course.TimeCreated,
		TimeModified:      course.TimeModified,
	}
}

// CoursesToResponse convierte un slice de Courses a slice de CourseResponse
func CoursesToResponse(courses []models.Course) []response.CourseResponse {
	responses := make([]response.CourseResponse, len(courses))
	for i, course := range courses {
		resp := CourseToResponse(&course)
		if resp != nil {
			responses[i] = *resp
		}
	}
	return responses
}

// CourseDetailsToResponse convierte CourseDetails a CourseDetailResponse
func CourseDetailsToResponse(details *repository.CourseDetails) *response.CourseDetailResponse {
	if details == nil {
		return nil
	}

	return &response.CourseDetailResponse{
		ID:               details.ID,
		FullName:         details.FullName,
		ShortName:        details.ShortName,
		IDNumber:         details.IDNumber,
		Format:           details.Format,
		Category:         details.Category,
		Groupings:        details.Groupings,
		Groups:           details.Groups,
		RoleAssignments:  details.RoleAssignments,
		EnrollmentMethod: details.EnrollmentMethod,
		Sections:         details.Sections,
	}
}

// DeleteCoursesWarningsToResponse convierte warnings del modelo al DTO
func DeleteCoursesWarningsToResponse(warnings []models.Warning) []response.Warning {
	if warnings == nil {
		return nil
	}

	result := make([]response.Warning, len(warnings))
	for i, warning := range warnings {
		result[i] = response.Warning{
			Item:        warning.Item,
			ItemID:      warning.ItemID,
			WarningCode: warning.WarningCode,
			Message:     warning.Message,
		}
	}
	return result
}

// UpdateCoursesWarningsToResponse convierte warnings del modelo al DTO (reutiliza la misma función)
func UpdateCoursesWarningsToResponse(warnings []models.Warning) []response.Warning {
	return DeleteCoursesWarningsToResponse(warnings)
}
