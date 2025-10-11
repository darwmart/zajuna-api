package services

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.repo.GetAllCourses()
}

func (s *CourseService) GetCoursesByCategory(categoryID uint) ([]models.Course, error) {
	return s.repo.GetCoursesByCategory(categoryID)
}
