package handlers

import (
	"net/http"
	"strconv"

	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service *services.CourseService
}

func NewCourseHandler(service *services.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

// GET /api/courses?categoryid=# (parámetro opcional)
func (h *CourseHandler) GetCourses(c *gin.Context) {
	categoryIDStr := c.Query("categoryid")

	// filtramos por categoría
	if categoryIDStr != "" {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "categoryid debe ser un número válido"})
			return
		}

		courses, err := h.service.GetCoursesByCategory(uint(categoryID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los cursos: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"courses": courses})
		return
	}

	// todos los cursos
	courses, err := h.service.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los cursos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

/*func (h *CourseHandler) GetCourseRoles(c *gin.Context) {
idParam := c.Param("id")
courseID, err := strconv.Atoi(idParam)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
	return
}

roles, err := h.service.GetCourseRoles(courseID)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}

c.JSON(http.StatusOK, roles)}*/

func (h *CourseHandler) GetCourseDetails(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	details, err := h.service.GetCourseDetails(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}

// DeleteCourses maneja la eliminación de cursos
func (h *CourseHandler) DeleteCourses(c *gin.Context) {
	var req models.DeleteCoursesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	response, err := h.service.DeleteCourses(req.CourseIDs)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting courses", "details": err.Error()})
		return
	}

	c.JSON(200, response)
}
