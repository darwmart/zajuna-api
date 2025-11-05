package services

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"
	"zajunaApi/internal/services/auth"
	authMocks "zajunaApi/internal/services/mocks"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	os.Setenv("SECRET", "testsecret")
	os.Setenv("ENROL_USER_ACTIVE", "1")
	os.Setenv("CONTEXT_COURSE", "50")
	os.Setenv("SITEID", "1")
}

// ============================================================================
// Tests para GetUsers
// ============================================================================

func TestGetUsers_Success_NoFilters(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	filters := map[string]string{}
	page := 1
	limit := 10

	expectedUsers := []models.User{
		{ID: 1, Username: "user1", Email: "user1@example.com", FirstName: "John", LastName: "Doe"},
		{ID: 2, Username: "user2", Email: "user2@example.com", FirstName: "Jane", LastName: "Smith"},
	}
	expectedTotal := int64(2)

	mockRepo.On("FindByFilters", filters, page, limit).Return(expectedUsers, expectedTotal, nil)

	// Act
	users, total, err := service.GetUsers(filters, page, limit)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, "user1", users[0].Username)
	assert.Equal(t, "user2", users[1].Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Success_WithFilters(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	filters := map[string]string{
		"firstname": "John",
		"email":     "john@example.com",
	}
	page := 1
	limit := 20

	expectedUsers := []models.User{
		{ID: 1, Username: "johndoe", Email: "john@example.com", FirstName: "John", LastName: "Doe"},
	}
	expectedTotal := int64(1)

	mockRepo.On("FindByFilters", filters, page, limit).Return(expectedUsers, expectedTotal, nil)

	// Act
	users, total, err := service.GetUsers(filters, page, limit)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, expectedTotal, total)
	assert.Equal(t, "John", users[0].FirstName)
	assert.Equal(t, "john@example.com", users[0].Email)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Success_EmptyResult(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	filters := map[string]string{"username": "nonexistent"}
	page := 1
	limit := 10

	emptyUsers := []models.User{}
	expectedTotal := int64(0)

	mockRepo.On("FindByFilters", filters, page, limit).Return(emptyUsers, expectedTotal, nil)

	// Act
	users, total, err := service.GetUsers(filters, page, limit)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 0, len(users))
	assert.Equal(t, int64(0), total)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Success_Pagination(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	filters := map[string]string{}
	page := 2
	limit := 5

	expectedUsers := []models.User{
		{ID: 6, Username: "user6", Email: "user6@example.com"},
		{ID: 7, Username: "user7", Email: "user7@example.com"},
		{ID: 8, Username: "user8", Email: "user8@example.com"},
		{ID: 9, Username: "user9", Email: "user9@example.com"},
		{ID: 10, Username: "user10", Email: "user10@example.com"},
	}
	expectedTotal := int64(50) // Total de usuarios en la BD

	mockRepo.On("FindByFilters", filters, page, limit).Return(expectedUsers, expectedTotal, nil)

	// Act
	users, total, err := service.GetUsers(filters, page, limit)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, int64(50), total)
	assert.Equal(t, "user6", users[0].Username)
	assert.Equal(t, "user10", users[4].Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	filters := map[string]string{}
	page := 1
	limit := 10

	expectedError := errors.New("database connection error")
	mockRepo.On("FindByFilters", filters, page, limit).Return(nil, int64(0), expectedError)

	// Act
	users, total, err := service.GetUsers(filters, page, limit)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Equal(t, int64(0), total)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_MultipleFilters(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	filters := map[string]string{
		"firstname": "John",
		"lastname":  "Doe",
		"email":     "john.doe@example.com",
		"username":  "johndoe",
	}
	page := 1
	limit := 10

	expectedUsers := []models.User{
		{
			ID:        1,
			Username:  "johndoe",
			Email:     "john.doe@example.com",
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	expectedTotal := int64(1)

	mockRepo.On("FindByFilters", filters, page, limit).Return(expectedUsers, expectedTotal, nil)

	// Act
	users, total, err := service.GetUsers(filters, page, limit)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, int64(1), total)
	mockRepo.AssertExpectations(t)
}

// ============================================================================
// Tests para DeleteUsers
// ============================================================================

func TestDeleteUsers_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	userIDs := []int{2, 3, 4}
	mockRepo.On("DeleteUsers", userIDs).Return(nil)

	// Act
	err := service.DeleteUsers(userIDs)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUsers_SingleUser(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	userIDs := []int{5}
	mockRepo.On("DeleteUsers", userIDs).Return(nil)

	// Act
	err := service.DeleteUsers(userIDs)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUsers_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	userIDs := []int{2, 3}
	expectedError := errors.New("database error")
	mockRepo.On("DeleteUsers", userIDs).Return(expectedError)

	// Act
	err := service.DeleteUsers(userIDs)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUsers_EmptyList(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	userIDs := []int{}
	mockRepo.On("DeleteUsers", userIDs).Return(nil)

	// Act
	err := service.DeleteUsers(userIDs)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// ============================================================================
// Tests para UpdateUsers
// ============================================================================

func TestUpdateUsers_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	users := []models.User{
		{
			ID:        2,
			FirstName: "John Updated",
			LastName:  "Doe Updated",
			Email:     "john.updated@example.com",
			City:      "New York",
			Country:   "US",
		},
	}

	expectedCount := int64(1)
	mockRepo.On("UpdateUsers", users).Return(expectedCount, nil)

	// Act
	count, err := service.UpdateUsers(users)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_MultipleUsers(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	users := []models.User{
		{ID: 2, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{ID: 3, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
		{ID: 4, FirstName: "Bob", LastName: "Johnson", Email: "bob@example.com"},
	}

	expectedCount := int64(3)
	mockRepo.On("UpdateUsers", users).Return(expectedCount, nil)

	// Act
	count, err := service.UpdateUsers(users)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(3), count)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	users := []models.User{
		{ID: 999, FirstName: "NonExistent", Email: "nonexistent@example.com"},
	}

	expectedError := errors.New("user not found")
	mockRepo.On("UpdateUsers", users).Return(int64(0), expectedError)

	// Act
	count, err := service.UpdateUsers(users)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, int64(0), count)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_PartialSuccess(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	users := []models.User{
		{ID: 2, FirstName: "John", Email: "john@example.com"},
		{ID: 3, FirstName: "Jane", Email: "jane@example.com"},
		{ID: 999, FirstName: "Invalid", Email: "invalid@example.com"}, // Este fallaría
	}

	// Simula que solo 2 fueron actualizados exitosamente
	expectedCount := int64(2)
	expectedError := errors.New("partial update: user 999 not found")
	mockRepo.On("UpdateUsers", users).Return(expectedCount, expectedError)

	// Act
	count, err := service.UpdateUsers(users)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, int64(2), count)
	assert.Contains(t, err.Error(), "partial update")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_EmptyList(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	users := []models.User{}
	mockRepo.On("UpdateUsers", users).Return(int64(0), nil)

	// Act
	count, err := service.UpdateUsers(users)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUsers_AllFieldsUpdate(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	users := []models.User{
		{
			ID:        2,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			City:      "San Francisco",
			Country:   "US",
			Lang:      "en",
			Timezone:  "America/Los_Angeles",
			Phone1:    "+1-555-1234",
		},
	}

	expectedCount := int64(1)
	mockRepo.On("UpdateUsers", mock.MatchedBy(func(u []models.User) bool {
		return len(u) == 1 &&
			u[0].FirstName == "John" &&
			u[0].Email == "john.doe@example.com" &&
			u[0].City == "San Francisco" &&
			u[0].Phone1 == "+1-555-1234"
	})).Return(expectedCount, nil)

	// Act
	count, err := service.UpdateUsers(users)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
	mockRepo.AssertExpectations(t)
}

// ============================================================================
// Tests para Constructor
// ============================================================================

func TestNewUserService(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
}

// ✅ Test de Login exitoso
func TestUserService_Login_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	userService := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	authMock := new(authMocks.MockAuthPlugin)

	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	hashedPassword, _ := HashPassword("password123")

	mockUser := &models.User{
		ID:        1,
		Username:  "testuser",
		Password:  hashedPassword,
		Suspended: 0,
		Auth:      "manual",
	}

	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(2, nil)
	sessionRepoMock.On("InsertSession", mock.AnythingOfType("*models.Sessions")).Return(nil)
	authMock.On("Login", "testuser", "password123").Return(true)

	// Mock del plugin manual (autenticación)

	token, err := userService.Login(req, "testuser", "password123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
	sessionRepoMock.AssertExpectations(t)
}

// ❌ Usuario no encontrado
func TestUserService_Login_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	req := httptest.NewRequest(http.MethodPost, "/login", nil)

	mockRepo.On("FindByUsername", "unknown").Return(nil, nil)

	token, err := service.Login(req, "unknown", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "usuario no encontrado")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

// ❌ Usuario suspendido
func TestUserService_Login_SuspendedUser(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	req := httptest.NewRequest(http.MethodPost, "/login", nil)

	mockUser := &models.User{ID: 1, Username: "testuser", Suspended: 1, Auth: "manual"}
	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Suspended Login")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

// ❌ Sin cursos vinculados
func TestUserService_Login_NoCourses(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	hashedPassword, _ := HashPassword("password123")

	req := httptest.NewRequest(http.MethodPost, "/login", nil)

	mockUser := &models.User{ID: 1, Username: "testuser", Password: hashedPassword, Suspended: 0, Auth: "manual"}
	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(0, nil)

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "El usuario no tiene cursos vinculados")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
}

// ❌ Contraseña incorrecta
func TestUserService_Login_InvalidPassword(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	hashedPassword, _ := HashPassword("password123")

	mockUser := &models.User{ID: 1, Username: "testuser", Password: hashedPassword, Suspended: 0, Auth: "manual"}
	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(2, nil)

	token, err := service.Login(req, "testuser", "wrongpass")

	assert.Error(t, err)
	assert.EqualError(t, err, "credenciales inválidas")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
}

// ❌ Error insertando sesión
func TestUserService_Login_InsertSessionError(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	//AuthPlugin

	userService := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	hashedPassword, _ := HashPassword("password123")

	req := httptest.NewRequest(http.MethodPost, "/login", nil)

	mockUser := &models.User{
		ID:        1,
		Username:  "testuser",
		Password:  hashedPassword,
		Suspended: 0,
		Auth:      "manual",
	}
	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(2, nil)
	sessionRepoMock.On("InsertSession", mock.Anything).Return(errors.New("insert failed"))

	token, err := userService.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "insert failed")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	sessionRepoMock.AssertExpectations(t)
}

// ✅ Logout exitoso
func TestUserService_Logout_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)

	sessionRepoMock.On("DeleteSession", "abc123").Return(nil)

	msg, err := service.Logout("abc123")

	assert.NoError(t, err)
	assert.Equal(t, "Sesion deleted", msg)
	sessionRepoMock.AssertExpectations(t)
}

// ❌ Error en logout
func TestUserService_Logout_Error(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	// Act
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	sessionRepoMock.On("DeleteSession", "badtoken").Return(errors.New("db error"))

	msg, err := service.Logout("badtoken")

	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
	assert.Empty(t, msg)
	sessionRepoMock.AssertExpectations(t)
}

// ❌ Error al buscar usuario
func TestUserService_Login_FindByUsernameError(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)

	mockRepo.On("FindByUsername", "testuser").Return(nil, errors.New("db error"))

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

// ❌ Método de autenticación no encontrado
func TestUserService_Login_AuthPluginNotFound(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	hashedPassword, _ := HashPassword("password123")

	mockUser := &models.User{
		ID:        1,
		Username:  "testuser",
		Password:  hashedPassword,
		Suspended: 0,
		Auth:      "inexistente",
	}

	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "método de autenticación no encontrado")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
}

// ❌ Error al contar cursos
func TestUserService_Login_CountUserCoursesError(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	hashedPassword, _ := HashPassword("password123")

	mockUser := &models.User{
		ID:        1,
		Username:  "testuser",
		Password:  hashedPassword,
		Suspended: 0,
		Auth:      "manual",
	}

	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(0, errors.New("db error"))

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "Error al buscar los cursos del usuario")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
}

// ❌ Error o fallo en el plugin de autenticación
func TestUserService_Login_AuthPluginError(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	authMock := new(authMocks.MockAuthPlugin)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	hashedPassword, _ := HashPassword("password123")

	mockUser := &models.User{
		ID:        1,
		Username:  "testuser",
		Password:  hashedPassword,
		Suspended: 0,
		Auth:      "manual",
	}

	// Configuración de mocks
	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(2, nil)

	// Simulamos que el plugin de autenticación falla
	authMock.On("Login", "testuser", "password123").Return(false, errors.New("auth failed"))

	// Aquí depende de cómo implementas auth.Get(), si usas un registro global, asegúrate
	// de registrar el mock para el tipo "manual".
	auth.Register("manual", authMock)

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.EqualError(t, err, "credenciales inválidas")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
	authMock.AssertExpectations(t)
}
func TestUserService_Login_TokenSigningError(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	sessionRepoMock := new(mocks.MockSessionsRepository)
	courseRepoMock := new(mocks.MockCourseRepository)
	authMock := new(authMocks.MockAuthPlugin)

	service := NewUserService(mockRepo, sessionRepoMock, courseRepoMock)
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	hashedPassword, _ := HashPassword("password123")

	mockUser := &models.User{
		ID:        1,
		Username:  "testuser",
		Password:  hashedPassword,
		Suspended: 0,
		Auth:      "manual",
	}

	mockRepo.On("FindByUsername", "testuser").Return(mockUser, nil)
	courseRepoMock.On("CountUserCourses", int(mockUser.ID)).Return(1, nil)
	authMock.On("Login", "testuser", "password123").Return(true, nil)
	auth.Register("manual", authMock)

	// Mock del signer para forzar el error
	originalSigner := signToken
	defer func() { signToken = originalSigner }() // restaurar después
	signToken = func(token *jwt.Token, secret string) (string, error) {
		return "", errors.New("fallo al firmar token (simulado)")
	}

	token, err := service.Login(req, "testuser", "password123")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Error al firmar token")
	assert.Empty(t, token)

	mockRepo.AssertExpectations(t)
	courseRepoMock.AssertExpectations(t)
	authMock.AssertExpectations(t)
}
func TestNormalizeLoopback(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"IPv6 loopback", "::1", "127.0.0.1"},
		{"IPv4 address", "192.168.1.10", "192.168.1.10"},
		{"Empty IP", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeLoopback(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestGetRemoteAddr_WithXForwardedFor(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.5")

	ip := getRemoteAddr(req)
	assert.Equal(t, "203.0.113.5", ip)
}
func TestGetRemoteAddr_WithInvalidXForwardedFor_NoRemoteAddr(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Forwarded-For", "invalid-ip")
	req.RemoteAddr = "" // forzamos que no haya TCP IP válida

	ip := getRemoteAddr(req)
	assert.Equal(t, "0.0.0.0", ip)
}
func TestPasswordVerify_CryptErrorByMock(t *testing.T) {
	// guardamos la función original y la restauramos al final
	original := cryptFunc
	defer func() { cryptFunc = original }()

	// sustituimos cryptFunc por una que siempre falle
	cryptFunc = func(password, salt string) (string, error) {
		return "", errors.New("forced error")
	}

	ok := PasswordVerify("any", "anyhash")
	assert.False(t, ok, "si cryptFunc devuelve error PasswordVerify debe devolver false")
}
