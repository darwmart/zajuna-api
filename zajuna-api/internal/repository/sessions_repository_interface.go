package repository

import "zajunaApi/internal/models"

// SessionsRepositoryInterface define los métodos que debe implementar un repository de usuarios
type SessionsRepositoryInterface interface {
	InsertSession(session *models.Sessions) error
	DeleteSession(sid string) error
	FindBySID(sid string) (*models.Sessions, error)
}
