package services

import (
	"zajunaApi/internal/dto/mapper"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	return s.repo.FindByFilters(filters, page, limit)
}

func (s *UserService) DeleteUsers(userIDs []int) error {
	return s.repo.DeleteUsers(userIDs)
}

func (s *UserService) UpdateUsers(users []models.User) (int64, error) {
	return s.repo.UpdateUsers(users)
}

func (s *UserService) ToggleUserStatus(userID uint) (int, error) {
	return s.repo.ToggleUserStatus(userID)
}

// GetEnrolledUsers obtiene usuarios matriculados en un curso con todas sus relaciones
func (s *UserService) GetEnrolledUsers(courseID int, options map[string]interface{}) ([]response.EnrolledUserResponse, int, error) {
	// Obtener usuarios matriculados del repository
	users, total, err := s.repo.GetEnrolledUsers(courseID, options)
	if err != nil {
		return nil, 0, err
	}

	// Construir respuestas con todos los datos relacionados
	responses := make([]response.EnrolledUserResponse, 0, len(users))
	for _, user := range users {
		userID := int(user.ID)

		// Obtener datos relacionados de cada usuario
		groups, _ := s.repo.GetUserGroupsInCourse(userID, courseID)
		roles, _ := s.repo.GetUserRolesInCourse(userID, courseID)
		customFields, _ := s.repo.GetUserCustomFields(userID)
		preferences, _ := s.repo.GetUserPreferences(userID)
		enrolledCourses, _ := s.repo.GetUserEnrolledCourses(userID)

		// Convertir a DTO usando el mapper
		userResp := mapper.EnrolledUserDetailToResponse(
			&user,
			groups,
			roles,
			customFields,
			preferences,
			enrolledCourses,
		)

		if userResp != nil {
			responses = append(responses, *userResp)
		}
	}

	return responses, total, nil
}
