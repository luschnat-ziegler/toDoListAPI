package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (appError AppError) AsMessage() *AppError {
	return &AppError{
		Message: appError.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

type ValidationError struct {
	Code          int               `json:",omitempty"`
	InvalidFields map[string]string `json:"invalid_fields"`
}

func NewValidationError(invalidFields map[string]string) *ValidationError {
	return &ValidationError{
		InvalidFields: invalidFields,
		Code:          http.StatusBadRequest,
	}
}

func (validationError ValidationError) AsMessage() *ValidationError {
	return &ValidationError{
		InvalidFields: validationError.InvalidFields,
	}
}
