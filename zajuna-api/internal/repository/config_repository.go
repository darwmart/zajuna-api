package repository

import (
	"errors"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type ConfigRepository struct {
	DB *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository {
	return &ConfigRepository{DB: db}
}

func (r *ConfigRepository) FindByName(name string) (*models.Config, error) {

	var config models.Config
	result := r.DB.First(&config, "name = ?", name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // no es un error real, solo que no existe el dato
		}
		return nil, result.Error // otro tipo de error (de conexión, SQL, etc.)
	}

	return &config, nil
}
