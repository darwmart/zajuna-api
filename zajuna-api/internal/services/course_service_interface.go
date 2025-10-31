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
	GetCourseDetails(courseID int) (*repository.CourseDetails, error)
	DeleteCourses(courseIDs []int) (*models.DeleteCoursesResponse, error)
	UpdateCourses(courses []request.UpdateCourseRequest) (*models.UpdateCoursesResponse, error)
}
