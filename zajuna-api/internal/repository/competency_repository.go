package repository

import (
	"errors"
	"fmt"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CompetencyRepository struct {
	DB *gorm.DB
}

func NewCompetencyRepository(db *gorm.DB) *CompetencyRepository {
	return &CompetencyRepository{DB: db}
}

func (r *CompetencyRepository) Create(competency *models.Competency) (*models.Competency, error) {
	var existing models.Competency
	err := r.DB.Table("mdl_competency").Where("idnumber = ?", competency.IDNumber).First(&existing).Error

	if err == nil {
		// Ya existe → error de negocio
		return nil, fmt.Errorf("idnumber '%s' ya está en uso", competency.IDNumber)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Error inesperado (bd caída, query inválida, etc.)
		return nil, err
	}

	err = r.DB.Table("mdl_competency").Create(competency).Error
	return competency, err
}

func (r *CompetencyRepository) Update(competency *models.Competency) (*models.Competency, error) {
	err := r.DB.Table("mdl_competency").Save(competency).Error
	return competency, err
}

func (r *CompetencyRepository) FindByID(id uint) (*models.Competency, error) {
	var competency models.Competency
	err := r.DB.Table("mdl_competency").First(&competency, id).Error
	if err != nil {
		return nil, err
	}
	return &competency, nil
}
