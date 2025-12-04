package handlers

import (
	"net/http"
	"strconv"
	"zajunaApi/internal/dto/mapper"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service services.CourseServiceInterface
}

func NewCourseHandler(service services.CourseServiceInterface) *CourseHandler {
	return &CourseHandler{service: service}
}

// GetCourses obtiene la lista de cursos con filtros opcionales
// @Summary      Listar cursos
// @Description  Obtiene cursos con filtros opcionales por categoría
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        categoryid  query     int  false  "Filtrar por ID de categoría"
// @Success      200         {object}  response.CourseListResponse
// @Failure      400         {object}  response.ErrorResponse
// @Failure      500         {object}  response.ErrorResponse
// @Router       /courses get
func (h *CourseHandler) GetCourses(c *gin.Context) {
	// 1. Parsear y validar request
	var req request.GetCoursesRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_PARAMS",
			"Parámetros de consulta inválidos",
			err.Error(),
		))
		return
	}

	// 2. Establecer valores por defecto
	req.SetDefaults()

	var courses []models.Course
	var err error

	// 3. Obtener cursos según filtros
	if req.HasCategoryFilter() {
		courses, err = h.service.GetCoursesByCategory(uint(req.CategoryID))
	} else {
		courses, err = h.service.GetAllCourses()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener los cursos",
			err.Error(),
		))
		return
	}

	// 4. Convertir modelos a DTOs
	coursesResponse := mapper.CoursesToResponse(courses)

	// 5. Crear respuesta
	listResponse := response.CourseListResponse{
		Courses: coursesResponse,
	}

	// 6. Responder
	c.JSON(http.StatusOK, listResponse)
}

// GetCourseDetails obtiene los detalles completos de un curso
// @Summary      Obtener detalles de curso
// @Description  Obtiene información detallada de un curso incluyendo roles, grupos y secciones
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  ID del curso
// @Success      200  {object}  response.CourseDetailResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /courses/{id}/details [get]
func (h *CourseHandler) GetCourseDetails(c *gin.Context) {
	// 1. Parsear y validar ID del parámetro URI
	var req request.GetCourseDetailsRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_ID",
			"ID de curso inválido",
			err.Error(),
		))
		return
	}

	// 2. Llamar al servicio
	details, err := h.service.GetCourseDetails(req.IDNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewErrorResponse(
			"NOT_FOUND",
			"Curso no encontrado",
			err.Error(),
		))
		return
	}

	// 3. Convertir a DTO
	detailsResponse := mapper.CourseDetailsToResponse(details)

	// 4. Responder
	c.JSON(http.StatusOK, detailsResponse)
}

// DeleteCourses elimina múltiples cursos
// @Summary      Eliminar cursos
// @Description  Elimina uno o más cursos por sus IDs
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        request body request.DeleteCoursesRequest true "IDs de cursos a eliminar"
// @Success      200 {object} response.DeleteCoursesResponse
// @Failure      400 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /courses [delete]
func (h *CourseHandler) DeleteCourses(c *gin.Context) {
	// 1. Parsear request
	var req request.DeleteCoursesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Validación adicional personalizada
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"VALIDATION_ERROR",
			err.Error(),
			nil,
		))
		return
	}

	// 3. Llamar al servicio.
	serviceResponse, err := h.service.DeleteCourses(req.CourseIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"DELETE_FAILED",
			"Error al eliminar cursos",
			err.Error(),
		))
		return
	}

	// 4. Convertir warnings a DTO
	warningsResponse := mapper.DeleteCoursesWarningsToResponse(serviceResponse.Warnings)

	// 5. Calcular número de cursos eliminados exitosamente
	deleted := len(req.CourseIDs) - len(serviceResponse.Warnings)

	// 6. Responder
	c.JSON(http.StatusOK, response.DeleteCoursesResponse{
		Message:  "Operación completada",
		Deleted:  deleted,
		Warnings: warningsResponse,
	})
}

// UpdateCourses actualiza múltiples cursos
// @Summary      Actualizar cursos
// @Description  Actualiza uno o más cursos con los datos proporcionados
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        request body request.UpdateCoursesRequest true "Cursos a actualizar"
// @Success      200 {object} response.UpdateCoursesResponse
// @Failure      400 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /courses [put]
func (h *CourseHandler) UpdateCourses(c *gin.Context) {
	// 1. Parsear request
	var req request.UpdateCoursesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Validación adicional personalizada
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"VALIDATION_ERROR",
			err.Error(),
			nil,
		))
		return
	}

	// 3. Llamar al servicio
	serviceResponse, err := h.service.UpdateCourses(req.Courses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"UPDATE_FAILED",
			"Error al actualizar cursos",
			err.Error(),
		))
		return
	}

	// 4. Convertir warnings a DTO
	warningsResponse := mapper.UpdateCoursesWarningsToResponse(serviceResponse.Warnings)

	// 5. Responder
	c.JSON(http.StatusOK, response.UpdateCoursesResponse{
		Warnings: warningsResponse,
	})
}

