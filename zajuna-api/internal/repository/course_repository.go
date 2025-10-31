package repository

import (
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

type CourseDetails struct {
	ID               int64            `json:"id"`
	FullName         string           `json:"fullName"`
	ShortName        string           `json:"shortName"`
	IDNumber         string           `json:"idNumber"`
	Format           string           `json:"format"`
	Category         string           `json:"category"`
	Groupings        int64            `json:"groupings"`
	Groups           int64            `json:"groups"`
	RoleAssignments  map[string]int64 `json:"roleAssignments" gorm:"-"`
	EnrollmentMethod string           `json:"enrollmentMethod"`
	Sections         []string         `json:"sections" gorm:"-"`
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
func (r *CourseRepository) GetCourseDetails(courseID int) (*CourseDetails, error) {
	// Reutilizamos GetCourseByID para obtener la información base
	course, err := r.GetCourseByID(uint(courseID))
	if err != nil {
		return nil, err
	}

	// Creamos el struct base de detalles
	details := CourseDetails{
		ID:        int64(course.ID),
		FullName:  course.FullName,
		ShortName: course.ShortName,
		IDNumber:  course.IDNumber,
		Format:    course.Format,
	}

	// Obtener categoría
	r.db.Table("mdl_course_categories").
		Select("name").
		Where("id = ?", course.Category).
		Scan(&details.Category)

	// Contar agrupamientos
	r.db.Table("mdl_groupings").Where("courseid = ?", courseID).Count(&details.Groupings)

	// Contar grupos
	r.db.Table("mdl_groups").Where("courseid = ?", courseID).Count(&details.Groups)

	// Asignaciones de roles
	roleAssignments := map[string]int64{}
	rows, _ := r.db.Table("mdl_role_assignments ra").
		Select("r.shortname, COUNT(ra.id) as total").
		Joins("JOIN mdl_context ctx ON ra.contextid = ctx.id").
		Joins("JOIN mdl_role r ON r.id = ra.roleid").
		Where("ctx.contextlevel = 50 AND ctx.instanceid = ?", courseID).
		Group("r.shortname").
		Rows()

	defer rows.Close()
	for rows.Next() {
		var role string
		var total int64
		rows.Scan(&role, &total)
		roleAssignments[role] = total
	}
	details.RoleAssignments = roleAssignments

	// Métodos de matriculación
	var enrolMethods []string
	r.db.Table("mdl_enrol").
		Select("DISTINCT enrol").
		Where("courseid = ?", courseID).
		Scan(&enrolMethods)
	if len(enrolMethods) == 0 {
		enrolMethods = []string{"Matriculación manual"}
	}
	details.EnrollmentMethod = enrolMethods[0]

	// Secciones
	var sectionNames []string
	r.db.Table("mdl_course_sections").
		Select("name").
		Where("course = ? AND name IS NOT NULL AND name != ''", courseID).
		Order("section ASC").
		Scan(&sectionNames)
	details.Sections = sectionNames

	return &details, nil
}

func (r *CourseRepository) DeleteCourses(courseIDs []int) ([]models.Warning, error) {
	var warnings []models.Warning

	// Validar si los cursos existen
	var existingIDs []int
	if err := r.db.Table("mdl_course").
		Where("id IN ?", courseIDs).
		Pluck("id", &existingIDs).Error; err != nil {
		return nil, err
	}

	// Generar warnings para los IDs inexistentes
	for _, id := range courseIDs {
		found := false
		for _, existing := range existingIDs {
			if id == existing {
				found = true
				break
			}
		}
		if !found {
			warnings = append(warnings, models.Warning{
				Item:        "course",
				ItemID:      id,
				WarningCode: "invalidcourseid",
				Message:     "Course ID not found in the database",
			})
		}
	}

	// Eliminar físicamente los cursos válidos
	if len(existingIDs) > 0 {
		if err := r.db.Table("mdl_course").
			Where("id IN ?", existingIDs).
			Delete(nil).Error; err != nil {
			return warnings, err
		}
	}

	return warnings, nil
}

// UpdateCourse actualiza un curso en la base de datos
func (r *CourseRepository) UpdateCourse(id int, updates map[string]interface{}) error {
	// Siempre actualizar timemodified
	updates["timemodified"] = gorm.Expr("EXTRACT(EPOCH FROM NOW())::INTEGER")

	// Ejecutar la actualización
	result := r.db.Table("mdl_course").
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	// Verificar que se actualizó al menos una fila
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
