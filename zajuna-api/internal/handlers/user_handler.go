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

type UserHandler struct {
	service services.UserServiceInterface
}

func NewUserHandler(service services.UserServiceInterface) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsers obtiene la lista de usuarios con filtros y paginación
// @Summary      Listar usuarios
// @Description  Obtiene usuarios con filtros opcionales y paginación
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        firstname  query     string  false  "Filtrar por nombre"
// @Param        lastname   query     string  false  "Filtrar por apellido"
// @Param        username   query     string  false  "Filtrar por username"
// @Param        email      query     string  false  "Filtrar por email"
// @Param        page       query     int     false  "Número de página"  default(1)
// @Param        limit      query     int     false  "Elementos por página"  default(15)
// @Success      200        {object}  response.PaginatedResponse
// @Failure      400        {object}  response.ErrorResponse
// @Failure      500        {object}  response.ErrorResponse
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	// 1. Parsear y validar request
	var req request.GetUsersRequest

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

	// 3. Convertir a filtros para el servicio
	filters := req.ToFilterMap()

	// 4. Llamar al servicio
	users, total, err := h.service.GetUsers(filters, req.Page, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener usuarios",
			err.Error(),
		))
		return
	}

	// 5. Convertir modelos a DTOs
	usersResponse := mapper.UsersToResponse(users)

	// 6. Crear respuesta paginada
	paginatedResponse := response.NewPaginatedResponse(
		usersResponse,
		req.Page,
		req.Limit,
		total,
	)

	// 7. Responder
	c.JSON(http.StatusOK, paginatedResponse)
}

// UpdateUsers actualiza múltiples usuarios
// @Summary      Actualizar usuarios
// @Description  Actualiza la información de uno o más usuarios
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body request.UpdateUsersRequest true "Lista de usuarios a actualizar"
// @Success      200 {object} response.UpdateUserResponse
// @Failure      400 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /users/update [put]
func (h *UserHandler) UpdateUsers(c *gin.Context) {
	// 1. Parsear y validar request
	var req request.UpdateUsersRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Convertir DTOs a modelos
	usersToUpdate := make([]models.User, len(req.Users))
	for i, userReq := range req.Users {
		usersToUpdate[i] = models.User{
			ID:        userReq.ID,
			FirstName: userReq.FirstName,
			LastName:  userReq.LastName,
			Email:     userReq.Email,
			City:      userReq.City,
			Country:   userReq.Country,
			Lang:      userReq.Lang,
			Timezone:  userReq.Timezone,
			Phone1:    userReq.Phone1,
		}
	}

	// 3. Llamar al servicio
	updated, err := h.service.UpdateUsers(usersToUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"UPDATE_FAILED",
			"Error al actualizar usuarios",
			err.Error(),
		))
		return
	}

	// 4. Responder
	c.JSON(http.StatusOK, response.UpdateUserResponse{
		Message:  "Usuarios actualizados correctamente",
		Updated:  updated,
		Warnings: []string{},
	})
}

// DeleteUsers suspende múltiples usuarios (soft delete)
// @Summary      Eliminar usuarios
// @Description  Suspende usuarios (soft delete) por sus IDs
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body request.DeleteUsersRequest true "IDs de usuarios a eliminar"
// @Success      200 {object} response.DeleteUserResponse
// @Failure      400 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /users [delete]
func (h *UserHandler) DeleteUsers(c *gin.Context) {
	// 1. Parsear request
	var req request.DeleteUsersRequest

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
	if err := h.service.DeleteUsers(req.UserIDs); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"DELETE_FAILED",
			"Error al suspender usuarios",
			err.Error(),
		))
		return
	}

	// 4. Responder
	c.JSON(http.StatusOK, response.DeleteUserResponse{
		Message: "Usuarios suspendidos correctamente",
		Deleted: len(req.UserIDs),
		Errors:  []string{},
	})
}

// GetEnrolledUsers obtiene usuarios matriculados en un curso
// @Summary      Obtener usuarios matriculados en un curso
// @Description  Obtiene la lista de usuarios matriculados en un curso con opciones de filtrado
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        courseid        path      int     true   "ID del curso"
// @Param        withcapability  query     string  false  "Filtrar por capacidad"
// @Param        groupid         query     int     false  "Filtrar por grupo"
// @Param        onlyactive      query     int     false  "Solo usuarios activos (1=sí, 0=no)"
// @Param        onlysuspended   query     int     false  "Solo usuarios suspendidos (1=sí, 0=no)"
// @Param        userfields      query     string  false  "Campos de usuario a retornar"
// @Param        limitfrom       query     int     false  "Offset de paginación"
// @Param        limitnumber     query     int     false  "Límite de resultados"
// @Param        sortby          query     string  false  "Campo de ordenamiento (id, firstname, lastname, siteorder)"
// @Param        sortdirection   query     string  false  "Dirección de ordenamiento (ASC, DESC)"
// @Success      200             {object}  response.EnrolledUsersListResponse
// @Failure      400             {object}  response.ErrorResponse
// @Failure      500             {object}  response.ErrorResponse
// @Router       /courses/{courseid}/users [get]
func (h *UserHandler) GetEnrolledUsers(c *gin.Context) {
	// 1. Parsear y validar courseID del path
	var uriReq request.GetEnrolledUsersRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_COURSE_ID",
			"ID de curso inválido",
			err.Error(),
		))
		return
	}

	// 2. Parsear opciones de query string
	var options request.EnrolledUsersOptions
	if err := c.ShouldBindQuery(&options); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_OPTIONS",
			"Opciones de consulta inválidas",
			err.Error(),
		))
		return
	}

	// 3. Establecer valores por defecto
	options.SetDefaults()

	// 4. Validar opciones
	if err := options.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"VALIDATION_ERROR",
			err.Error(),
			nil,
		))
		return
	}

	// 5. Convertir opciones a map para el service
	optionsMap := map[string]interface{}{
		"sortby":        options.SortBy,
		"sortdirection": options.SortDirection,
		"limitnumber":   options.LimitNumber,
	}

	if options.WithCapability != "" {
		optionsMap["withcapability"] = options.WithCapability
	}
	if options.GroupID > 0 {
		optionsMap["groupid"] = options.GroupID
	}
	if options.OnlyActive > 0 {
		optionsMap["onlyactive"] = options.OnlyActive
	}
	if options.OnlySuspended > 0 {
		optionsMap["onlysuspended"] = options.OnlySuspended
	}
	if options.LimitFrom > 0 {
		optionsMap["limitfrom"] = options.LimitFrom
	}

	// 6. Llamar al servicio
	users, total, err := h.service.GetEnrolledUsers(uriReq.CourseID, optionsMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener usuarios matriculados",
			err.Error(),
		))
		return
	}

	// 7. Responder
	c.JSON(http.StatusOK, response.EnrolledUsersListResponse{
		Users: users,
		Total: total,
	})
}
