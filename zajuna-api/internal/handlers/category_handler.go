package handlers

import (
	"net/http"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

// CategoryHandler maneja las solicitudes relacionadas con categorías
type CategoryHandler struct {
	service *services.CategoryService
}

// NewCategoryHandler constructor para inyectar el servicio
func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetCategories maneja GET /api/categories
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
