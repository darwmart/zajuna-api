package repository

import (
	"fmt"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// Obtener todos los cursos
func (r *CourseRepository) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course

	if err := r.db.Table("mdl_course").
		Order("fullname").
		Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

// Obtener cursos por categoría (útil para filtrar)
func (r *CourseRepository) GetCoursesByCategory(categoryID uint) ([]models.Course, error) {
	var courses []models.Course

	if err := r.db.Table("mdl_course").
		Where("category = ?", categoryID).
		Order("fullname").
		Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

// Obtener un curso por ID
func (r *CourseRepository) GetCourseByID(id uint) (*models.Course, error) {
	var course models.Course

	if err := r.db.Table("mdl_course").
		Where("id = ?", id).
		First(&course).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

// GetRoleAssignments obtiene el número de usuarios por rol en un curso.
func (r *CourseRepository) GetRoleAssignments(courseID int) (map[string]int64, error) {
	type Result struct {
		RoleName string
		Total    int64
	}

	var results []Result

	// Obtener el path del contexto del curso
	var courseContext struct {
		Path string
	}

	err := r.db.Table("mdl_context").
		Select("path").
		Where("instanceid = ? AND contextlevel = 50", courseID).
		Scan(&courseContext).Error

	if err != nil {
		return nil, fmt.Errorf("error buscando contexto del curso: %v", err)
	}
	if courseContext.Path == "" {
		return nil, fmt.Errorf("no se encontró el contexto del curso (id=%d)", courseID)
	}

	// Buscar todos los roles en ese contexto o en contextos “padres”
	err = r.db.Table("mdl_role_assignments AS ra").
		Select("r.shortname AS role_name, COUNT(ra.userid) AS total").
		Joins("JOIN mdl_context ctx ON ctx.id = ra.contextid").
		Joins("JOIN mdl_role r ON r.id = ra.roleid").
		Where("ctx.path LIKE ?", courseContext.Path+"%").
		Group("r.shortname").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	roleCount := make(map[string]int64)
	for _, r := range results {
		roleCount[r.RoleName] = r.Total
	}

	return roleCount, nil
}
