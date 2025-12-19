package repository

import (
	"time"
	"zajunaApi/internal/models"

	"golang.org/x/crypto/bcrypt"
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

	// Búsqueda global (q) - Busca en todos los campos como Moodle
	if q := filters["q"]; q != "" {
		searchPattern := "%" + q + "%"
		query = query.Where(`
			firstname ILIKE ?
			OR lastname ILIKE ?
			OR username ILIKE ?
			OR email ILIKE ?
		`,
			searchPattern,
			searchPattern,
			searchPattern,
			searchPattern,
		)
	} else {
		// Aplicar filtros individuales solo si NO hay búsqueda global
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
	}

	// Contar total antes de paginar
	query.Count(&total)

	// Ordenar alfabéticamente por firstname, luego por lastname
	query = query.Order("firstname ASC, lastname ASC")

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
		// Construir mapa dinámico solo con los campos que fueron enviados
		updates := make(map[string]interface{})

		// El handler solo asigna campos desde punteros no-nil,
		// por lo que solo los campos enviados en el request tendrán valores
		if u.FirstName != "" {
			updates["firstname"] = u.FirstName
		}
		if u.LastName != "" {
			updates["lastname"] = u.LastName
		}
		if u.Email != "" {
			updates["email"] = u.Email
		}
		if u.City != "" {
			updates["city"] = u.City
		}
		if u.Country != "" {
			updates["country"] = u.Country
		}
		if u.Lang != "" {
			updates["lang"] = u.Lang
		}
		if u.Timezone != "" {
			updates["timezone"] = u.Timezone
		}
		if u.Phone1 != "" {
			updates["phone1"] = u.Phone1
		}

		// Para suspended: siempre incluir si fue enviado (handler lo copia solo si no es nil)
		// Como models.User tiene Suspended como int (no puntero), necesitamos
		// verificar indirectamente. El handler marca suspended=-1 cuando NO se envió.
		if u.Suspended >= 0 && u.Suspended <= 1 {
			updates["suspended"] = u.Suspended
		}

		if u.Deleted >= 0 && u.Deleted <= 1 {
			updates["deleted"] = u.Deleted
		}

		// Si no hay campos para actualizar, saltar
		if len(updates) == 0 {
			continue
		}

		if err := r.DB.Model(&models.User{}).
			Where("id = ?", u.ID).
			Updates(updates).Error; err != nil {
			return total, err
		}
		total++
	}
	return total, nil
}

// ToggleUserStatus cambia el estado suspended de un usuario (0 <-> 1)
func (r *UserRepository) ToggleUserStatus(userID uint) (int, error) {
	var user models.User

	// Buscar el usuario
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return -1, err
	}

	// Cambiar el estado al opuesto
	newStatus := 0
	if user.Suspended == 0 {
		newStatus = 1
	}

	// Actualizar en la base de datos
	if err := r.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("suspended", newStatus).Error; err != nil {
		return -1, err
	}

	return newStatus, nil
}

// EnrolledUserDetail representa los detalles completos de un usuario matriculado
type EnrolledUserDetail struct {
	models.User
	LastCourseAccess int64 `gorm:"column:lastcourseaccess"`
}

