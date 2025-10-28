package response

import "time"

// ErrorResponse representa una respuesta de error estandarizada
type ErrorResponse struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// SuccessResponse representa una respuesta exitosa genérica
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PaginationMeta contiene la metadata de paginación
type PaginationMeta struct {
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}

// PaginatedResponse representa una respuesta paginada genérica
type PaginatedResponse struct {
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// NewPaginatedResponse crea una nueva respuesta paginada
func NewPaginatedResponse(data interface{}, page, limit int, total int64) *PaginatedResponse {
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &PaginatedResponse{
		Data: data,
		Pagination: PaginationMeta{
			Page:        page,
			Limit:       limit,
			Total:       total,
			TotalPages:  totalPages,
			HasNext:     page < totalPages,
			HasPrevious: page > 1,
		},
	}
}

// NewErrorResponse crea una nueva respuesta de error
func NewErrorResponse(code, message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Code:      code,
		Message:   message,
		Details:   details,
		Timestamp: time.Now(),
	}
}

// NewSuccessResponse crea una nueva respuesta exitosa
func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Data:    data,
	}
}
