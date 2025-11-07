package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Tests para GetCourses
func TestGetCourses_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	expectedCourses := []models.Course{
		{ID: 1, FullName: "Course 1", ShortName: "C1", Category: 1},
		{ID: 2, FullName: "Course 2", ShortName: "C2", Category: 1},
	}

	mockService.On("GetAllCourses").Return(expectedCourses, nil)

	router := gin.New()
	router.GET("/courses", handler.GetCourses)

	req, _ := http.NewRequest(http.MethodGet, "/courses", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var courseResponse response.CourseListResponse
	err := json.Unmarshal(w.Body.Bytes(), &courseResponse)
	assert.NoError(t, err)
	assert.NotNil(t, courseResponse.Courses)
	assert.Equal(t, 2, len(courseResponse.Courses))
	assert.Equal(t, "Course 1", courseResponse.Courses[0].FullName)

	mockService.AssertExpectations(t)
}

func TestGetCourses_WithCategoryFilter(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	expectedCourses := []models.Course{
		{ID: 1, FullName: "Course 1", ShortName: "C1", Category: 5},
	}

	mockService.On("GetCoursesByCategory", uint(5)).Return(expectedCourses, nil)

	router := gin.New()
	router.GET("/courses", handler.GetCourses)

	req, _ := http.NewRequest(http.MethodGet, "/courses?categoryid=5", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var courseResponse response.CourseListResponse
	err := json.Unmarshal(w.Body.Bytes(), &courseResponse)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(courseResponse.Courses))

	mockService.AssertExpectations(t)
}

func TestGetCourses_EmptyList(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	emptyCourses := []models.Course{}
	mockService.On("GetAllCourses").Return(emptyCourses, nil)

	router := gin.New()
	router.GET("/courses", handler.GetCourses)

	req, _ := http.NewRequest(http.MethodGet, "/courses", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var courseResponse response.CourseListResponse
	err := json.Unmarshal(w.Body.Bytes(), &courseResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(courseResponse.Courses))

	mockService.AssertExpectations(t)
}

func TestGetCourses_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	expectedError := errors.New("database connection error")
	mockService.On("GetAllCourses").Return(nil, expectedError)

	router := gin.New()
	router.GET("/courses", handler.GetCourses)

	req, _ := http.NewRequest(http.MethodGet, "/courses", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "FETCH_ERROR", errorResponse.Code)
	assert.Equal(t, "Error al obtener los cursos", errorResponse.Message)

	mockService.AssertExpectations(t)
}

func TestGetCourses_InvalidParams(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	router := gin.New()
	router.GET("/courses", handler.GetCourses)

	req, _ := http.NewRequest(http.MethodGet, "/courses?categoryid=-1", nil)
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

// Tests para GetCourseDetails
func TestGetCourseDetails_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	idNumber := "TC001"
	expectedDetails := &repository.CourseDetails{
		ID:               1,
		FullName:         "Test Course",
		ShortName:        "TC1",
		IDNumber:         idNumber,
		Format:           "weeks",
		Category:         "Test Category",
		Groupings:        2,
		Groups:           3,
		RoleAssignments:  map[string]int64{"student": 10, "teacher": 2},
		EnrollmentMethod: "manual",
		Sections:         []string{"Section 1", "Section 2"},
	}

	mockService.On("GetCourseDetails", idNumber).Return(expectedDetails, nil)

	router := gin.New()
	router.GET("/courses/:idnumber/details", handler.GetCourseDetails)

	req, _ := http.NewRequest(http.MethodGet, "/courses/TC001/details", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var detailsResponse response.CourseDetailResponse
	err := json.Unmarshal(w.Body.Bytes(), &detailsResponse)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), detailsResponse.ID)
	assert.Equal(t, "Test Course", detailsResponse.FullName)
	assert.Equal(t, int64(2), detailsResponse.Groupings)

	mockService.AssertExpectations(t)
}

func TestGetCourseDetails_NotFound(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	idNumber := "NONEXISTENT"
	expectedError := errors.New("course not found")
	mockService.On("GetCourseDetails", idNumber).Return(nil, expectedError)

	router := gin.New()
	router.GET("/courses/:idnumber/details", handler.GetCourseDetails)

	req, _ := http.NewRequest(http.MethodGet, "/courses/NONEXISTENT/details", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "NOT_FOUND", errorResponse.Code)
	assert.Equal(t, "Curso no encontrado", errorResponse.Message)

	mockService.AssertExpectations(t)
}

func TestGetCourseDetails_InvalidID(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	router := gin.New()
	router.GET("/courses/:id/details", handler.GetCourseDetails)

	req, _ := http.NewRequest(http.MethodGet, "/courses/invalid/details", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse response.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, "INVALID_ID", errorResponse.Code)
}