// GetEnrolledUsers obtiene usuarios matriculados en un curso con opciones de filtrado
func (r *UserRepository) GetEnrolledUsers(courseID int, options map[string]interface{}) ([]EnrolledUserDetail, int, error) {
	var users []EnrolledUserDetail

	// Query base: usuarios matriculados en el curso
	query := r.DB.Table("mdl_user u").
		Select("u.*, COALESCE(ula.timeaccess, 0) as lastcourseaccess").
		Joins("JOIN mdl_user_enrolments ue ON ue.userid = u.id").
		Joins("JOIN mdl_enrol e ON e.id = ue.enrolid").
		Joins("LEFT JOIN mdl_user_lastaccess ula ON ula.userid = u.id AND ula.courseid = ?", courseID).
		Where("e.courseid = ? AND u.deleted = 0", courseID)

	// Filtro onlyactive: solo usuarios con matriculaciones activas
	if onlyActive, ok := options["onlyactive"].(int); ok && onlyActive == 1 {
		query = query.Where("ue.status = 0") // 0 = active
		// Filtrar por restricciones de tiempo si aplica
		query = query.Where("(ue.timestart = 0 OR ue.timestart <= EXTRACT(EPOCH FROM NOW()))")
		query = query.Where("(ue.timeend = 0 OR ue.timeend >= EXTRACT(EPOCH FROM NOW()))")
	}

	// Filtro onlysuspended: solo usuarios suspendidos
	if onlySuspended, ok := options["onlysuspended"].(int); ok && onlySuspended == 1 {
		query = query.Where("ue.status = 1") // 1 = suspended
	}

	// Filtro por grupo
	if groupID, ok := options["groupid"].(int); ok && groupID > 0 {
		query = query.Joins("JOIN mdl_groups_members gm ON gm.userid = u.id").
			Where("gm.groupid = ?", groupID)
	}

	// Filtro por capacidad (requiere verificar role capabilities)
	if withCapability, ok := options["withcapability"].(string); ok && withCapability != "" {
		// Obtener el contexto del curso
		query = query.Joins(`
			JOIN mdl_role_assignments ra ON ra.userid = u.id
			JOIN mdl_context ctx ON ctx.id = ra.contextid
			JOIN mdl_role_capabilities rc ON rc.roleid = ra.roleid
		`).Where("ctx.contextlevel = 50 AND ctx.instanceid = ?", courseID).
			Where("rc.capability = ? AND rc.permission = 1", withCapability)
	}

	// Contar total antes de aplicar límites
	var total int64
	query.Count(&total)

	// Ordenamiento
	sortBy := "u.id"
	sortDirection := "ASC"
	if sb, ok := options["sortby"].(string); ok && sb != "" {
		switch sb {
		case "firstname":
			sortBy = "u.firstname"
		case "lastname":
			sortBy = "u.lastname"
		case "siteorder":
			sortBy = "u.id" // Moodle usa ID para siteorder por defecto
		default:
			sortBy = "u.id"
		}
	}
	if sd, ok := options["sortdirection"].(string); ok && (sd == "DESC" || sd == "ASC") {
		sortDirection = sd
	}
	query = query.Order(sortBy + " " + sortDirection)

	// Paginación
	if limitFrom, ok := options["limitfrom"].(int); ok && limitFrom > 0 {
		query = query.Offset(limitFrom)
	}
	if limitNumber, ok := options["limitnumber"].(int); ok && limitNumber > 0 {
		query = query.Limit(limitNumber)
	}

	// Ejecutar query
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

// GetUserGroupsInCourse obtiene los grupos de un usuario en un curso específico
func (r *UserRepository) GetUserGroupsInCourse(userID, courseID int) ([]map[string]interface{}, error) {
	var groups []map[string]interface{}

	err := r.DB.Table("mdl_groups g").
		Select("g.id, g.name, g.description, g.descriptionformat").
		Joins("JOIN mdl_groups_members gm ON gm.groupid = g.id").
		Where("gm.userid = ? AND g.courseid = ?", userID, courseID).
		Scan(&groups).Error

	return groups, err
}

// GetUserRolesInCourse obtiene los roles de un usuario en un curso específico
func (r *UserRepository) GetUserRolesInCourse(userID, courseID int) ([]map[string]interface{}, error) {
	var roles []map[string]interface{}

	err := r.DB.Table("mdl_role r").
		Select("r.id as roleid, r.name, r.shortname, r.sortorder").
		Joins("JOIN mdl_role_assignments ra ON ra.roleid = r.id").
		Joins("JOIN mdl_context ctx ON ctx.id = ra.contextid").
		Where("ra.userid = ? AND ctx.contextlevel = 50 AND ctx.instanceid = ?", userID, courseID).
		Scan(&roles).Error

	return roles, err
}

// GetUserCustomFields obtiene los campos personalizados de un usuario
func (r *UserRepository) GetUserCustomFields(userID int) ([]map[string]interface{}, error) {
	var customFields []map[string]interface{}

	err := r.DB.Table("mdl_user_info_data uid").
		Select("uif.datatype as type, uid.data as value, uif.name, uif.shortname").
		Joins("JOIN mdl_user_info_field uif ON uif.id = uid.fieldid").
		Where("uid.userid = ?", userID).
		Scan(&customFields).Error

	return customFields, err
}

// GetUserPreferences obtiene las preferencias de un usuario
func (r *UserRepository) GetUserPreferences(userID int) ([]map[string]interface{}, error) {
	var preferences []map[string]interface{}

	err := r.DB.Table("mdl_user_preferences").
		Select("name, value").
		Where("userid = ?", userID).
		Scan(&preferences).Error

	return preferences, err
}

// GetUserEnrolledCourses obtiene los cursos en los que está matriculado un usuario
func (r *UserRepository) GetUserEnrolledCourses(userID int) ([]map[string]interface{}, error) {
	var courses []map[string]interface{}

	err := r.DB.Table("mdl_course c").
		Select("c.id, c.fullname, c.shortname").
		Joins("JOIN mdl_enrol e ON e.courseid = c.id").
		Joins("JOIN mdl_user_enrolments ue ON ue.enrolid = e.id").
		Where("ue.userid = ? AND ue.status = 0", userID).
		Scan(&courses).Error

	return courses, err
}

// FindByID obtiene un usuario por su ID (solo activos y no eliminados, igual que Moodle)
func (r *UserRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User

	// Buscar usuario con las mismas condiciones que Moodle:
	// - deleted = 0 (no eliminado)
	// - suspended = 0 (no suspendido)
	// - confirmed = 1 (email confirmado)
	err := r.DB.Where("id = ? AND deleted = 0 AND suspended = 0 AND confirmed = 1", userID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser crea un nuevo usuario en la base de datos
func (r *UserRepository) CreateUser(user *models.User, password string) error {
	// Hash del password usando bcrypt (como Moodle)
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	// Establecer valores por defecto compatibles con Moodle
	now := timeNow().Unix()
	user.Auth = "manual"
	user.Confirmed = 1
	user.MNetHostID = 1
	user.Deleted = 0
	user.Suspended = 0
	user.MailFormat = 1
	user.MailDigest = 0
	user.MailDisplay = 2
	user.AutoSubscribe = 1
	user.TrackForums = 0
	user.TimeCreated = now
	user.TimeModified = now
	user.FirstAccess = 0
	user.LastAccess = 0
	user.LastLogin = 0
	user.CurrentLogin = 0
	user.CalendarType = "gregorian"

	// Establecer valores por defecto si están vacíos
	if user.Lang == "" {
		user.Lang = "es"
	}
	if user.Timezone == "" {
		user.Timezone = "America/Bogota"
	}
	if user.Country == "" {
		user.Country = "CO"
	}

	// Crear un mapa con todos los campos del usuario + password
	userData := map[string]interface{}{
		"username":       user.Username,
		"password":       hashedPassword,
		"firstname":      user.FirstName,
		"lastname":       user.LastName,
		"email":          user.Email,
		"auth":           user.Auth,
		"confirmed":      user.Confirmed,
		"mnethostid":     user.MNetHostID,
		"deleted":        user.Deleted,
		"suspended":      user.Suspended,
		"mailformat":     user.MailFormat,
		"maildigest":     user.MailDigest,
		"maildisplay":    user.MailDisplay,
		"autosubscribe":  user.AutoSubscribe,
		"trackforums":    user.TrackForums,
		"timecreated":    user.TimeCreated,
		"timemodified":   user.TimeModified,
		"firstaccess":    user.FirstAccess,
		"lastaccess":     user.LastAccess,
		"lastlogin":      user.LastLogin,
		"currentlogin":   user.CurrentLogin,
		"calendartype":   user.CalendarType,
		"lang":           user.Lang,
		"timezone":       user.Timezone,
		"city":           user.City,
		"country":        user.Country,
	}

	// Agregar campos opcionales solo si no están vacíos
	if user.Phone1 != "" {
		userData["phone1"] = user.Phone1
	}
	if user.Phone2 != "" {
		userData["phone2"] = user.Phone2
	}
	if user.Institution != "" {
		userData["institution"] = user.Institution
	}
	if user.Department != "" {
		userData["department"] = user.Department
	}
	if user.Address != "" {
		userData["address"] = user.Address
	}
	if user.IDNumber != "" {
		userData["idnumber"] = user.IDNumber
	}
	if user.Description != "" {
		userData["description"] = user.Description
	}
	if user.Theme != "" {
		userData["theme"] = user.Theme
	}
	if user.Interests != "" {
		userData["interests"] = user.Interests
	}

	// Insertar el usuario
	result := r.DB.Table("mdl_user").Create(userData)
	if result.Error != nil {
		return result.Error
	}

	// Obtener el ID del usuario recién creado
	var createdUser models.User
	if err := r.DB.Table("mdl_user").
		Where("username = ?", user.Username).
		First(&createdUser).Error; err != nil {
		return err
	}

	// Actualizar el user original con el ID generado
	user.ID = createdUser.ID

	return nil
}

// hashPassword hashea una contraseña usando bcrypt (compatible con Moodle)
func hashPassword(password string) (string, error) {
	// Moodle usa bcrypt con cost 10
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// timeNow retorna el tiempo actual (puede ser mockeado en tests)
func timeNow() time.Time {
	return time.Now()
}

