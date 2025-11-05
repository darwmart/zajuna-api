package services

import (
	"net/http"
	"zajunaApi/internal/models"
)

// UserServiceInterface define los metodos que debe implementar un servicio de usuarios
type UserServiceInterface interface {
	GetUsers(filters map[string]string, page, limit int) ([]models.User, int64, error)
	DeleteUsers(userIDs []int) error
	UpdateUsers(users []models.User) (int64, error)
	Login(r *http.Request, username, password string) (string, error)
	Logout(sid string) (string, error)
}
