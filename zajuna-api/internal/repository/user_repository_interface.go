package repository

import "zajunaApi/internal/models"

// UserRepositoryInterface define los métodos que debe implementar un repository de usuarios
type UserRepositoryInterface interface {
	FindByFilters(filters map[string]string, page, limit int) ([]models.User, int64, error)
	DeleteUsers(userIDs []int) error
	UpdateUsers(users []models.User) (int64, error)
	FindByUsername(username string) (*models.User, error)
}
