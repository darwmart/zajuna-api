package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ============================================================================
// Tests para GetUsers
// ============================================================================

func TestGetUsers_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	expectedUsers := []models.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@test.com", Username: "johndoe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@test.com", Username: "janesmith"},
	}

	mockService.On("GetUsers", map[string]string{}, 1, 15).Return(expectedUsers, int64(2), nil)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var paginatedResponse response.PaginatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &paginatedResponse)
	assert.NoError(t, err)
	assert.NotNil(t, paginatedResponse.Data)
	assert.Equal(t, int64(2), paginatedResponse.Pagination.Total)

	mockService.AssertExpectations(t)
}

func TestGetUsers_WithFilters(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	expectedUsers := []models.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@test.com"},
	}

	filters := map[string]string{"firstname": "John"}
	mockService.On("GetUsers", filters, 1, 15).Return(expectedUsers, int64(1), nil)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	req, _ := http.NewRequest(http.MethodGet, "/users?firstname=John", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var paginatedResponse response.PaginatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &paginatedResponse)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), paginatedResponse.Pagination.Total)

	mockService.AssertExpectations(t)
}

func TestGetUsers_EmptyList(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	emptyUsers := []models.User{}
	mockService.On("GetUsers", map[string]string{}, 1, 15).Return(emptyUsers, int64(0), nil)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var paginatedResponse response.PaginatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &paginatedResponse)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), paginatedResponse.Pagination.Total)

	mockService.AssertExpectations(t)
}

func TestGetUsers_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	expectedError := errors.New("database connection error")
	mockService.On("GetUsers", map[string]string{}, 1, 15).Return(nil, int64(0), expectedError)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "FETCH_ERROR", errorResponse.Code)
	assert.Equal(t, "Error al obtener usuarios", errorResponse.Message)

	mockService.AssertExpectations(t)
}

func TestGetUsers_Pagination(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	users := []models.User{
		{ID: 11, FirstName: "User11", LastName: "Test"},
		{ID: 12, FirstName: "User12", LastName: "Test"},
	}

	mockService.On("GetUsers", map[string]string{}, 2, 10).Return(users, int64(25), nil)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	req, _ := http.NewRequest(http.MethodGet, "/users?page=2&limit=10", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var paginatedResponse response.PaginatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &paginatedResponse)
	assert.NoError(t, err)
	assert.Equal(t, 2, paginatedResponse.Pagination.Page)
	assert.Equal(t, int64(25), paginatedResponse.Pagination.Total)

	mockService.AssertExpectations(t)
}

// ============================================================================
// Tests para UpdateUsers
// ============================================================================

func TestUpdateUsers_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	requestBody := map[string]interface{}{
		"users": []map[string]interface{}{
			{
				"id":        1,
				"firstname": "UpdatedJohn",
				"lastname":  "UpdatedDoe",
				"email":     "updated@test.com",
			},
		},
	}

	mockService.On("UpdateUsers", []models.User{
		{ID: 1, FirstName: "UpdatedJohn", LastName: "UpdatedDoe", Email: "updated@test.com"},
	}).Return(int64(1), nil)

	router := gin.New()
	router.PUT("/users/update", handler.UpdateUsers)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/users/update", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var updateResponse response.UpdateUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &updateResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Usuarios actualizados correctamente", updateResponse.Message)
	assert.Equal(t, int64(1), updateResponse.Updated)

	mockService.AssertExpectations(t)
}

func TestUpdateUsers_EmptyRequest(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	requestBody := map[string]interface{}{
		"users": []map[string]interface{}{},
	}

	router := gin.New()
	router.PUT("/users/update", handler.UpdateUsers)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/users/update", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	// La validación binding:"required,min=1" de Gin rechaza arrays vacíos
	assert.Equal(t, "INVALID_JSON", errorResponse.Code)
}

func TestUpdateUsers_InvalidJSON(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.PUT("/users/update", handler.UpdateUsers)

	req, _ := http.NewRequest(http.MethodPut, "/users/update", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "INVALID_JSON", errorResponse.Code)
}