// Tests para DeleteCourses
func TestDeleteCourses_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courseid": []int{2, 3, 4},
	}

	serviceResponse := &models.DeleteCoursesResponse{
		Warnings: []models.Warning{},
	}

	mockService.On("DeleteCourses", []int{2, 3, 4}).Return(serviceResponse, nil)

	router := gin.New()
	router.DELETE("/courses", handler.DeleteCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var deleteResponse response.DeleteCoursesResponse
	err := json.Unmarshal(w.Body.Bytes(), &deleteResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Operación completada", deleteResponse.Message)
	assert.Equal(t, 3, deleteResponse.Deleted)
	assert.Equal(t, 0, len(deleteResponse.Warnings))

	mockService.AssertExpectations(t)
}

func TestDeleteCourses_WithWarnings(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courseid": []int{2, 3},
	}

	serviceResponse := &models.DeleteCoursesResponse{
		Warnings: []models.Warning{
			{Item: "course", ItemID: 3, WarningCode: "cannotdeletecourse", Message: "Cannot delete course 3"},
		},
	}

	mockService.On("DeleteCourses", []int{2, 3}).Return(serviceResponse, nil)

	router := gin.New()
	router.DELETE("/courses", handler.DeleteCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var deleteResponse response.DeleteCoursesResponse
	err := json.Unmarshal(w.Body.Bytes(), &deleteResponse)
	assert.NoError(t, err)
	assert.Equal(t, 1, deleteResponse.Deleted)
	assert.Equal(t, 1, len(deleteResponse.Warnings))

	mockService.AssertExpectations(t)
}

func TestDeleteCourses_ValidationError_SiteCourse(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courseid": []int{1}, // ID 1 is the site course
	}

	router := gin.New()
	router.DELETE("/courses", handler.DeleteCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/courses", bytes.NewBuffer(body))
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
}

func TestDeleteCourses_InvalidJSON(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	router := gin.New()
	router.DELETE("/courses", handler.DeleteCourses)

	req, _ := http.NewRequest(http.MethodDelete, "/courses", bytes.NewBuffer([]byte("invalid json")))
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

func TestDeleteCourses_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courseid": []int{2, 3},
	}

	expectedError := errors.New("database delete failed")
	mockService.On("DeleteCourses", []int{2, 3}).Return(nil, expectedError)

	router := gin.New()
	router.DELETE("/courses", handler.DeleteCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodDelete, "/courses", bytes.NewBuffer(body))
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

// Tests para UpdateCourses
func TestUpdateCourses_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courses": []map[string]interface{}{
			{
				"id":       2,
				"fullname": "Updated Course Name",
			},
		},
	}

	serviceResponse := &models.UpdateCoursesResponse{
		Warnings: []models.Warning{},
	}

	mockService.On("UpdateCourses", []request.UpdateCourseRequest{
		{ID: 2, FullName: "Updated Course Name"},
	}).Return(serviceResponse, nil)

	router := gin.New()
	router.PUT("/courses", handler.UpdateCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var updateResponse response.UpdateCoursesResponse
	err := json.Unmarshal(w.Body.Bytes(), &updateResponse)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(updateResponse.Warnings))

	mockService.AssertExpectations(t)
}

func TestUpdateCourses_WithWarnings(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courses": []map[string]interface{}{
			{"id": 2, "fullname": "Course A"},
			{"id": 999}, // Este generará un warning
		},
	}

	serviceResponse := &models.UpdateCoursesResponse{
		Warnings: []models.Warning{
			{Item: "course", ItemID: 999, WarningCode: "updatefailed", Message: "Failed to update course 999"},
		},
	}

	mockService.On("UpdateCourses", []request.UpdateCourseRequest{
		{ID: 2, FullName: "Course A"},
		{ID: 999},
	}).Return(serviceResponse, nil)

	router := gin.New()
	router.PUT("/courses", handler.UpdateCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var updateResponse response.UpdateCoursesResponse
	err := json.Unmarshal(w.Body.Bytes(), &updateResponse)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(updateResponse.Warnings))

	mockService.AssertExpectations(t)
}

func TestUpdateCourses_ValidationError_SiteCourse(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courses": []map[string]interface{}{
			{"id": 1, "fullname": "Try to update site course"},
		},
	}

	router := gin.New()
	router.PUT("/courses", handler.UpdateCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/courses", bytes.NewBuffer(body))
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
}

func TestUpdateCourses_InvalidJSON(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	router := gin.New()
	router.PUT("/courses", handler.UpdateCourses)

	req, _ := http.NewRequest(http.MethodPut, "/courses", bytes.NewBuffer([]byte("invalid json")))
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

func TestUpdateCourses_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.MockCourseService)
	handler := NewCourseHandler(mockService)

	requestBody := map[string]interface{}{
		"courses": []map[string]interface{}{
			{"id": 2, "fullname": "Test"},
		},
	}

	expectedError := errors.New("database update failed")
	mockService.On("UpdateCourses", []request.UpdateCourseRequest{
		{ID: 2, FullName: "Test"},
	}).Return(nil, expectedError)

	router := gin.New()
	router.PUT("/courses", handler.UpdateCourses)

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPut, "/courses", bytes.NewBuffer(body))
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

func TestNewCourseHandler(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockCourseService)

	// Act
	handler := NewCourseHandler(mockService)

	// Assert
	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.service)
}
