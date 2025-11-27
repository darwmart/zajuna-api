package handlers

import (
	"net/http"
	"zajunaApi/internal/dto/mapper"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services"

	"github.com/gin-gonic/gin"
)

// CategoryHandler maneja las solicitudes relacionadas con categorías
type CategoryHandler struct {
	service services.CategoryServiceInterface
}

// NewCategoryHandler constructor para inyectar el servicio
func NewCategoryHandler(service services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// GetCategories obtiene la lista de categorías
// @Summary      Listar categorías
// @Description  Obtiene todas las categorías de cursos disponibles
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.CategoryListResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	// 1. Llamar al servicio
	categories, err := h.service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"FETCH_ERROR",
			"Error al obtener las categorías",
			err.Error(),
		))
		return
	}

	// 2. Convertir modelos a DTOs
	categoriesResponse := mapper.CategoriesToResponse(categories)

	// 3. Crear respuesta
	listResponse := response.CategoryListResponse{
		Categories: categoriesResponse,
	}

	// 4. Responder
	c.JSON(http.StatusOK, listResponse)
}

// MoveCategory mueve una categoría antes de otra categoría especificada
// @Summary      Mover categoría
// @Description  Mueve una categoría antes de otra categoría especificada (beforeid=0 mueve al final)
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        request body request.MoveCategoryRequest true "Datos de la categoría a mover"
// @Success      200  {object}  response.MoveCategoryResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /categories/move [post]
// MoveCategory mueve una categoría dentro del árbol jerárquico
// Soporta:
//   - Reordenamiento entre categorías hermanas (cambio de sortorder)
//   - Cambio de padre (cambio de parentid)
func (h *CategoryHandler) MoveCategory(c *gin.Context) {
	// 1. Parsear y validar request
	var req request.MoveCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Validación adicional
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"VALIDATION_ERROR",
			err.Error(),
			nil,
		))
		return
	}

	// 3. Llamar al servicio
	if err := h.service.MoveCategory(req.ID, req.BeforeID, req.ParentID); err != nil {
		// Verificar si es error de "no encontrado"
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.NewErrorResponse(
				"CATEGORY_NOT_FOUND",
				"Categoría o padre no encontrado",
				err.Error(),
			))
			return
		}

		// Verificar si es error de validación (categorías con diferente parent)
		if err.Error() == "invalid value" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				"INVALID_OPERATION",
				"La categoría beforeid debe tener el mismo padre que el nuevo padre especificado",
				err.Error(),
			))
			return
		}

		// Otros errores de base de datos
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"MOVE_FAILED",
			"Error al mover la categoría",
			err.Error(),
		))
		return
	}

	// 4. Construir respuesta
	message := "Categoría movida correctamente"
	if req.BeforeID == 0 {
		message = "Categoría movida al final correctamente"
	}
	if req.ParentID != nil {
		message = "Categoría movida y reasignada correctamente"
	}

	// Determinar el nuevo padre para la respuesta
	var newParent uint
	if req.ParentID != nil {
		newParent = *req.ParentID
	}

	c.JSON(http.StatusOK, response.MoveCategoryResponse{
		Message:    message,
		CategoryID: req.ID,
		NewParent:  newParent,
	})
}

// CreateCategories crea una o más categorías siguiendo las reglas de Moodle 4.3
// @Summary      Crear categorías
// @Description  Crea una o más categorías de cursos con jerarquía automática
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        request body request.CreateCategoriesRequest true "Datos de las categorías a crear"
// @Success      200  {object}  response.CategoryListResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /categories [post]
func (h *CategoryHandler) CreateCategories(c *gin.Context) {
	// 1. Parsear y validar request
	var req request.CreateCategoriesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"INVALID_JSON",
			"JSON inválido o campos requeridos faltantes",
			err.Error(),
		))
		return
	}

	// 2. Convertir request a modelos
	var categories []models.Category
	for _, catReq := range req.Categories {
		category := models.Category{
			Name:              catReq.Name,
			Parent:            catReq.Parent,
			IDNumber:          catReq.IDNumber,
			Description:       catReq.Description,
			DescriptionFormat: catReq.DescriptionFormat,
			Theme:             catReq.Theme,
		}
		categories = append(categories, category)
	}

	// 3. Llamar al servicio
	createdCategories, err := h.service.CreateCategories(categories)
	if err != nil {
		// Verificar si es error de "no encontrado" (padre no existe)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.NewErrorResponse(
				"PARENT_NOT_FOUND",
				"La categoría padre especificada no existe",
				err.Error(),
			))
			return
		}

		// Otros errores de base de datos
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"CREATE_FAILED",
			"Error al crear las categorías",
			err.Error(),
		))
		return
	}

	// 4. Obtener el listado completo actualizado de categorías después de crear
	allCategories, err := h.service.GetCategories()
	if err != nil {
		// Si falla al obtener la lista completa, devolver solo las creadas
		categoriesResponse := mapper.CategoriesToResponse(createdCategories)
		c.JSON(http.StatusOK, response.CategoryListResponse{
			Categories: categoriesResponse,
		})
		return
	}

	// 5. Convertir todas las categorías a DTOs
	categoriesResponse := mapper.CategoriesToResponse(allCategories)

	// 6. Crear respuesta con la lista completa
	listResponse := response.CategoryListResponse{
		Categories: categoriesResponse,
	}

	// 7. Responder con la lista completa para que el frontend pueda actualizar su estado
	c.JSON(http.StatusOK, listResponse)
}
