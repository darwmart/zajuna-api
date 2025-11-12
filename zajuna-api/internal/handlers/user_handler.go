package handlers

import (
	"net/http"
	"time"
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
		user := models.User{
			ID:        userReq.ID,
			Suspended: -1, // Marca para indicar "no enviado"
			Deleted:   -1, // Marca para indicar "no enviado"
		}

		// Solo copiar campos que no son nil (fueron enviados en el request)
		if userReq.FirstName != nil {
			user.FirstName = *userReq.FirstName
		}
		if userReq.LastName != nil {
			user.LastName = *userReq.LastName
		}
		if userReq.Email != nil {
			user.Email = *userReq.Email
		}
		if userReq.City != nil {
			user.City = *userReq.City
		}
		if userReq.Country != nil {
			user.Country = *userReq.Country
		}
		if userReq.Lang != nil {
			user.Lang = *userReq.Lang
		}
		if userReq.Timezone != nil {
			user.Timezone = *userReq.Timezone
		}
		if userReq.Phone1 != nil {
			user.Phone1 = *userReq.Phone1
		}
		if userReq.Suspended != nil {
			user.Suspended = *userReq.Suspended
		}
		if userReq.Deleted != nil {
			user.Deleted = *userReq.Deleted
		}

		usersToUpdate[i] = user
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

// ToggleUserStatus cambia el estado suspended de un usuario (activo <-> suspendido)
// @Summary      Cambiar estado de usuario
// @Description  Alterna el estado suspended de un usuario entre 0 (activo) y 1 (suspendido)
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID del usuario"
// @Success      200  {object}  response.ToggleUserStatusResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /users/{id}/toggle-status [put]
func (h *UserHandler) ToggleUserStatus(c *gin.Context) {
	// 1. Obtener el ID del usuario desde la URL
	var uriParam struct {
		ID uint `uri:"id" binding:"required,min=1"`
	}

	if err := c.ShouldBindUri(&uriParam); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_USER_ID",
			"ID de usuario inválido",
			err.Error(),
		))
		return
	}

	// 2. Llamar al servicio para cambiar el estado
	newStatus, err := h.service.ToggleUserStatus(uriParam.ID)
	if err != nil {
		// Verificar si es error de "no encontrado"
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.NewErrorResponse(
				"USER_NOT_FOUND",
				"Usuario no encontrado",
				err.Error(),
			))
			return
		}

		// Otros errores de base de datos
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"TOGGLE_STATUS_FAILED",
			"Error al cambiar el estado del usuario",
			err.Error(),
		))
		return
	}

	// 3. Construir respuesta
	statusText := "activo"
	message := "Usuario activado correctamente"
	if newStatus == 1 {
		statusText = "suspendido"
		message = "Usuario suspendido correctamente"
	}

	c.JSON(http.StatusOK, response.ToggleUserStatusResponse{
		Message:    message,
		UserID:     uriParam.ID,
		NewStatus:  newStatus,
		StatusText: statusText,
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

// Login autentica a un usuario y genera una cookie de sesión.
//
// @Summary      Iniciar sesión
// @Description  Valida credenciales del usuario, genera un token y lo almacena como cookie HttpOnly.
// @Tags         auth
// @Accept       json
// @Produce      json
//
// @Param        body   body      object  true   "Credenciales del usuario"
// @Success      200    {object}  map[string]interface{}   "Inicio de sesión exitoso"
// @Failure      400    {object}  map[string]string         "Error en credenciales o body inválido"
// @Failure      401    {object}  map[string]string         "No autorizado"
// @Router       /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	token, err := h.service.Login(c.Request, body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Authorization", token, 3600*3, "", "", false, true)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  time.Now().Add(3 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		Domain:   "",
		Secure:   true,
	})

	c.JSON(http.StatusOK, gin.H{})
}

// Logout elimina la sesión del usuario y limpia la cookie Authorization.
//
// @Summary      Cerrar sesión
// @Description  Borra la cookie de autenticación y elimina la sesión en base de datos.
// @Tags         auth
// @Produce      json
// @Security     CookieAuth
//
// @Success      200    {object}  map[string]string    "Sesión eliminada correctamente"
// @Failure      400    {object}  map[string]string    "Error eliminando la sesión"
// @Failure      401    {object}  map[string]string    "No autorizado"
// @Router       /logout [post]
func (h *UserHandler) Logout(c *gin.Context) {

	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := h.service.Logout(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//c.SetCookie("Authorization", "", -1, "", "", false, true)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/", // o el path con el que se creó
		Domain:   "",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true, // debe coincidir con el original
	})

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
