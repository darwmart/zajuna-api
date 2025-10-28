package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/mapper"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service *services.CourseService
}

func NewCourseHandler(service *services.CourseService) *CourseHandler {
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
// @Router       /courses [get]
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
// @Param        id   path      int  true  "ID del curso"
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
	details, err := h.service.GetCourseDetails(req.ID)
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

	// 3. Llamar al servicio
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
