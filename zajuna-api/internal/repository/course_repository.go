package repository

import (
	"zajunaApi/internal/dto/request"
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

// Obtener un curso por IDNumber
func (r *CourseRepository) GetCourseByIDNumber(idnumber string) (*models.Course, error) {
	var course models.Course

	if err := r.db.Table("mdl_course").
		Where("idnumber = ?", idnumber).
		First(&course).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

// GetRoleAssignments obtiene el número de usuarios por rol en un curso.
func (r *CourseRepository) GetCourseDetails(idnumber string) (*CourseDetails, error) {
	// Buscar curso por idnumber
	course, err := r.GetCourseByIDNumber(idnumber)
	if err != nil {
		return nil, err
	}

	courseID := int(course.ID)

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

// SearchCourses busca cursos según criterio y valor
func (r *CourseRepository) SearchCourses(criteriaName, criteriaValue string, page, perPage int) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64

	query := r.db.Table("mdl_course")

	// Aplicar el criterio de búsqueda
	switch criteriaName {
	case "search":
		// Búsqueda por texto en fullname, shortname o idnumber
		searchPattern := "%" + criteriaValue + "%"
		query = query.Where("fullname ILIKE ? OR shortname ILIKE ? OR idnumber ILIKE ?",
			searchPattern, searchPattern, searchPattern)

	case "categoryid":
		// Búsqueda por categoría
		query = query.Where("category = ?", criteriaValue)

	case "id":
		// Búsqueda por ID
		query = query.Where("id = ?", criteriaValue)

	case "idnumber":
		// Búsqueda por idnumber
		query = query.Where("idnumber = ?", criteriaValue)

	default:
		return nil, 0, nil
	}

	// Contar total
	query.Count(&total)

	// Ordenar alfabéticamente
	query = query.Order("fullname ASC")

	// Aplicar paginación si perPage > 0
	if perPage > 0 {
		offset := page * perPage
		query = query.Offset(offset).Limit(perPage)
	}

	// Ejecutar consulta
	if err := query.Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	return courses, total, nil
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

// UpdateCourseFormatOptions actualiza o inserta opciones de formato de curso
func (r *CourseRepository) UpdateCourseFormatOptions(courseID int, options []request.CourseFormatOption) error {
	// Primero verificar que el curso existe
	var exists bool
	if err := r.db.Table("mdl_course").
		Select("1").
		Where("id = ?", courseID).
		Limit(1).
		Find(&exists).Error; err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	// Obtener el formato del curso
	var format string
	if err := r.db.Table("mdl_course").
		Select("format").
		Where("id = ?", courseID).
		Scan(&format).Error; err != nil {
		return err
	}

	// Procesar cada opción de formato
	for _, option := range options {
		// Verificar si la opción ya existe
		var existingID int
		err := r.db.Table("mdl_course_format_options").
			Select("id").
			Where("courseid = ? AND format = ? AND name = ?", courseID, format, option.Name).
			Scan(&existingID).Error

		if err == gorm.ErrRecordNotFound || existingID == 0 {
			// Insertar nueva opción
			if err := r.db.Exec(`
				INSERT INTO mdl_course_format_options (courseid, format, sectionid, name, value)
				VALUES (?, ?, 0, ?, ?)
			`, courseID, format, option.Name, option.Value).Error; err != nil {
				return err
			}
		} else {
			// Actualizar opción existente
			if err := r.db.Table("mdl_course_format_options").
				Where("id = ?", existingID).
				Update("value", option.Value).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// UpdateCourseCustomFields actualiza o inserta campos personalizados de curso
func (r *CourseRepository) UpdateCourseCustomFields(courseID int, fields []request.CustomField) error {
	// Primero verificar que el curso existe
	var exists bool
	if err := r.db.Table("mdl_course").
		Select("1").
		Where("id = ?", courseID).
		Limit(1).
		Find(&exists).Error; err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	// Obtener el context ID del curso
	var contextID int
	if err := r.db.Table("mdl_context").
		Select("id").
		Where("contextlevel = 50 AND instanceid = ?", courseID).
		Scan(&contextID).Error; err != nil {
		return err
	}

	// Procesar cada campo personalizado
	for _, field := range fields {
		// Buscar el ID del campo personalizado por shortname
		var fieldID int
		err := r.db.Table("mdl_customfield_field").
			Select("id").
			Where("shortname = ?", field.ShortName).
			Scan(&fieldID).Error

		if err == gorm.ErrRecordNotFound || fieldID == 0 {
			// El campo personalizado no existe, continuar con el siguiente
			continue
		}

		// Verificar si ya existe un valor para este campo y curso
		var existingDataID int
		err = r.db.Table("mdl_customfield_data").
			Select("id").
			Where("fieldid = ? AND instanceid = ? AND contextid = ?", fieldID, courseID, contextID).
			Scan(&existingDataID).Error

		if err == gorm.ErrRecordNotFound || existingDataID == 0 {
			// Insertar nuevo valor
			if err := r.db.Exec(`
				INSERT INTO mdl_customfield_data (fieldid, instanceid, contextid, value, valueformat, timecreated, timemodified)
				VALUES (?, ?, ?, ?, 0, EXTRACT(EPOCH FROM NOW())::INTEGER, EXTRACT(EPOCH FROM NOW())::INTEGER)
			`, fieldID, courseID, contextID, field.Value).Error; err != nil {
				return err
			}
		} else {
			// Actualizar valor existente
			if err := r.db.Exec(`
				UPDATE mdl_customfield_data
				SET value = ?, timemodified = EXTRACT(EPOCH FROM NOW())::INTEGER
				WHERE id = ?
			`, field.Value, existingDataID).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
