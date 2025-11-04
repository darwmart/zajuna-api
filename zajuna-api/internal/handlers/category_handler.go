package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/mapper"
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
