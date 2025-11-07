package services

import (
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

// CourseServiceInterface define los métodos que debe implementar un servicio de cursos
type CourseServiceInterface interface {
	GetAllCourses() ([]models.Course, error)
	GetCoursesByCategory(categoryID uint) ([]models.Course, error)
	GetCourseDetails(idnumber string) (*repository.CourseDetails, error)
	DeleteCourses(courseIDs []int) (*models.DeleteCoursesResponse, error)
	UpdateCourses(courses []request.UpdateCourseRequest) (*models.UpdateCoursesResponse, error)
	SearchCourses(criteriaName, criteriaValue string, page, perPage int) ([]models.Course, int64, error)
}
