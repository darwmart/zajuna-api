package repository

import (
	"fmt"
	"time"
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
		Order("sortorder ASC").
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
		Order("sortorder ASC").
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

// GetCourseDetailsByID obtiene los detalles completos de un curso por su ID
func (r *CourseRepository) GetCourseDetailsByID(id int) (*CourseDetails, error) {
	// Buscar curso por ID
	course, err := r.GetCourseByID(uint(id))
	if err != nil {
		return nil, err
	}

	return r.buildCourseDetails(course)
}

// GetRoleAssignments obtiene el número de usuarios por rol en un curso.
func (r *CourseRepository) GetCourseDetails(idnumber string) (*CourseDetails, error) {
	// Buscar curso por idnumber
	course, err := r.GetCourseByIDNumber(idnumber)
	if err != nil {
		return nil, err
	}

	return r.buildCourseDetails(course)
}

// buildCourseDetails construye los detalles completos de un curso
func (r *CourseRepository) buildCourseDetails(course *models.Course) (*CourseDetails, error) {
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

	// Ordenar por sortorder (orden definido en Moodle)
	query = query.Order("sortorder ASC")

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

// MoveCourse mueve un curso a una nueva categoría y opcionalmente lo posiciona antes de otro curso
func (r *CourseRepository) MoveCourse(id int, categoryID int, beforeID *int) error {
	// 1. Verificar que el curso a mover existe
	var courseToMove models.Course
	if err := r.db.Table("mdl_course").
		Where("id = ?", id).
		First(&courseToMove).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// No permitir mover el curso site (ID=1)
	if id == 1 {
		return gorm.ErrInvalidValue
	}

	// 2. Verificar que la categoría destino existe (si no es 0)
	if categoryID > 0 {
		var category models.Category
		if err := r.db.Table("mdl_course_categories").
			Where("id = ?", categoryID).
			First(&category).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return gorm.ErrRecordNotFound
			}
			return err
		}
	}

	// 3. Si beforeID está especificado, verificar que existe y está en la misma categoría destino
	if beforeID != nil && *beforeID > 0 {
		var targetCourse models.Course
		if err := r.db.Table("mdl_course").
			Where("id = ?", *beforeID).
			First(&targetCourse).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return gorm.ErrRecordNotFound
			}
			return err
		}

		// Verificar que el curso objetivo está en la categoría destino
		if targetCourse.Category != categoryID {
			return gorm.ErrInvalidValue
		}
	}

	// 4. Obtener todos los cursos de la categoría destino, ordenados por sortorder
	var siblings []models.Course
	if err := r.db.Table("mdl_course").
		Where("category = ?", categoryID).
		Order("sortorder ASC, id ASC").
		Find(&siblings).Error; err != nil {
		return err
	}

	// 5. Construir la nueva lista de cursos reordenados
	var reordered []models.Course

	if beforeID == nil || *beforeID == 0 {
		// Mover al final
		for _, course := range siblings {
			if course.ID != id {
				reordered = append(reordered, course)
			}
		}
		reordered = append(reordered, courseToMove)
	} else {
		// Encontrar la posición actual del curso a mover
		var currentIndex int
		for i, course := range siblings {
			if course.ID == id {
				currentIndex = i
				break
			}
		}

		// Encontrar la posición del beforeID
		var beforeIndex int
		for i, course := range siblings {
			if course.ID == *beforeID {
				beforeIndex = i
				break
			}
		}

		// Determinar el curso adyacente con el que intercambiar
		var adjacentIndex int
		if beforeIndex < currentIndex {
			// Mover hacia arriba: intercambiar con el anterior (currentIndex - 1)
			adjacentIndex = currentIndex - 1
		} else {
			// Mover hacia abajo: intercambiar con el siguiente (currentIndex + 1)
			adjacentIndex = currentIndex + 1
		}

		// Construir el nuevo array intercambiando posiciones
		for i, course := range siblings {
			if i == currentIndex {
				// En la posición del curso a mover, poner el adyacente
				reordered = append(reordered, siblings[adjacentIndex])
			} else if i == adjacentIndex {
				// En la posición del adyacente, poner el curso a mover
				reordered = append(reordered, courseToMove)
			} else {
				// Los demás se mantienen igual
				reordered = append(reordered, course)
			}
		}
	}

	// 6. Si cambió de categoría, actualizar el campo category
	if courseToMove.Category != categoryID {
		if err := r.db.Table("mdl_course").
			Where("id = ?", id).
			Update("category", categoryID).Error; err != nil {
			return err
		}
	}

	// 7. Obtener el sortorder de la categoría destino
	var category models.Category
	if err := r.db.Table("mdl_course_categories").
		Where("id = ?", categoryID).
		First(&category).Error; err != nil {
		return err
	}

	// 8. Actualizar sortorder secuencialmente para todos los cursos de la categoría destino
	// usando la misma lógica que las categorías:
	// sortorder = sortorder_categoria + posición (autoincremental por unidad)
	for i, course := range reordered {
		newSortOrder := category.SortOrder + (i + 1)
		if err := r.db.Table("mdl_course").
			Where("id = ?", course.ID).
			Update("sortorder", newSortOrder).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetCoursesWhereUserIsTeacher obtiene TODOS los cursos donde el usuario está matriculado
// Esto replica el comportamiento de Moodle: "Mis cursos" muestra todos los cursos donde el usuario
// tiene una matrícula activa, independientemente del rol (estudiante, profesor, etc.)
// Compatible con core_enrol_get_users_courses de Moodle
func (r *CourseRepository) GetCoursesWhereUserIsTeacher(userID int) ([]models.Course, error) {
	var courses []models.Course

	// Query para obtener cursos donde el usuario está matriculado
	// Replica el comportamiento de Moodle usando mdl_enrol y mdl_user_enrolments
	// Busca en mdl_user_enrolments donde:
	// 1. userid = userID
	// 2. status = 0 (matrícula activa)
	// 3. La matrícula no ha expirado (timestart y timeend)
	// 4. El curso es visible
	currentTime := time.Now().Unix()

	err := r.db.Table("mdl_course AS c").
		Select("DISTINCT c.*").
		Joins("JOIN mdl_enrol AS e ON e.courseid = c.id").
		Joins("JOIN mdl_user_enrolments AS ue ON ue.enrolid = e.id").
		Where("ue.userid = ?", userID).
		Where("ue.status = 0"). // 0 = matrícula activa
		Where("(ue.timestart = 0 OR ue.timestart <= ?)", currentTime).
		Where("(ue.timeend = 0 OR ue.timeend >= ?)", currentTime).
		Where("c.visible = 1"). // Solo cursos visibles
		Order("c.sortorder ASC").
		Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}

// ============================================================================
// STRUCTS - Compatible 100% con core_course_get_contents de Moodle 4.x
// ============================================================================

// CourseModule representa un módulo/actividad dentro de una sección
// Compatible con la estructura devuelta por core_course_get_contents
type CourseModule struct {
	ID        int    `json:"id"`        // cm.id - ID del course_module
	Module    int    `json:"module"`    // cm.module - ID del tipo de módulo
	Instance  int    `json:"instance"`  // cm.instance - ID de la instancia específica
	Name      string `json:"name"`      // Nombre del módulo (desde mdl_assign, mdl_quiz, etc.)
	ModName   string `json:"modname"`   // Tipo: assign, quiz, resource, forum, etc.
	ModIcon   string `json:"modicon"`   // URL del icono (opcional)
	ModPlural string `json:"modplural"` // Nombre plural del módulo (opcional)

	Indent      int  `json:"indent"`      // Nivel de indentación
	Visible     int  `json:"visible"`     // 0 = oculto, 1 = visible
	UserVisible bool `json:"uservisible"` // Si el usuario puede verlo

	URL string `json:"url"` // URL del módulo: /mod/{modname}/view.php?id={cmid}

	Completion int `json:"completion"` // Estado de completitud (0, 1, 2)

	// Campos opcionales que Moodle puede devolver
	VisibleOnCoursePage int    `json:"visibleoncoursepage"` // Si es visible en la página del curso
	AvailabilityInfo    string `json:"availabilityinfo"`    // Info de restricciones
}

// CourseSection representa una sección del curso con sus módulos
// Compatible 100% con core_course_get_contents de Moodle
type CourseSection struct {
	// Campos básicos de mdl_course_sections
	ID            int    `json:"id"`            // ID de la sección
	Section       int    `json:"section"`       // Número de sección (0, 1, 2...)
	Name          string `json:"name"`          // Nombre de la sección
	Summary       string `json:"summary"`       // Resumen HTML
	SummaryFormat int    `json:"summaryformat"` // Formato (1 = HTML)

	// Visibilidad
	Visible     int  `json:"visible"`     // 0 = oculta, 1 = visible
	UserVisible bool `json:"uservisible"` // Calculado: si el usuario puede verla

	// Jerarquía (Formato Flexsections)
	Parent *int `json:"parent"` // ID de la sección padre (null = raíz)

	// Restricciones y disponibilidad
	AvailabilityInfo string `json:"availabilityinfo"` // HTML con info de restricciones

	// Campos opcionales
	HiddenFromStudents bool `json:"hiddenfromstudents"` // Si está oculta para estudiantes

	// Módulos de la sección
	Modules []CourseModule `json:"modules" gorm:"-"` // Actividades/recursos
}

// SectionWithChildren representa una sección con su jerarquía de subsecciones
// Utilizado para construir el árbol jerárquico del formato flexsections
type SectionWithChildren struct {
	CourseSection                       // Embebe todos los campos de CourseSection
	Children      []SectionWithChildren `json:"subsections,omitempty"` // Subsecciones anidadas
}

// ============================================================================
// FUNCIÓN PRINCIPAL: GetCourseContent
// Replica fielmente core_course_get_contents de Moodle 4.x
// ============================================================================

// GetCourseContent obtiene el contenido completo del curso (secciones y módulos)
// Compatible 100% con core_course_get_contents de Moodle Web Services
//
// Parámetros:
//   - courseID: ID del curso (mdl_course.id)
//
// Retorna:
//   - Array de CourseSection PLANAS con campo parent (igual que Moodle)
//   - El cliente construirá la jerarquía usando el campo parent
//   - Error si hay problemas de base de datos
func (r *CourseRepository) GetCourseContent(courseID int) ([]CourseSection, error) {
	var sections []CourseSection

	// ================================================================
	// 1. OBTENER TODAS LAS SECCIONES DEL CURSO
	// ================================================================
	err := r.db.Table("mdl_course_sections").
		Select("id, section, name, summary, summaryformat, visible, sequence").
		Where("course = ?", courseID).
		Order("section ASC").
		Scan(&sections).Error

	if err != nil {
		return nil, fmt.Errorf("error obteniendo secciones: %w", err)
	}

	// ================================================================
	// 2. ENRIQUECER CADA SECCIÓN
	// ================================================================
	for i := range sections {
		sectionID := sections[i].ID

		// parent (desde mdl_course_format_options para formato flexsections)
		var parentValue *int
		r.db.Table("mdl_course_format_options").
			Select("CAST(value AS INTEGER)").
			Where("courseid = ? AND sectionid = ? AND name = 'parent'", courseID, sectionID).
			Scan(&parentValue)
		sections[i].Parent = parentValue

		// availability
		var availability string
		r.db.Table("mdl_course_sections").
			Select("availability").
			Where("id = ?", sectionID).
			Scan(&availability)

		if availability != "" && availability != "null" {
			sections[i].AvailabilityInfo = "Restricciones de acceso aplicadas"
		}

		sections[i].UserVisible = sections[i].Visible == 1
		sections[i].HiddenFromStudents = sections[i].Visible == 0

		// módulos
		modules, err := r.getCourseModules(courseID, sectionID)
		if err == nil {
			sections[i].Modules = modules
		}
	}

	// ================================================================
	// RETORNAR SECCIONES PLANAS (como Moodle core_course_get_contents)
	// El frontend construirá la jerarquía usando el campo 'parent'
	// ================================================================
	return sections, nil
}

// ============================================================================
// FUNCIÓN AUXILIAR: getCourseModules
// Obtiene los módulos (actividades/recursos) de una sección específica
// ============================================================================

func (r *CourseRepository) getCourseModules(courseID int, sectionID int) ([]CourseModule, error) {
	var modules []CourseModule

	// ========================================================================
	// PASO 1: Obtener módulos desde mdl_course_modules
	// ========================================================================
	err := r.db.Table("mdl_course_modules AS cm").
		Select(`
			cm.id,
			cm.module,
			cm.instance,
			cm.indent,
			cm.visible,
			cm.completion,
			cm.visibleoncoursepage,
			m.name as modname
		`).
		Joins("JOIN mdl_modules AS m ON m.id = cm.module").
		Where("cm.course = ? AND cm.section = ? AND cm.deletioninprogress = 0", courseID, sectionID).
		Order("cm.id ASC").
		Scan(&modules).Error

	if err != nil {
		return nil, fmt.Errorf("error al obtener módulos: %w", err)
	}

	// ========================================================================
	// PASO 2: Para cada módulo, obtener su NOMBRE REAL desde su tabla
	// ========================================================================
	for j := range modules {
		// Validar que ModName no esté vacío
		if modules[j].ModName == "" {
			continue
		}

		// Determinar la tabla del módulo
		moduleTable := getModuleTableName(modules[j].ModName)
		if moduleTable == "" || moduleTable == "mdl_" {
			continue
		}

		// Obtener el nombre del módulo desde su tabla específica
		var moduleName string
		r.db.Table(moduleTable).
			Select("name").
			Where("id = ?", modules[j].Instance).
			Scan(&moduleName)

		// Asignar nombre
		if moduleName != "" {
			modules[j].Name = moduleName
		} else {
			modules[j].Name = fmt.Sprintf("%s #%d", modules[j].ModName, modules[j].Instance)
		}

		// ====================================================================
		// PASO 3: Generar URL correcta (formato Moodle)
		// ====================================================================
		modules[j].URL = fmt.Sprintf("/mod/%s/view.php?id=%d", modules[j].ModName, modules[j].ID)

		// ====================================================================
		// PASO 4: Calcular UserVisible
		// ====================================================================
		modules[j].UserVisible = modules[j].Visible == 1

		// ====================================================================
		// PASO 5: Campos opcionales (compatibilidad con Moodle)
		// ====================================================================
		modules[j].ModIcon = fmt.Sprintf("/theme/image.php/boost/%s/1/monologo", modules[j].ModName)
		modules[j].ModPlural = getModulePluralName(modules[j].ModName)
	}

	return modules, nil
}

// ============================================================================
// FUNCIONES AUXILIARES
// ============================================================================

// getModuleTableName retorna el nombre de la tabla Moodle para cada tipo de módulo
func getModuleTableName(modname string) string {
	switch modname {
	case "assign":
		return "mdl_assign"
	case "quiz":
		return "mdl_quiz"
	case "resource":
		return "mdl_resource"
	case "forum":
		return "mdl_forum"
	case "page":
		return "mdl_page"
	case "url":
		return "mdl_url"
	case "folder":
		return "mdl_folder"
	case "book":
		return "mdl_book"
	case "label":
		return "mdl_label"
	case "h5pactivity":
		return "mdl_h5pactivity"
	case "choice":
		return "mdl_choice"
	case "data":
		return "mdl_data"
	case "feedback":
		return "mdl_feedback"
	case "glossary":
		return "mdl_glossary"
	case "lesson":
		return "mdl_lesson"
	case "scorm":
		return "mdl_scorm"
	case "survey":
		return "mdl_survey"
	case "wiki":
		return "mdl_wiki"
	case "workshop":
		return "mdl_workshop"
	case "chat":
		return "mdl_chat"
	case "lti":
		return "mdl_lti"
	default:
		if modname != "" {
			return "mdl_" + modname
		}
		return ""
	}
}

// getModulePluralName retorna el nombre plural de un tipo de módulo (para UI)
func getModulePluralName(modname string) string {
	pluralNames := map[string]string{
		"assign":      "Tareas",
		"quiz":        "Cuestionarios",
		"resource":    "Archivos",
		"forum":       "Foros",
		"page":        "Páginas",
		"url":         "URLs",
		"folder":      "Carpetas",
		"book":        "Libros",
		"label":       "Etiquetas",
		"h5pactivity": "Actividades H5P",
		"choice":      "Consultas",
		"data":        "Bases de datos",
		"feedback":    "Retroalimentaciones",
		"glossary":    "Glosarios",
		"lesson":      "Lecciones",
		"scorm":       "Paquetes SCORM",
		"survey":      "Encuestas",
		"wiki":        "Wikis",
		"workshop":    "Talleres",
		"chat":        "Chats",
		"lti":         "Herramientas externas",
	}

	if plural, exists := pluralNames[modname]; exists {
		return plural
	}
	return modname
}
