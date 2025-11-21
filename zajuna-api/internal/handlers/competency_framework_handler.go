package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

type CompetencyFrameworkHandler struct {
	service services.CompetencyFrameworkServiceInterface
}

func NewCompetencyFrameworkHandler(service services.CompetencyFrameworkServiceInterface) *CompetencyFrameworkHandler {
	return &CompetencyFrameworkHandler{service: service}
}
func (h *CompetencyFrameworkHandler) CreateCompetencyFramework(c *gin.Context) {
	// 1. Parsear request
	var req request.CreateCompetencyFrameworkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Convertir DTO a modelo (siguiendo el patrón solicitado)
	framework := &models.CompetencyFramework{
		ShortName:          req.ShortName,
		IDNumber:           req.IDNumber,
		Description:        req.Description,
		DescriptionFormat:  req.DescriptionFormat,
		Visible:            req.Visible,
		ScaleID:            req.ScaleID,
		ScaleConfiguration: req.ScaleConfiguration,
		ContextID:          req.ContextID,
		Taxonomies:         req.Taxonomies,
	}

	// Obtener el token
	sid, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 3. Llamar a la capa de servicio
	createdFramework, err := h.service.CreateCompetencyFramework(sid, framework)
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
		ID:       createdFramework.ID,
		FullName: createdFramework.ShortName,
		Message:  "Competency created successfully",
	})
}
