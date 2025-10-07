package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zajunaApi/internal/services"
)

type CourseHandler struct {
	service *services.CourseService
}

func NewCourseHandler(service *services.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

// GET /api/courses?categoryId=3
func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	// Obtener categoryId desde los parámetros de la URL
	categoryIDStr := r.URL.Query().Get("categoryid")
	if categoryIDStr == "" {
		http.Error(w, "Falta el parámetro categoryid", http.StatusBadRequest)
		return
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		http.Error(w, "categoryId debe ser un número válido", http.StatusBadRequest)
		return
	}

	// Llamar al servicio
	courses, err := h.service.GetCoursesByCategory(categoryID)
	if err != nil {
		http.Error(w, "Error al obtener los cursos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"courses": courses,
	})
}
