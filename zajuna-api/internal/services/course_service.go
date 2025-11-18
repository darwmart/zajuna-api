package services

import (
	"fmt"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type CourseService struct {
	repo repository.CourseRepositoryInterface
}

func NewCourseService(repo repository.CourseRepositoryInterface) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.repo.GetAllCourses()
}

func (s *CourseService) GetCoursesByCategory(categoryID uint) ([]models.Course, error) {
	return s.repo.GetCoursesByCategory(categoryID)
}

//func (s *CourseService) GetCourseRoles(courseID int) (map[string]int64, error) {
//return s.repo.GetRoleAssignments(courseID)}

func (s *CourseService) GetCourseDetails(idnumber string) (*repository.CourseDetails, error) {
	return s.repo.GetCourseDetails(idnumber)
}

func (s *CourseService) DeleteCourses(courseIDs []int) (*models.DeleteCoursesResponse, error) {
	warnings, err := s.repo.DeleteCourses(courseIDs)
	if err != nil {
		return nil, err
	}
	return &models.DeleteCoursesResponse{Warnings: warnings}, nil
}

// UpdateCourses actualiza múltiples cursos (compatible con Moodle)
func (s *CourseService) UpdateCourses(courses []request.UpdateCourseRequest) (*models.UpdateCoursesResponse, error) {
	var warnings []models.Warning

	for _, course := range courses {
		// Construir map de updates solo con campos no vacíos/nil
		updates := make(map[string]interface{})

		// Campos de texto
		if course.FullName != "" {
			updates["fullname"] = course.FullName
		}
		if course.ShortName != "" {
			updates["shortname"] = course.ShortName
		}
		if course.IDNumber != "" {
			updates["idnumber"] = course.IDNumber
		}
		if course.Summary != "" {
			updates["summary"] = course.Summary
		}
		if course.Format != "" {
			updates["format"] = course.Format
		}
		if course.Lang != "" {
			updates["lang"] = course.Lang
		}
		if course.ForceTheme != "" {
			updates["theme"] = course.ForceTheme
		}

		// Campos punteros (int)
		if course.CategoryID != nil {
			updates["category"] = *course.CategoryID
		}
		if course.SummaryFormat != nil {
			updates["summaryformat"] = *course.SummaryFormat
		}
		if course.ShowGrades != nil {
			updates["showgrades"] = *course.ShowGrades
		}
		if course.NewsItems != nil {
			updates["newsitems"] = *course.NewsItems
		}
		if course.NumSections != nil {
			updates["numsections"] = *course.NumSections
		}
		if course.MaxBytes != nil {
			updates["maxbytes"] = *course.MaxBytes
		}
		if course.ShowReports != nil {
			updates["showreports"] = *course.ShowReports
		}
		if course.Visible != nil {
			updates["visible"] = *course.Visible
		}
		if course.HiddenSections != nil {
			updates["hiddensections"] = *course.HiddenSections
		}
		if course.GroupMode != nil {
			updates["groupmode"] = *course.GroupMode
		}
		if course.GroupModeForce != nil {
			updates["groupmodeforce"] = *course.GroupModeForce
		}
		if course.DefaultGroupingID != nil {
			updates["defaultgroupingid"] = *course.DefaultGroupingID
		}
		if course.EnableCompletion != nil {
			updates["enablecompletion"] = *course.EnableCompletion
		}
		if course.CompletionNotify != nil {
			updates["completionnotify"] = *course.CompletionNotify
		}

		// Campos punteros (int64)
		if course.StartDate != nil {
			updates["startdate"] = *course.StartDate
		}
		if course.EndDate != nil {
			updates["enddate"] = *course.EndDate
		}

		// Si no hay campos para actualizar, agregar warning
		if len(updates) == 0 {
			warnings = append(warnings, models.Warning{
				Item:        "course",
				ItemID:      course.ID,
				WarningCode: "nofieldstoupdate",
				Message:     fmt.Sprintf("No fields provided to update course %d", course.ID),
			})
			continue
		}

		// Intentar actualizar el curso
		err := s.repo.UpdateCourse(course.ID, updates)
		if err != nil {
			warnings = append(warnings, models.Warning{
				Item:        "course",
				ItemID:      course.ID,
				WarningCode: "updatefailed",
				Message:     fmt.Sprintf("Failed to update course %d: %s", course.ID, err.Error()),
			})
			continue
		}

		// Procesar courseformatoptions si están presentes
		if len(course.CourseFormatOptions) > 0 {
			if err := s.repo.UpdateCourseFormatOptions(course.ID, course.CourseFormatOptions); err != nil {
				warnings = append(warnings, models.Warning{
					Item:        "course",
					ItemID:      course.ID,
					WarningCode: "formatoptionsfailed",
					Message:     fmt.Sprintf("Failed to update format options for course %d: %s", course.ID, err.Error()),
				})
			}
		}

		// Procesar customfields si están presentes
		if len(course.CustomFields) > 0 {
			if err := s.repo.UpdateCourseCustomFields(course.ID, course.CustomFields); err != nil {
				warnings = append(warnings, models.Warning{
					Item:        "course",
					ItemID:      course.ID,
					WarningCode: "customfieldsfailed",
					Message:     fmt.Sprintf("Failed to update custom fields for course %d: %s", course.ID, err.Error()),
				})
			}
		}
	}

	return &models.UpdateCoursesResponse{Warnings: warnings}, nil
}

// SearchCourses busca cursos según criterio y valor
func (s *CourseService) SearchCourses(criteriaName, criteriaValue string, page, perPage int) ([]models.Course, int64, error) {
	return s.repo.SearchCourses(criteriaName, criteriaValue, page, perPage)
}

func (s *CourseService) AddCompetencyToCourse() {

}
