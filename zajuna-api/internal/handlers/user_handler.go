package handlers

import (
	"net/http"
	"strconv"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

type DeleteUsersRequest struct {
	UserIDs []int `json:"userids"`
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	// Extraer parámetros de la query
	filters := map[string]string{
		"firstname": c.Query("firstname"),
		"lastname":  c.Query("lastname"),
		"username":  c.Query("username"),
		"email":     c.Query("email"),
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "15"))

	users, total, err := h.service.GetUsers(filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	c.JSON(http.StatusOK, gin.H{
		"users":       users,
		"total":       total,
		"page":        page,
		"total_pages": totalPages,
	})

}

// DELETE /api/users
func (h *UserHandler) DeleteUsers(c *gin.Context) {
	var req struct {
		UserIDs []int `json:"userids"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	if len(req.UserIDs) == 0 {
		c.JSON(400, gin.H{"error": "Debe proporcionar al menos un ID de usuario"})
		return
	}

	if err := h.service.DeleteUsers(req.UserIDs); err != nil {
		c.JSON(500, gin.H{"error": "Error eliminando usuarios", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuarios suspendidos correctamente"})
}
