package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CompetencyHandler struct {
	service services.CompetencyServiceInterface
}

func NewCompetencyHandler(service services.CompetencyServiceInterface) *CompetencyHandler {
	return &CompetencyHandler{service: service}
}

func (h *CompetencyHandler) CreateCompetency(c *gin.Context) {
	// 1. Parsear request
	var req request.CreateCompetencyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Convertir DTO a modelo
	competency := &models.Competency{
		ShortName:             req.ShortName,
		Description:           req.Description,
		DescriptionFormat:     req.DescriptionFormat,
		IDNumber:              req.IDNumber,
		CompetencyFrameworkID: req.CompetencyFrameworkID,
		ParentID:              req.ParentID,
		Path:                  "/0/",
		SortOrder:             0,
		RuleType:              req.RuleType,
		RuleOutcome:           req.RuleOutcome,
		RuleConfig:            req.RuleConfig,
		ScaleID:               req.ScaleID,
		ScaleConfiguration:    req.ScaleConfiguration,
	}

	// Obtener el token
	sid, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 3. Llamar a la capa de servicio
	competency, err = h.service.CreateCompetency(sid, competency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"SERVER_ERROR",
			"Error interno del servidor",
			err.Error(),
		))
		return
	}

	// 4. Retornar respuesta exitosa
	c.JSON(http.StatusOK, response.CreateCompetencyResponse{
		ID:       competency.ID,
		FullName: competency.ShortName,
		Message:  "Competency created successfully",
	})
}
