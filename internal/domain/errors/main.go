package errors

import (
	"fmt"
)

type ValidationError struct {
	message string
}

func NewValidationError(message string) error {
	return &ValidationError{message: message}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.message)
}

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) error {
	return &NotFoundError{message: message}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error: %s", e.message)
}