func TestUpdateUsers_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	requestBody := map[string]interface{}{
		"users": []map[string]interface{}{
			{
				"id":        1,
				"firstname": "Test",
				"lastname":  "User",
				"email":     "test@test.com",
			},
		},
	}

	expectedError := errors.New("database update failed")
	mockService.On("UpdateUsers", []models.User{
		{ID: 1, FirstName: "Test", LastName: "User", Email: "test@test.com"},
	}).Return(int64(0), expectedError)

	router := gin.New()
	router.PUT("/users/update", handler.UpdateUsers)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/users/update", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "UPDATE_FAILED", errorResponse.Code)

	mockService.AssertExpectations(t)
}

// ============================================================================
// Tests para DeleteUsers
// ============================================================================

func TestDeleteUsers_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	requestBody := map[string]interface{}{
		"userids": []int{1, 2, 3},
	}

	mockService.On("DeleteUsers", []int{1, 2, 3}).Return(nil)

	router := gin.New()
	router.DELETE("/users", handler.DeleteUsers)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var deleteResponse response.DeleteUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &deleteResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Usuarios suspendidos correctamente", deleteResponse.Message)
	assert.Equal(t, 3, deleteResponse.Deleted)

	mockService.AssertExpectations(t)
}

func TestDeleteUsers_InvalidJSON(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.DELETE("/users", handler.DeleteUsers)

	req, _ := http.NewRequest(http.MethodDelete, "/users", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "INVALID_JSON", errorResponse.Code)
}

func TestDeleteUsers_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	requestBody := map[string]interface{}{
		"userids": []int{1, 2},
	}

	expectedError := errors.New("database delete failed")
	mockService.On("DeleteUsers", []int{1, 2}).Return(expectedError)

	router := gin.New()
	router.DELETE("/users", handler.DeleteUsers)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "DELETE_FAILED", errorResponse.Code)

	mockService.AssertExpectations(t)
}

func TestDeleteUsers_DuplicateIDs(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	requestBody := map[string]interface{}{
		"userids": []int{1, 2, 1}, // ID duplicado
	}

	router := gin.New()
	router.DELETE("/users", handler.DeleteUsers)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "VALIDATION_ERROR", errorResponse.Code)
	assert.Contains(t, errorResponse.Message, "duplicado")
}

func TestGetUsers_InvalidQueryParams(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	// Parámetros inválidos: page negativo
	req, _ := http.NewRequest(http.MethodGet, "/users?page=-1", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "INVALID_PARAMS", errorResponse.Code)
}

func TestNewUserHandler(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockUserService)

	// Act
	handler := NewUserHandler(mockService)

	// Assert
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.service)
}
func TestUserHandler_Login_InvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/login", handler.Login)

	req := httptest.NewRequest("POST", "/login", strings.NewReader("{invalid json"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestUserHandler_Login_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/login", handler.Login)

	body := `{"Username":"user", "Password":"pass"}`

	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	mockService.
		On("Login", req, "user", "pass").
		Return("", errors.New("bad credentials"))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}
func TestUserHandler_Login_EmptyToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/login", handler.Login)

	body := `{"Username":"user", "Password":"pass"}`

	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	mockService.
		On("Login", req, "user", "pass").
		Return("", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}
func TestUserHandler_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/login", handler.Login)

	body := `{"Username":"user","Password":"pass"}`

	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	mockToken := "abc123"

	mockService.
		On("Login", req, "user", "pass").
		Return(mockToken, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	cookie := w.Result().Cookies()
	assert.NotEmpty(t, cookie)
	assert.Equal(t, "Authorization", cookie[0].Name)
	assert.Equal(t, mockToken, cookie[0].Value)

	mockService.AssertExpectations(t)
}
func TestUserHandler_Logout_NoCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/logout", handler.Logout)

	req := httptest.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
func TestUserHandler_Logout_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/logout", handler.Logout)

	req := httptest.NewRequest("POST", "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "token123"})

	mockService.
		On("Logout", "token123").
		Return("", errors.New("logout failed"))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}
func TestUserHandler_Logout_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.POST("/logout", handler.Logout)

	req := httptest.NewRequest("POST", "/logout", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "token123"})

	mockService.
		On("Logout", "token123").
		Return("Sesion deleted", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	cookie := w.Result().Cookies()
	assert.NotEmpty(t, cookie)
	assert.Equal(t, -1, cookie[0].MaxAge)
	assert.Equal(t, "", cookie[0].Value)

	mockService.AssertExpectations(t)
}
