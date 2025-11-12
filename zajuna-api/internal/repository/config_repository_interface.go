package repository

import "zajunaApi/internal/models"

type ConfigRepositoryInterface interface {
	FindByName(name string) (*models.Config, error)
}
