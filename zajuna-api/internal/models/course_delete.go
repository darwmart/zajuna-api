package models

// Petición para eliminar cursos
type DeleteCoursesRequest struct {
	CourseIDs []int `json:"courseids" binding:"required"`
}

// Estructura de advertencia (warnings)
type Warning struct {
	Item        string `json:"item,omitempty"`
	ItemID      int    `json:"itemid,omitempty"`
	WarningCode string `json:"warningcode"`
	Message     string `json:"message"`
}

// Respuesta estándar de la API Moodle-like
type DeleteCoursesResponse struct {
	Warnings []Warning `json:"warnings,omitempty"`
}
