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

//func (s *CourseService) GetCourseRoles(courseID int) (map[string]int64, error) {
//return s.repo.GetRoleAssignments(courseID)}

func (s *CourseService) GetCourseDetails(courseID int) (*repository.CourseDetails, error) {
	return s.repo.GetCourseDetails(courseID)
}

func (s *CourseService) DeleteCourses(courseIDs []int) (*models.DeleteCoursesResponse, error) {
	warnings, err := s.repo.DeleteCourses(courseIDs)
	if err != nil {
		return nil, err
	}
	return &models.DeleteCoursesResponse{Warnings: warnings}, nil
}
