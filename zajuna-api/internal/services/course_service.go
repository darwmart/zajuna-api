package services

import (
	"fmt"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type CourseService struct {
	repo     repository.CourseRepositoryInterface
	permRepo *repository.PermissionRepository
}

func NewCourseService(repo repository.CourseRepositoryInterface) *CourseService {
	return &CourseService{repo: repo, permRepo: nil}
}

// SetPermissionRepository establece el repositorio de permisos (inyección de dependencia opcional)
func (s *CourseService) SetPermissionRepository(permRepo *repository.PermissionRepository) {
	s.permRepo = permRepo
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
		if course.SortOrder != nil {
			updates["sortorder"] = *course.SortOrder
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

// MoveCourses mueve múltiples cursos a nuevas categorías
// Compatible con core_course_move_courses de Moodle 4.3
func (s *CourseService) MoveCourses(courses []request.MoveCourseRequest) error {
	for _, course := range courses {
		if err := s.repo.MoveCourse(course.ID, course.CategoryID, course.BeforeID); err != nil {
			return err
		}
	}
	return nil
}

// GetCoursesWhereUserIsTeacher obtiene los cursos donde el usuario es instructor
// Busca cursos donde el usuario tiene rol de editingteacher o teacher
func (s *CourseService) GetCoursesWhereUserIsTeacher(userID int) ([]models.Course, error) {
	return s.repo.GetCoursesWhereUserIsTeacher(userID)
}

// GetCourseContent obtiene el contenido completo del curso (secciones y módulos)
// Compatible con core_course_get_contents de Moodle
// Filtra el contenido según los permisos del usuario
func (s *CourseService) GetCourseContent(courseID int, userID int) ([]repository.CourseSection, error) {
	// Obtener todas las secciones (sin filtrar)
	sections, err := s.repo.GetCourseContent(courseID)
	if err != nil {
		return nil, err
	}

	// Si no tenemos repositorio de permisos, devolver sin filtrar
	if s.permRepo == nil {
		return sections, nil
	}

	// Verificar si el usuario es site admin
	isSiteAdmin, _ := s.permRepo.IsSiteAdmin(userID)

	// Verificar si el usuario puede ver secciones ocultas
	canViewHiddenSections := false
	if !isSiteAdmin {
		// Obtener el contexto del curso
		courseContext, err := s.permRepo.GetContextByLevelAndInstance(50, courseID) // 50 = CONTEXT_COURSE
		if err == nil {
			permission, _ := s.permRepo.GetUserCapabilityInContext(userID, "moodle/course:viewhiddensections", courseContext.ID)
			canViewHiddenSections = (permission == 1) // CAP_ALLOW
		}
	} else {
		canViewHiddenSections = true
	}

	// Filtrar secciones basándose en permisos
	filteredSections := []repository.CourseSection{}
	for _, section := range sections {
		// Si la sección está oculta y el usuario no puede ver secciones ocultas, saltarla
		if section.Visible == 0 && !canViewHiddenSections {
			continue
		}

		// Filtrar módulos ocultos si el usuario no puede verlos
		if !canViewHiddenSections {
			filteredModules := []repository.CourseModule{}
			for _, module := range section.Modules {
				if module.Visible == 1 {
					filteredModules = append(filteredModules, module)
				}
			}
			section.Modules = filteredModules
		}

		// Actualizar UserVisible
		section.UserVisible = section.Visible == 1 || canViewHiddenSections

		filteredSections = append(filteredSections, section)
	}

	return filteredSections, nil
}
