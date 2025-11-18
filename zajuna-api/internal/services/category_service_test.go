package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	expectedCategories := []models.Category{
		{
			ID:          1,
			Name:        "Category 1",
			Description: "Description 1",
			SortOrder:   1,
		},
		{
			ID:          2,
			Name:        "Category 2",
			Description: "Description 2",
			SortOrder:   2,
		},
		{
			ID:          3,
			Name:        "Category 3",
			Description: "Description 3",
			SortOrder:   3,
		},
	}

	mockRepo.On("GetAllCategories").Return(expectedCategories, nil)

	// Act
	result, err := service.GetCategories()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "Category 1", result[0].Name)
	assert.Equal(t, "Category 2", result[1].Name)
	assert.Equal(t, "Category 3", result[2].Name)
	assert.Equal(t, 1, result[0].SortOrder)
	assert.Equal(t, 2, result[1].SortOrder)
	assert.Equal(t, 3, result[2].SortOrder)
	mockRepo.AssertExpectations(t)
}

func TestGetCategories_EmptyList(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	emptyCategories := []models.Category{}
	mockRepo.On("GetAllCategories").Return(emptyCategories, nil)

	// Act
	result, err := service.GetCategories()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))
	mockRepo.AssertExpectations(t)
}

func TestGetCategories_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	expectedError := errors.New("database connection error")
	mockRepo.On("GetAllCategories").Return(nil, expectedError)

	// Act
	result, err := service.GetCategories()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	assert.EqualError(t, err, "database connection error")
	mockRepo.AssertExpectations(t)
}

func TestGetCategories_LargeDataset(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	// Simular un dataset grande
	largeCategories := make([]models.Category, 100)
	for i := 0; i < 100; i++ {
		largeCategories[i] = models.Category{
			ID:        uint(i + 1),
			Name:      "Category " + string(rune(i+1)),
			SortOrder: i + 1,
		}
	}

	mockRepo.On("GetAllCategories").Return(largeCategories, nil)

	// Act
	result, err := service.GetCategories()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 100, len(result))
	assert.Equal(t, uint(1), result[0].ID)
	assert.Equal(t, uint(100), result[99].ID)
	mockRepo.AssertExpectations(t)
}

func TestNewCategoryService(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCategoryRepository)

	// Act
	service := NewCategoryService(mockRepo)

	// Assert
	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.repo)
}

