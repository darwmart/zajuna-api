package handlers

import (
	"encoding/json"
	"net/http"
	"zajunaApi/internal/services"
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
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
