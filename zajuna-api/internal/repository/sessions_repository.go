package repository

import (
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	DB *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) *SessionsRepository {
	return &SessionsRepository{DB: db}
}

func (r *SessionsRepository) InsertSession(session *models.Sessions) error {

	result := r.DB.Create(session)

	return result.Error
}

func (r *SessionsRepository) DeleteSession(sid string) error {
	result := r.DB.Unscoped().Where("sid = ?", sid).Delete(&models.Sessions{})
	return result.Error
}
