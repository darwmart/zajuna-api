package repository

import "zajunaApi/internal/models"

type RoleCapabilityRepositoryInterface interface {
	FindByUserID(userID int64, roles []string, capability string) (*[]models.RoleCapability, error)
}
