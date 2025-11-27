package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CompetencyPlanHandler struct {
	service services.CompetencyPlanServiceInterface
}

func NewCompetencyPlanHandler(service services.CompetencyPlanServiceInterface) *CompetencyPlanHandler {
	return &CompetencyPlanHandler{service: service}
}

func (h *CompetencyPlanHandler) CreateCompetencyPlan(c *gin.Context) {
	var req request.CreateCompetencyPlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	competencyPlan := &models.CompetencyPlan{
		Name:              req.Name,
		Description:       req.Description,
		DescriptionFormat: req.DescriptionFormat,
		UserID:            req.UserID,
		TemplateID:        req.TemplateID,
		OrigTemplateID:    req.OrigTemplateID,
		Status:            req.Status,
		DueDate:           req.DueDate,
		ReviewerID:        req.ReviewerID,
	}
	// Obtener el token
	sid, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	plan, err := h.service.CreateCompetencyPlan(sid, competencyPlan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"SERVER_ERROR",
			"Error interno del servidor",
			err.Error(),
		))
		return
	}
	// 4. Retornar respuesta exitosa
	c.JSON(http.StatusOK, response.CreateCompetencyPlanResponse{
		ID:      plan.ID,
		Name:    plan.Name,
		Message: "Plan creado correctamente",
	})
}
