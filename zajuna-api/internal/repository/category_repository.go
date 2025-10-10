package repository

import (
	"gorm.io/gorm"
	"zajunaApi/internal/models"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	if err := r.db.Table("mdl_course_categories").
		Order("sortorder").
		Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
