package repository

import (
	"errors"
	"fmt"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CompetencyFrameworkRepository struct {
	DB *gorm.DB
}

func NewCompetencyFrameworkRepository(db *gorm.DB) *CompetencyFrameworkRepository {
	return &CompetencyFrameworkRepository{DB: db}
}
func (r *CompetencyFrameworkRepository) Create(competency *models.CompetencyFramework) (*models.CompetencyFramework, error) {
	var existing models.Competency
	err := r.DB.Table("mdl_competency_framework").Where("idnumber = ?", competency.IDNumber).First(&existing).Error

	if err == nil {
		// Ya existe → error de negocio
		return nil, fmt.Errorf("idnumber '%s' ya está en uso", competency.IDNumber)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Error inesperado (bd caída, query inválida, etc.)
		return nil, err
	}

	err = r.DB.Table("mdl_competency_framework").Create(competency).Error
	return competency, err
}
func (r *CompetencyFrameworkRepository) FindByID(id uint) (*models.CompetencyFramework, error) {
	var competencyFramework models.CompetencyFramework
	result := r.DB.Table("mdl_competency_framework").First(&competencyFramework, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // no es un error real, solo que no existe el usuario
		}
		return nil, result.Error // otro tipo de error (de conexión, SQL, etc.)
	}
	return &competencyFramework, nil
}
