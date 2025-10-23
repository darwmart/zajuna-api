package mapper

import (
	"fmt"
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/models"
)

// UserToResponse convierte un modelo User a UserResponse
func UserToResponse(user *models.User) *response.UserResponse {
	if user == nil {
		return nil
	}

	fullName := fmt.Sprintf("%s %s", user.FirstName, user.LastName)

	return &response.UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		FullName:          fullName,
		Email:             user.Email,
		Address:           user.Address,
		Phone1:            user.Phone1,
		Phone2:            user.Phone2,
		Department:        user.Department,
		Institution:       user.Institution,
		IDNumber:          user.IDNumber,
		Interests:         user.Interests,
		FirstAccess:       user.FirstAccess,
		LastAccess:        user.LastAccess,
		Auth:              user.Auth,
		Suspended:         user.Suspended,
		Confirmed:         user.Confirmed,
		Lang:              user.Lang,
		Theme:             user.Theme,
		Timezone:          user.Timezone,
		MailFormat:        user.MailFormat,
		Description:       user.Description,
		DescriptionFormat: user.DescriptionFormat,
		City:              user.City,
		Country:           user.Country,
		ProfileImageSmall: user.ProfileImageSmall,
		ProfileImage:      user.ProfileImage,
		CustomFields:      mapUserCustomFields(user.CustomFields),
		Preferences:       mapUserPreferences(user.Preferences),
	}
}

// UsersToResponse convierte un slice de Users a slice de UserResponse
func UsersToResponse(users []models.User) []response.UserResponse {
	responses := make([]response.UserResponse, len(users))
	for i, user := range users {
		resp := UserToResponse(&user)
		if resp != nil {
			responses[i] = *resp
		}
	}
	return responses
}

// mapUserCustomFields convierte custom fields del modelo al DTO
func mapUserCustomFields(fields []models.UserCustomField) []response.UserCustomField {
	if fields == nil {
		return nil
	}

	result := make([]response.UserCustomField, len(fields))
	for i, field := range fields {
		result[i] = response.UserCustomField{
			Type:         field.Type,
			Value:        field.Value,
			DisplayValue: field.DisplayValue,
			Name:         field.Name,
			ShortName:    field.ShortName,
		}
	}
	return result
}

// mapUserPreferences convierte preferencias del modelo al DTO
func mapUserPreferences(prefs []models.UserPreference) []response.UserPreference {
	if prefs == nil {
		return nil
	}

	result := make([]response.UserPreference, len(prefs))
	for i, pref := range prefs {
		result[i] = response.UserPreference{
			Name:  pref.Name,
			Value: pref.Value,
		}
	}
	return result
}
