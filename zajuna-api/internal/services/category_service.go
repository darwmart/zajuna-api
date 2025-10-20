package services

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

// CategoryService maneja la lógica de categorías
type CategoryService struct {
	repo *repository.CategoryRepository
}

// NewCategoryService crea un nuevo servicio de categorías
func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetCategories devuelve todas las categorías
func (s *CategoryService) GetCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}
