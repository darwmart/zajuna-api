package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCategoryService)
	handler := NewCategoryHandler(mockService)

	expectedCategories := []models.Category{
		{ID: 1, Name: "Category 1", Description: "Description 1", SortOrder: 1},
		{ID: 2, Name: "Category 2", Description: "Description 2", SortOrder: 2},
	}

	mockService.On("GetCategories").Return(expectedCategories, nil)

	// Configurar router
	router := gin.New()
	router.GET("/categories", handler.GetCategories)

	// Crear request
	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response response.CategoryListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Categories)
	assert.Equal(t, 2, len(response.Categories))
	assert.Equal(t, "Category 1", response.Categories[0].Name)
	assert.Equal(t, "Category 2", response.Categories[1].Name)

	mockService.AssertExpectations(t)
}

func TestGetCategories_EmptyList(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCategoryService)
	handler := NewCategoryHandler(mockService)

	emptyCategories := []models.Category{}
	mockService.On("GetCategories").Return(emptyCategories, nil)

	router := gin.New()
	router.GET("/categories", handler.GetCategories)

	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response response.CategoryListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Categories)
	assert.Equal(t, 0, len(response.Categories))

	mockService.AssertExpectations(t)
}

func TestGetCategories_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCategoryService)
	handler := NewCategoryHandler(mockService)

	expectedError := errors.New("database connection error")
	mockService.On("GetCategories").Return(nil, expectedError)

	router := gin.New()
	router.GET("/categories", handler.GetCategories)

	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "FETCH_ERROR", errorResponse.Code)
	assert.Equal(t, "Error al obtener las categorías", errorResponse.Message)

	mockService.AssertExpectations(t)
}

func TestGetCategories_LargeDataset(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCategoryService)
	handler := NewCategoryHandler(mockService)

	// Simular un dataset grande
	largeCategories := make([]models.Category, 50)
	for i := 0; i < 50; i++ {
		largeCategories[i] = models.Category{
			ID:        uint(i + 1),
			Name:      "Category",
			SortOrder: i + 1,
		}
	}

	mockService.On("GetCategories").Return(largeCategories, nil)

	router := gin.New()
	router.GET("/categories", handler.GetCategories)

	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response response.CategoryListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 50, len(response.Categories))

	mockService.AssertExpectations(t)
}

func TestGetCategories_JSONFormat(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCategoryService)
	handler := NewCategoryHandler(mockService)

	categories := []models.Category{
		{ID: 1, Name: "Test Category", Description: "Test Description", SortOrder: 1},
	}

	mockService.On("GetCategories").Return(categories, nil)

	router := gin.New()
	router.GET("/categories", handler.GetCategories)

	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	// Verificar que es JSON valido
	var response response.CategoryListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}

func TestNewCategoryHandler(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockCategoryService)

	// Act
	handler := NewCategoryHandler(mockService)

	// Assert
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.service)
}
