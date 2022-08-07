// Package model houses all custom/translated core models from the application
// i.e we might not want to accept/return all fields in storage in rest endpoint.
// this is just similar to how we have our graphql messages written out and defined
package model

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateRequest validates a struct with the request body
func ValidateRequest(request interface{}) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			return fmt.Errorf(fieldError{fieldErr}.String())
		}
	}

	return nil
}

type fieldError struct {
	err validator.FieldError
}

// String() returns the string value of the error.
func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
