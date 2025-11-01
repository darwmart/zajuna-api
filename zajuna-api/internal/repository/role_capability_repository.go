package repository

import (
	"errors"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type RoleCapabilityRepository struct {
	DB *gorm.DB
}

func NewRoleCapabilityRepository(db *gorm.DB) *RoleCapabilityRepository {
	return &RoleCapabilityRepository{DB: db}
}

func (r *RoleCapabilityRepository) FindByUserID(userID int64, roles []string, capability string) (*[]models.RoleCapability, error) {

	var roleCapability []models.RoleCapability

	subquery := r.DB.Table("mdl_role_assignments").
		Select("DISTINCT roleid").
		Where("userid = ?", userID).Or("roleid IN(?)", roles)

	result := r.DB.Debug().Table("mdl_role_capabilities").
		Where("roleid IN (?)", subquery).
		Where("capability = ?", capability).
		Find(&roleCapability)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // no es un error real, solo que no existe el usuario
		}
		return nil, result.Error // otro tipo de error (de conexión, SQL, etc.)
	}

	return &roleCapability, nil
}
