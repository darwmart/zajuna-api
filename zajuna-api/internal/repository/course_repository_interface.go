package repository

import "zajunaApi/internal/models"

// CourseRepositoryInterface define los métodos que debe implementar un repository de cursos
type CourseRepositoryInterface interface {
	GetAllCourses() ([]models.Course, error)
	GetCoursesByCategory(categoryID uint) ([]models.Course, error)
	GetCourseByID(id uint) (*models.Course, error)
	GetCourseByIDNumber(idnumber string) (*models.Course, error)
	GetCourseDetails(idnumber string) (*CourseDetails, error)
	DeleteCourses(courseIDs []int) ([]models.Warning, error)
	UpdateCourse(id int, updates map[string]interface{}) error
	SearchCourses(criteriaName, criteriaValue string, page, perPage int) ([]models.Course, int64, error)
}