// MoveCourses mueve múltiples cursos a nuevas categorías
// @Summary      Mover cursos
// @Description  Mueve uno o más cursos a nuevas categorías (compatible con core_course_move_courses de Moodle)
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        request body request.MoveCoursesRequest true "Cursos a mover"
// @Success      200 {object} response.MoveCoursesResponse
// @Failure      400 {object} response.ErrorResponse
// @Failure      404 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /courses/move [post]
func (h *CourseHandler) MoveCourses(c *gin.Context) {

	// 1. Parsear request
	var req request.MoveCoursesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Validar cada curso individualmente
	for i, course := range req.Courses {
		if err := course.Validate(); err != nil {
			valErr := err.(*request.ValidationError)
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				"VALIDATION_ERROR",
				valErr.Message,
				map[string]interface{}{
					"field": "courses[" + string(rune(i)) + "]." + valErr.Field,
				},
			))
			return
		}
	}

	// 3. Llamar al servicio
	if err := h.service.MoveCourses(req.Courses); err != nil {
		// Verificar si es error de "no encontrado"
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.NewErrorResponse(
				"COURSE_NOT_FOUND",
				"Curso o categoría no encontrada",
				err.Error(),
			))
			return
		}

		// Verificar si es error de validación
		if err.Error() == "invalid value" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				"INVALID_OPERATION",
				"Operación inválida: verifica que el beforeid exista en la categoría destino",
				err.Error(),
			))
			return
		}

		// Otros errores de base de datos
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"MOVE_FAILED",
			"Error al mover cursos",
			err.Error(),
		))
		return
	}

	// 4. Responder
	c.JSON(http.StatusOK, response.MoveCoursesResponse{
		Message: "Cursos movidos correctamente",
		Moved:   len(req.Courses),
	})
}

// GetMyCourses obtiene los cursos donde el usuario autenticado es instructor
// @Summary      Listar mis cursos
// @Description  Obtiene todos los cursos donde el usuario autenticado tiene rol de instructor (editingteacher o teacher)
// @Tags         courses
// @Accept       json
// @Produce      json
// @Success      200 {object} response.CourseListResponse
// @Failure      401 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /courses/my-courses [get]
func (h *CourseHandler) GetMyCourses(c *gin.Context) {
	// 1. Obtener el userID del contexto (establecido por AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(
			"UNAUTHORIZED",
			"Usuario no autenticado",
			nil,
		))
		return
	}

	// 2. Convertir userID a int
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"INTERNAL_ERROR",
			"Error al procesar ID de usuario",
			nil,
		))
		return
	}

	// 3. Obtener cursos del servicio
	courses, err := h.service.GetCoursesWhereUserIsTeacher(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener los cursos",
			err.Error(),
		))
		return
	}

	// 4. Convertir modelos a DTOs
	coursesResponse := mapper.CoursesToResponse(courses)

	// 5. Crear respuesta
	listResponse := response.CourseListResponse{
		Courses: coursesResponse,
	}

	// 6. Responder
	c.JSON(http.StatusOK, listResponse)
}

// GetCourseContent godoc
// @Summary      Obtener contenido del curso (secciones y módulos)
// @Description  Obtiene todas las secciones y actividades de un curso. Compatible con core_course_get_contents de Moodle
// @Tags         courses
// @Security     BearerAuth
// @Produce      json
// @Param        id path int true "ID del curso"
// @Success      200 {array} repository.CourseSection
// @Failure      400 {object} response.ErrorResponse
// @Failure      401 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /courses/{id}/content [get]
func (h *CourseHandler) GetCourseContent(c *gin.Context) {
	// 1. Obtener userID del contexto (establecido por AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse(
			"UNAUTHORIZED",
			"Usuario no autenticado",
			"",
		))
		return
	}

	// Convertir a int
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"INTERNAL_ERROR",
			"Error al obtener ID de usuario",
			"",
		))
		return
	}

	// 2. Obtener courseID del path parameter
	courseIDStr := c.Param("id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_COURSE_ID",
			"ID de curso inválido",
			err.Error(),
		))
		return
	}

	// 3. Obtener contenido del curso desde el servicio (filtrado por permisos)
	sections, err := h.service.GetCourseContent(courseID, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener el contenido del curso",
			err.Error(),
		))
		return
	}

	// 4. Responder con las secciones
	c.JSON(http.StatusOK, sections)
}
