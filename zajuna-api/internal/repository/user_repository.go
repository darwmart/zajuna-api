package repository

import (
	"errors"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FindByFilters busca usuarios con filtros dinámicos
func (r *UserRepository) FindByFilters(filters map[string]string, page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.DB.Model(&models.User{})

	// Aplicar filtros solo si vienen con datos
	if firstname := filters["firstname"]; firstname != "" {
		query = query.Where("firstname ILIKE ?", "%"+firstname+"%")
	}
	if lastname := filters["lastname"]; lastname != "" {
		query = query.Where("lastname ILIKE ?", "%"+lastname+"%")
	}
	if username := filters["username"]; username != "" {
		query = query.Where("username ILIKE ?", "%"+username+"%")
	}
	if email := filters["email"]; email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}

	// Contar total antes de paginar
	query.Count(&total)

	// Aplicar paginación
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// DeleteUsers suspende usuarios por sus IDs
func (r *UserRepository) DeleteUsers(userIDs []int) error {
	return r.DB.Table("mdl_user").
		Where("id IN ?", userIDs).
		Update("suspended", 1).Error
}

func (r *UserRepository) UpdateUsers(users []models.User) (int64, error) {
	var total int64
	for _, u := range users {
		if err := r.DB.Model(&models.User{}).
			Where("id = ?", u.ID).
			Updates(map[string]interface{}{
				"firstname": u.FirstName,
				"lastname":  u.LastName,
				"email":     u.Email,
				"city":      u.City,
				"country":   u.Country,
				"lang":      u.Lang,
				"timezone":  u.Timezone,
				"phone1":    u.Phone1,
			}).Error; err != nil {
			return total, err
		}
		total++
	}
	return total, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// No se encontró el usuario: no es un error para la capa superior
			return nil, nil
		}
		// Otro tipo de error (conexión, SQL, etc.)
		return nil, result.Error
	}

	return &user, nil
}
