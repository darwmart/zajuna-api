package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Init inicializa el validador con validaciones personalizadas
func Init() {
	validate = validator.New()

	// Usar el nombre del tag json como nombre del campo en los errores
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Registrar validaciones personalizadas
	validate.RegisterValidation("moodle_username", validateMoodleUsername)
	validate.RegisterValidation("moodle_shortname", validateMoodleShortname)
}

// Validate valida una estructura
func Validate(s interface{}) error {
	if validate == nil {
		Init()
	}

	err := validate.Struct(s)
	if err != nil {
		return FormatValidationErrors(err)
	}
	return nil
}

// ValidateVar valida una variable individual
func ValidateVar(field interface{}, tag string) error {
	if validate == nil {
		Init()
	}

	err := validate.Var(field, tag)
	if err != nil {
		return FormatValidationErrors(err)
	}
	return nil
}

// FormatValidationErrors formatea los errores de validación
func FormatValidationErrors(err error) error {
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var messages []string
	for _, e := range validationErrors {
		messages = append(messages, formatFieldError(e))
	}

	return &ValidationError{
		Errors: messages,
	}
}

// formatFieldError formatea un error de campo individual
func formatFieldError(e validator.FieldError) string {
	field := e.Field()

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("El campo '%s' es requerido", field)
	case "email":
		return fmt.Sprintf("El campo '%s' debe ser un email válido", field)
	case "min":
		return fmt.Sprintf("El campo '%s' debe tener al menos %s caracteres", field, e.Param())
	case "max":
		return fmt.Sprintf("El campo '%s' no debe exceder %s caracteres", field, e.Param())
	case "len":
		return fmt.Sprintf("El campo '%s' debe tener exactamente %s caracteres", field, e.Param())
	case "oneof":
		return fmt.Sprintf("El campo '%s' debe ser uno de: %s", field, e.Param())
	case "moodle_username":
		return fmt.Sprintf("El campo '%s' contiene caracteres no válidos", field)
	case "moodle_shortname":
		return fmt.Sprintf("El campo '%s' debe ser un nombre corto válido de Moodle", field)
	default:
		return fmt.Sprintf("El campo '%s' es inválido", field)
	}
}

// validateMoodleUsername valida un username
// Solo permite alfanuméricos, guiones, underscores, puntos y @
func validateMoodleUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if len(username) < 2 || len(username) > 100 {
		return false
	}

	// Moodle username: alfanuméricos, -, _, ., @
	for _, char := range username {
		if !isValidMoodleUsernameChar(char) {
			return false
		}
	}
	return true
}

// validateMoodleShortname valida un shortname
func validateMoodleShortname(fl validator.FieldLevel) bool {
	shortname := fl.Field().String()
	if len(shortname) < 1 || len(shortname) > 100 {
		return false
	}
	// No debe contener espacios
	return !strings.Contains(shortname, " ")
}

// isValidMoodleUsernameChar verifica si un carácter es válido para username
func isValidMoodleUsernameChar(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == '-' || char == '_' || char == '.' || char == '@'
}

// ValidationError representa un error de validación
type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Errors, "; ")
}
