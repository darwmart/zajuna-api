package services

import "zajunaApi/internal/models"

// CategoryServiceInterface define los metodos que debe implementar un servicio de categorias
type CategoryServiceInterface interface {
	GetCategories() ([]models.Category, error)
}
