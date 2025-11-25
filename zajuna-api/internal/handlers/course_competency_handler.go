package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CourseCompetencyHandler struct {
	service services.CourseCompetencyServiceInterface
}

func NewCourseCompetencyHandler(service services.CourseCompetencyServiceInterface) *CourseCompetencyHandler {
	return &CourseCompetencyHandler{service: service}
}
func (h *CourseCompetencyHandler) AddCompetencyToCourse(c *gin.Context) {

	var req request.CreateCourseCompetencyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	courseCompetency := &models.CourseCompetency{
		CourseID:     req.CourseID,
		CompetencyID: req.CompetencyID,
		RuleOutcome:  request.OUTCOME_EVIDENCE,
		SortOrder:    0,
	}
	// Obtener el token
	sid, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err = h.service.AddCompetencyToCourse(sid, courseCompetency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"INTERNAL_SERVER_ERROR",
			"Error al agregar la competencia al curso",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, response.NewSuccessResponse("Competencia agregada correctamente al curso", nil))

}
