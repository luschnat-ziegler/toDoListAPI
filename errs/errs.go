/*
 * package: errs
 * --------------------
 * Includes definitions of types representing custom errors.
 */

package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

/*
 * Method: AppError.AsMessage
 * --------------------
 * Instantiates an AppError with the Message of the provided AppError
 * and the Code being reset to the zero value. This is needed for
 * serialization.
 *
 * returns: a pointer to an AppError with a zero value (int 0) code.
 */

func (appError AppError) AsMessage() *AppError {
	return &AppError{
		Message: appError.Message,
	}
}

/*
 * Function: NewNotFoundError
 * --------------------
 * Instantiates an AppError with the provided message and code 404.
 *
 * message: a string providing information on the error.
 *
 * returns: a pointer to an AppError.
 */

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

/*
 * Function: NewInternalError
 * --------------------
 * Instantiates an AppError with the provided message and code 500.
 *
 * message: a string providing information on the error.
 *
 * returns: a pointer to an AppError.
 */

func NewInternalError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

/*
 * Function: NewBadRequestError
 * --------------------
 * Instantiates an AppError with the provided message and code 400.
 *
 * message: a string providing information on the error.
 *
 * returns: a pointer to an AppError.
 */

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

/*
 * Function: NewValidationError
 * --------------------
 * Instantiates a Validation error with the provided field info and code 400.
 *
 * invalidFields: a map with field names and violation information.
 *
 * returns: a pointer to a ValidationError.
 */

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
