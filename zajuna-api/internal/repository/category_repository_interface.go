package repository

import "zajunaApi/internal/models"

// CategoryRepositoryInterface define los métodos que debe implementar un repository de categorías
type CategoryRepositoryInterface interface {
	GetAllCategories() ([]models.Category, error)
}
