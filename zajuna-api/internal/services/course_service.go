package services

import "zajunaApi/internal/repository"

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetCoursesByCategory(categoryID int) ([]repository.Course, error) {
	return s.repo.GetCoursesByCategory(categoryID)
}
