package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ============================================================================
// Tests para GetUsers
// ============================================================================

func TestGetUsers_Success_NoFilters(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

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
	service := NewUserService(mockRepo)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
}
