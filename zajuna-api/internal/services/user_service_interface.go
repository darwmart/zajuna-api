package services

import (
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
)

// UserServiceInterface define los metodos que debe implementar un servicio de usuarios
type UserServiceInterface interface {
	GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error)
	DeleteUsers(userIDs []int) error
	UpdateUsers(users []models.User) (int64, error)
	ToggleUserStatus(userID uint) (int, error)
	GetEnrolledUsers(courseID int, options map[string]interface{}) ([]response.EnrolledUserResponse, int, error)
}
