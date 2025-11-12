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

func TestHasCapability_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ARRANGE
	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	// Simular respuestas de ConfigRepository
	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)

	// Simular sesión válida
	mockSessionRepo.On("FindBySID", "valid-token").Return(&models.Sessions{UserID: 1}, nil)

	// Simular permisos (uno con CAP_ALLOW)
	mockRoleCapabilityRepo.On("FindByUserID", int64(1), []string{"2", "3"}, "mod/test:view").
		Return(&[]models.RoleCapability{
			{Permission: CAP_ALLOW},
		}, nil)

	// Crear router con middleware
	router := gin.New()
	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	// Crear request con cookie válida
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})
	w := httptest.NewRecorder()

	// ACT
	router.ServeHTTP(w, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Access granted")

	mockConfigRepo.AssertExpectations(t)
	mockSessionRepo.AssertExpectations(t)
	mockRoleCapabilityRepo.AssertExpectations(t)
}

func TestHasCapability_Unauthorized_NoPermission(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)
	mockSessionRepo.On("FindBySID", "valid-token").Return(&models.Sessions{UserID: 1}, nil)

	// Devuelve permisos sin CAP_ALLOW
	mockRoleCapabilityRepo.On("FindByUserID", int64(1), []string{"2", "3"}, "mod/test:view").
		Return(&[]models.RoleCapability{
			{Permission: CAP_PREVENT},
		}, nil)

	router := gin.New()
	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestHasCapability_MissingCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)

	router := gin.New()
	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestHasCapability_DBErrorFirstConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(nil, errors.New("error in db"))
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)

	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHasCapability_DBErrorSecondConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(nil, errors.New("error in db"))

	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHasCapability_DBErrorSearchingSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)
	mockSessionRepo.On("FindBySID", "valid-token").Return(nil, errors.New("error in db"))

	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
func TestHasCapability_DBErrorSearchingCapabilities(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)
	mockSessionRepo.On("FindBySID", "valid-token").Return(&models.Sessions{UserID: 1}, nil)

	mockRoleCapabilityRepo.On("FindByUserID", int64(1), []string{"2", "3"}, "mod/test:view").Return(nil, errors.New("error in db"))

	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
func TestHasCapability_ProhibitRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockConfigRepo := new(mocks.MockConfigRepository)
	mockSessionRepo := new(mocks.MockSessionsRepository)
	mockRoleCapabilityRepo := new(mocks.MockRoleCapabilityRepository)

	// Mock configuración
	mockConfigRepo.On("FindByName", "defaultuserroleid").Return(&models.Config{Value: "2"}, nil)
	mockConfigRepo.On("FindByName", "defaultfrontpageroleid").Return(&models.Config{Value: "3"}, nil)

	// Mock sesión válida
	mockSessionRepo.On("FindBySID", "valid-token").Return(&models.Sessions{UserID: 10}, nil)

	// Mock capabilities con prohibición
	mockRoleCapabilityRepo.On("FindByUserID", int64(10), []string{"2", "3"}, "mod/test:view").
		Return(&[]models.RoleCapability{
			{Permission: CAP_PROHIBIT},
		}, nil)

	// Middleware + ruta
	router.Use(HasCapability(mockConfigRepo, mockSessionRepo, mockRoleCapabilityRepo, "mod/test:view"))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid-token"})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRoleCapabilityRepo.AssertExpectations(t)
}
