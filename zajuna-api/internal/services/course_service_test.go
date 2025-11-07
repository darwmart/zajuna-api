package services

import (
	"errors"
	"testing"
	"zajunaApi/internal/dto/request"
	"zajunaApi/internal/models"
	"zajunaApi/internal/repository"
	"zajunaApi/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllCourses_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	expectedCourses := []models.Course{
		{ID: 1, FullName: "Course 1", ShortName: "C1"},
		{ID: 2, FullName: "Course 2", ShortName: "C2"},
	}

	mockRepo.On("GetAllCourses").Return(expectedCourses, nil)

	// Act
	result, err := service.GetAllCourses()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Course 1", result[0].FullName)
	assert.Equal(t, "Course 2", result[1].FullName)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCourses_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	expectedError := errors.New("database error")
	mockRepo.On("GetAllCourses").Return(nil, expectedError)

	// Act
	result, err := service.GetAllCourses()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCoursesByCategory_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	categoryID := uint(5)
	expectedCourses := []models.Course{
		{ID: 1, FullName: "Course 1", ShortName: "C1", Category: int(categoryID)},
	}

	mockRepo.On("GetCoursesByCategory", categoryID).Return(expectedCourses, nil)

	// Act
	result, err := service.GetCoursesByCategory(categoryID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, int(categoryID), result[0].Category)
	mockRepo.AssertExpectations(t)
}

func TestGetCoursesByCategory_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	categoryID := uint(5)
	expectedError := errors.New("database error")
	mockRepo.On("GetCoursesByCategory", categoryID).Return(nil, expectedError)

	// Act
	result, err := service.GetCoursesByCategory(categoryID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCourseDetails_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	idNumber := "COURSE-001"
	expectedDetails := &repository.CourseDetails{
		ID:        10,
		FullName:  "Test Course",
		ShortName: "TC",
		IDNumber:  idNumber,
		Category:  "Category 1",
		Groups:    5,
		Groupings: 2,
		RoleAssignments: map[string]int64{
			"student": 20,
			"teacher": 2,
		},
	}

	mockRepo.On("GetCourseDetails", idNumber).Return(expectedDetails, nil)

	// Act
	result, err := service.GetCourseDetails(idNumber)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Course", result.FullName)
	assert.Equal(t, int64(5), result.Groups)
	assert.Equal(t, int64(20), result.RoleAssignments["student"])
	mockRepo.AssertExpectations(t)
}

func TestGetCourseDetails_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	idNumber := "NONEXISTENT"
	expectedError := errors.New("course not found")
	mockRepo.On("GetCourseDetails", idNumber).Return(nil, expectedError)

	// Act
	result, err := service.GetCourseDetails(idNumber)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCourses_Success_NoWarnings(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	courseIDs := []int{2, 3, 4}
	mockRepo.On("DeleteCourses", courseIDs).Return([]models.Warning{}, nil)

	// Act
	result, err := service.DeleteCourses(courseIDs)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Warnings))
	mockRepo.AssertExpectations(t)
}

func TestDeleteCourses_Success_WithWarnings(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	courseIDs := []int{2, 999}
	warnings := []models.Warning{
		{
			Item:        "course",
			ItemID:      999,
			WarningCode: "invalidcourseid",
			Message:     "Course ID not found in the database",
		},
	}

	mockRepo.On("DeleteCourses", courseIDs).Return(warnings, nil)

	// Act
	result, err := service.DeleteCourses(courseIDs)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Warnings))
	assert.Equal(t, "invalidcourseid", result.Warnings[0].WarningCode)
	assert.Equal(t, 999, result.Warnings[0].ItemID)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCourses_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	courseIDs := []int{2, 3}
	expectedError := errors.New("database error")
	mockRepo.On("DeleteCourses", courseIDs).Return(nil, expectedError)

	// Act
	result, err := service.DeleteCourses(courseIDs)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCourses_Success_NoWarnings(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	categoryID := 5
	courses := []request.UpdateCourseRequest{
		{
			ID:         2,
			FullName:   "Updated Course Name",
			ShortName:  "UCN",
			CategoryID: &categoryID,
		},
	}

	// Mock espera cualquier map de updates que contenga los campos correctos
	mockRepo.On("UpdateCourse", 2, mock.MatchedBy(func(updates map[string]interface{}) bool {
		return updates["fullname"] == "Updated Course Name" &&
			updates["shortname"] == "UCN" &&
			updates["category"] == 5
	})).Return(nil)

	// Act
	result, err := service.UpdateCourses(courses)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Warnings))
	mockRepo.AssertExpectations(t)
}

