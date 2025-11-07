package request

import "fmt"

// GetEnrolledUsersRequest representa la solicitud para obtener usuarios matriculados
type GetEnrolledUsersRequest struct {
	CourseID int `uri:"courseid" binding:"required"`
}

// EnrolledUsersOptions representa las opciones de filtrado
type EnrolledUsersOptions struct {
	WithCapability  string `form:"withcapability"`
	GroupID         int    `form:"groupid"`
	OnlyActive      int    `form:"onlyactive"`
	OnlySuspended   int    `form:"onlysuspended"`
	UserFields      string `form:"userfields"`
	LimitFrom       int    `form:"limitfrom"`
	LimitNumber     int    `form:"limitnumber"`
	SortBy          string `form:"sortby"`          // id, firstname, lastname, siteorder
	SortDirection   string `form:"sortdirection"`   // ASC or DESC
}

// SetDefaults establece valores por defecto
func (o *EnrolledUsersOptions) SetDefaults() {
	if o.SortBy == "" {
		o.SortBy = "id"
	}
	if o.SortDirection == "" {
		o.SortDirection = "ASC"
	}
	if o.LimitNumber == 0 {
		o.LimitNumber = 100 // Límite por defecto
	}
}

// Validate valida las opciones
func (o *EnrolledUsersOptions) Validate() error {
	// No se puede usar onlyactive y onlysuspended al mismo tiempo
	if o.OnlyActive == 1 && o.OnlySuspended == 1 {
		return fmt.Errorf("onlyactive and onlysuspended cannot be used together")
	}
	return nil
}
