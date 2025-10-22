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
	c.SetCookie("Authorization", token, 3600*3, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

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
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
