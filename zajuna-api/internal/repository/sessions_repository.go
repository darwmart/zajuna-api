package repository

import (
	"errors"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	DB *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) *SessionsRepository {
	return &SessionsRepository{DB: db}
}

func (r *SessionsRepository) InsertSession(session models.Sessions) error {

	result := r.DB.Create(session)

	return result.Error
}

func (r *SessionsRepository) DeleteSession(sid string) error {
	result := r.DB.Unscoped().Where("sid = ?", sid).Delete(&models.Sessions{})
	return result.Error
}

func (r *SessionsRepository) FindBySID(sid string) (*models.Sessions, error) {

	var session models.Sessions
	result := r.DB.First(&session, "sid = ?", sid)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // no es un error real, solo que no existe el usuario
		}
		return nil, result.Error // otro tipo de error (de conexión, SQL, etc.)
	}

	return &session, nil
}
