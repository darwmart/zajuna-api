package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize := 30

	// armar filtros
	filters := map[string]string{
		"id":        q.Get("id"),
		"username":  q.Get("username"),
		"firstname": q.Get("firstname"),
		"lastname":  q.Get("lastname"),
		"email":     q.Get("email"),
		"idnumber":  q.Get("idnumber"),
		"auth":      q.Get("auth"),
	}

	users, total, err := h.service.GetUsers(filters, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var totalPages int
	var hasNextPage, hasPreviousPage bool
	currentPage := page

	if filters["id"] != "" || filters["username"] != "" || filters["firstname"] != "" ||
		filters["lastname"] != "" || filters["email"] != "" || filters["idnumber"] != "" || filters["auth"] != "" {
		totalPages = 1
		currentPage = 1
	} else {
		totalPages = (total + pageSize - 1) / pageSize
		hasNextPage = page < totalPages
		hasPreviousPage = page > 1
	}

	response := models.APIResponse{
		Users:           users,
		Total:           total,
		CurrentPage:     currentPage,
		PageSize:        pageSize,
		TotalPages:      totalPages,
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