func TestUpdateCourses_NoFieldsToUpdate(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	courses := []request.UpdateCourseRequest{
		{
			ID: 2,
			// No fields to update
		},
	}

	// Act
	result, err := service.UpdateCourses(courses)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Warnings))
	assert.Equal(t, "nofieldstoupdate", result.Warnings[0].WarningCode)
	assert.Equal(t, 2, result.Warnings[0].ItemID)
	assert.Contains(t, result.Warnings[0].Message, "No fields provided")
	mockRepo.AssertNotCalled(t, "UpdateCourse")
}

func TestUpdateCourses_UpdateFailed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	courses := []request.UpdateCourseRequest{
		{
			ID:       999,
			FullName: "Non-existent Course",
		},
	}

	updateError := errors.New("course not found")
	mockRepo.On("UpdateCourse", 999, mock.Anything).Return(updateError)

	// Act
	result, err := service.UpdateCourses(courses)

	// Assert
	assert.NoError(t, err) // El servicio no retorna error, solo warnings
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Warnings))
	assert.Equal(t, "updatefailed", result.Warnings[0].WarningCode)
	assert.Equal(t, 999, result.Warnings[0].ItemID)
	assert.Contains(t, result.Warnings[0].Message, "Failed to update")
	mockRepo.AssertExpectations(t)
}

func TestUpdateCourses_MultipleCourses_MixedResults(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	categoryID := 3
	courses := []request.UpdateCourseRequest{
		{
			ID:        2,
			FullName:  "Course Success",
			ShortName: "CS",
		},
		{
			ID: 3,
			// No fields - should generate warning
		},
		{
			ID:         999,
			CategoryID: &categoryID,
			// Should fail to update - generates warning
		},
	}

	mockRepo.On("UpdateCourse", 2, mock.Anything).Return(nil)
	mockRepo.On("UpdateCourse", 999, mock.Anything).Return(errors.New("not found"))

	// Act
	result, err := service.UpdateCourses(courses)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.Warnings))

	// Verify warning types
	warningCodes := make(map[string]int)
	for _, w := range result.Warnings {
		warningCodes[w.WarningCode]++
	}
	assert.Equal(t, 1, warningCodes["nofieldstoupdate"])
	assert.Equal(t, 1, warningCodes["updatefailed"])

	mockRepo.AssertExpectations(t)
}

func TestUpdateCourses_AllFieldsUpdate(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockCourseRepository)
	service := NewCourseService(mockRepo)

	categoryID := 5
	summaryFormat := 1
	startDate := int64(1609459200)
	endDate := int64(1640995200)
	visible := 1

	courses := []request.UpdateCourseRequest{
		{
			ID:            2,
			FullName:      "Complete Update",
			ShortName:     "CU",
			CategoryID:    &categoryID,
			IDNumber:      "ID123",
			Summary:       "Course summary",
			SummaryFormat: &summaryFormat,
			Format:        "topics",
			StartDate:     &startDate,
			EndDate:       &endDate,
			Visible:       &visible,
		},
	}

	mockRepo.On("UpdateCourse", 2, mock.MatchedBy(func(updates map[string]interface{}) bool {
		return updates["fullname"] == "Complete Update" &&
			updates["shortname"] == "CU" &&
			updates["category"] == 5 &&
			updates["idnumber"] == "ID123" &&
			updates["summary"] == "Course summary" &&
			updates["summaryformat"] == 1 &&
			updates["format"] == "topics" &&
			updates["startdate"] == int64(1609459200) &&
			updates["enddate"] == int64(1640995200) &&
			updates["visible"] == 1
	})).Return(nil)

	// Act
	result, err := service.UpdateCourses(courses)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Warnings))
	mockRepo.AssertExpectations(t)
}
