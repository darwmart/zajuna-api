package services

import (
	"fmt"

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
	for i := range users {
		user := &users[i]
		userID := int(user.ID)

		// Obtener datos relacionados de cada usuario
		groups, _ := s.repo.GetUserGroupsInCourse(userID, courseID)
		roles, _ := s.repo.GetUserRolesInCourse(userID, courseID)
		customFields, _ := s.repo.GetUserCustomFields(userID)
		preferences, _ := s.repo.GetUserPreferences(userID)
		enrolledCourses, _ := s.repo.GetUserEnrolledCourses(userID)

		// Convertir a DTO usando el mapper
		userResp := mapper.EnrolledUserDetailToResponse(
			user,
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

// GetUserByID obtiene un usuario por su ID (solo usuarios activos y no eliminados)
func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.repo.FindByID(userID)
}

// CreateUsers crea múltiples usuarios (compatible con Moodle 4.3 core_user_create_users)
func (s *UserService) CreateUsers(usersData []struct {
	User     *models.User
	Password string
}) ([]response.CreateUserResponse, error) {
	var createdUsers []response.CreateUserResponse

	for _, userData := range usersData {
		// Validar que el username no exista
		existingUsers, _, err := s.repo.FindByFilters(map[string]string{"username": userData.User.Username}, 1, 1)
		if err == nil && len(existingUsers) > 0 {
			return nil, fmt.Errorf("username '%s' ya existe", userData.User.Username)
		}

		// Validar que el email no exista
		existingUsers, _, err = s.repo.FindByFilters(map[string]string{"email": userData.User.Email}, 1, 1)
		if err == nil && len(existingUsers) > 0 {
			return nil, fmt.Errorf("email '%s' ya está en uso", userData.User.Email)
		}

		// Crear el usuario
		err = s.repo.CreateUser(userData.User, userData.Password)
		if err != nil {
			return nil, fmt.Errorf("error al crear usuario '%s': %v", userData.User.Username, err)
		}

		createdUsers = append(createdUsers, response.CreateUserResponse{
			ID:       userData.User.ID,
			Username: userData.User.Username,
		})
	}

	return createdUsers, nil
}
