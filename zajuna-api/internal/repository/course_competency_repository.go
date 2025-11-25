package repository

import (
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CourseCompetencyRepository struct {
	DB *gorm.DB
}

func NewCourseCompetencyRepository(db *gorm.DB) *CourseCompetencyRepository {
	return &CourseCompetencyRepository{DB: db}
}

func (r *CourseCompetencyRepository) AddCompetencyToCourse(courseCompetency *models.CourseCompetency) error {
	err := r.DB.Table("mdl_competency_coursecomp").Create(&courseCompetency).Error
	return err
}

func (r *CourseCompetencyRepository) Exists(courseID uint, competencyID uint) (bool, error) {
	var count int64
	err := r.DB.Table("mdl_competency_coursecomp").Where("courseid = ? AND competencyid = ?", courseID, competencyID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
