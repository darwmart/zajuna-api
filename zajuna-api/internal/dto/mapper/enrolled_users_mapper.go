package mapper

import (
	"zajunaApi/internal/dto/response"
	"zajunaApi/internal/repository"
)

// EnrolledUserDetailToResponse convierte EnrolledUserDetail a EnrolledUserResponse
func EnrolledUserDetailToResponse(
	user *repository.EnrolledUserDetail,
	groups []map[string]interface{},
	roles []map[string]interface{},
	customFields []map[string]interface{},
	preferences []map[string]interface{},
	enrolledCourses []map[string]interface{},
) *response.EnrolledUserResponse {
	if user == nil {
		return nil
	}

	resp := &response.EnrolledUserResponse{
		ID:                   int64(user.ID),
		Username:             user.Username,
		FirstName:            user.FirstName,
		LastName:             user.LastName,
		FullName:             user.FirstName + " " + user.LastName,
		Email:                user.Email,
		Address:              user.Address,
		Phone1:               user.Phone1,
		Phone2:               user.Phone2,
		Department:           user.Department,
		Institution:          user.Institution,
		IDNumber:             user.IDNumber,
		Interests:            user.Interests,
		FirstAccess:          user.FirstAccess,
		LastAccess:           user.LastAccess,
		LastCourseAccess:     user.LastCourseAccess,
		Description:          user.Description,
		DescriptionFormat:    user.DescriptionFormat,
		City:                 user.City,
		Country:              user.Country,
		ProfileImageURLSmall: user.ProfileImageSmall,
		ProfileImageURL:      user.ProfileImage,
	}

	// Mapear custom fields
	if len(customFields) > 0 {
		resp.CustomFields = make([]response.UserCustomField, 0, len(customFields))
		for _, cf := range customFields {
			resp.CustomFields = append(resp.CustomFields, response.UserCustomField{
				Type:      getStringFromMap(cf, "type"),
				Value:     getStringFromMap(cf, "value"),
				Name:      getStringFromMap(cf, "name"),
				ShortName: getStringFromMap(cf, "shortname"),
			})
		}
	}

	// Mapear groups
	if len(groups) > 0 {
		resp.Groups = make([]response.EnrolledUserGroup, 0, len(groups))
		for _, g := range groups {
			resp.Groups = append(resp.Groups, response.EnrolledUserGroup{
				ID:                getIntFromMap(g, "id"),
				Name:              getStringFromMap(g, "name"),
				Description:       getStringFromMap(g, "description"),
				DescriptionFormat: getIntFromMap(g, "descriptionformat"),
			})
		}
	}

	// Mapear roles
	if len(roles) > 0 {
		resp.Roles = make([]response.EnrolledUserRole, 0, len(roles))
		for _, r := range roles {
			resp.Roles = append(resp.Roles, response.EnrolledUserRole{
				RoleID:    int64(getIntFromMap(r, "roleid")),
				Name:      getStringFromMap(r, "name"),
				ShortName: getStringFromMap(r, "shortname"),
				SortOrder: getIntFromMap(r, "sortorder"),
			})
		}
	}

	// Mapear preferences
	if len(preferences) > 0 {
		resp.Preferences = make([]response.UserPreference, 0, len(preferences))
		for _, p := range preferences {
			resp.Preferences = append(resp.Preferences, response.UserPreference{
				Name:  getStringFromMap(p, "name"),
				Value: getStringFromMap(p, "value"),
			})
		}
	}

	// Mapear enrolled courses
	if len(enrolledCourses) > 0 {
		resp.EnrolledCourses = make([]response.UserEnrolledCourse, 0, len(enrolledCourses))
		for _, ec := range enrolledCourses {
			resp.EnrolledCourses = append(resp.EnrolledCourses, response.UserEnrolledCourse{
				ID:        int64(getIntFromMap(ec, "id")),
				FullName:  getStringFromMap(ec, "fullname"),
				ShortName: getStringFromMap(ec, "shortname"),
			})
		}
	}

	return resp
}

// Helper functions para extraer valores de maps
func getStringFromMap(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getIntFromMap(m map[string]interface{}, key string) int {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case int:
			return v
		case int64:
			return int(v)
		case float64:
			return int(v)
		}
	}
	return 0
}
