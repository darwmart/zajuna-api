package services

import (
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
)

// CategoryService maneja la lógica de categorías
type CategoryService struct {
	repo repository.CategoryRepositoryInterface
}

// NewCategoryService crea un nuevo servicio de categorías
func NewCategoryService(repo repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetCategories devuelve todas las categorías
func (s *CategoryService) GetCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

// MoveCategory mueve una categoría antes de otra y/o cambia su padre
func (s *CategoryService) MoveCategory(id uint, beforeid uint, parentid *uint) error {
	return s.repo.MoveCategory(id, beforeid, parentid)
}
