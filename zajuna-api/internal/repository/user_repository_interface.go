package repository

import "zajunaApi/internal/models"

// UserRepositoryInterface define los métodos que debe implementar un repository de usuarios
type UserRepositoryInterface interface {
	FindByFilters(filters map[string]string, page, limit int) ([]models.User, int64, error)
	DeleteUsers(userIDs []int) error
	UpdateUsers(users []models.User) (int64, error)
	ToggleUserStatus(userID uint) (int, error)
	GetEnrolledUsers(courseID int, options map[string]interface{}) ([]EnrolledUserDetail, int, error)
	GetUserGroupsInCourse(userID, courseID int) ([]map[string]interface{}, error)
	GetUserRolesInCourse(userID, courseID int) ([]map[string]interface{}, error)
	GetUserCustomFields(userID int) ([]map[string]interface{}, error)
	GetUserPreferences(userID int) ([]map[string]interface{}, error)
	GetUserEnrolledCourses(userID int) ([]map[string]interface{}, error)
	FindByUsername(username string) (*models.User, error)
}
