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
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
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

	// 2. Validar que haya al menos un usuario
	if len(req.Users) == 0 {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"EMPTY_REQUEST",
			"Debe proporcionar al menos un usuario para actualizar",
			nil,
		))
		return
	}

	// 3. Convertir DTOs a modelos
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

	// 4. Llamar al servicio
	updated, err := h.service.UpdateUsers(usersToUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"UPDATE_FAILED",
			"Error al actualizar usuarios",
			err.Error(),
		))
		return
	}

	// 5. Responder
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
