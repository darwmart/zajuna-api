package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// --- TESTS ---
func TestRequireAuth_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockSessionsRepository)
	mockSession := &models.Sessions{
		ID:  1,
		SID: "valid-token",
	}

	mockRepo.On("FindBySID", "valid-token").Return(mockSession, nil)

	router := gin.New()
	router.Use(RequireAuth(mockRepo))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
	mockRepo.AssertExpectations(t)
}

func TestRequireAuth_NoCookie(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockSessionsRepository)
	router := gin.New()
	router.Use(RequireAuth(mockRepo))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRepo.AssertNotCalled(t, "FindBySID")
}

func TestRequireAuth_DBError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockSessionsRepository)
	mockRepo.On("FindBySID", "valid-token").Return(nil, errors.New("error in db"))

	router := gin.New()
	router.Use(RequireAuth(mockRepo))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestRequireAuth_NoSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockRepo := new(mocks.MockSessionsRepository)
	mockRepo.On("FindBySID", "invalid-token").Return(nil, nil)

	r.Use(RequireAuth(mockRepo))
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "invalid-token"})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRepo.AssertExpectations(t)
}
