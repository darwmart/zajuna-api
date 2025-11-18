package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/mapper"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

// CategoryHandler maneja las solicitudes relacionadas con categorías
type CategoryHandler struct {
	service services.CategoryServiceInterface
}

// NewCategoryHandler constructor para inyectar el servicio
func NewCategoryHandler(service services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetCategories obtiene la lista de categorías
// @Summary      Listar categorías
// @Description  Obtiene todas las categorías de cursos disponibles
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.CategoryListResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	// 1. Llamar al servicio
	categories, err := h.service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener las categorías",
			err.Error(),
		))
		return
	}

	// 2. Convertir modelos a DTOs
	categoriesResponse := mapper.CategoriesToResponse(categories)

	// 3. Crear respuesta
	listResponse := response.CategoryListResponse{
		Categories: categoriesResponse,
	}

	// 4. Responder
	c.JSON(http.StatusOK, listResponse)
}

// MoveCategory mueve una categoría antes de otra categoría especificada
// @Summary      Mover categoría
// @Description  Mueve una categoría antes de otra categoría especificada (beforeid=0 mueve al final)
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        request body request.MoveCategoryRequest true "Datos de la categoría a mover"
// @Success      200  {object}  response.MoveCategoryResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /categories/move [post]
// MoveCategory mueve una categoría dentro del árbol jerárquico
// Soporta:
//   - Reordenamiento entre categorías hermanas (cambio de sortorder)
//   - Cambio de padre (cambio de parentid)
func (h *CategoryHandler) MoveCategory(c *gin.Context) {
	// 1. Parsear y validar request
	var req request.MoveCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Validación adicional
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"VALIDATION_ERROR",
			err.Error(),
			nil,
		))
		return
	}

	// 3. Llamar al servicio
	if err := h.service.MoveCategory(req.ID, req.BeforeID, req.ParentID); err != nil {
		// Verificar si es error de "no encontrado"
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.NewErrorResponse(
				"CATEGORY_NOT_FOUND",
				"Categoría o padre no encontrado",
				err.Error(),
			))
			return
		}

		// Verificar si es error de validación (categorías con diferente parent)
		if err.Error() == "invalid value" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				"INVALID_OPERATION",
				"La categoría beforeid debe tener el mismo padre que el nuevo padre especificado",
				err.Error(),
			))
			return
		}

		// Otros errores de base de datos
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"MOVE_FAILED",
			"Error al mover la categoría",
			err.Error(),
		))
		return
	}

	// 4. Construir respuesta
	message := "Categoría movida correctamente"
	if req.BeforeID == 0 {
		message = "Categoría movida al final correctamente"
	}
	if req.ParentID != nil {
		message = "Categoría movida y reasignada correctamente"
	}

	// Determinar el nuevo padre para la respuesta
	var newParent uint
	if req.ParentID != nil {
		newParent = *req.ParentID
	}

	c.JSON(http.StatusOK, response.MoveCategoryResponse{
		Message:    message,
		CategoryID: req.ID,
		NewParent:  newParent,
	})
}
