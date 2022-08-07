package model

import (
	"errors"
	"fmt"
)

var (
	// ErrEmptyValidIDs if array of ids has no valid id
	ErrEmptyValidIDs = errors.New("no valid id has been submitted")

	// ErrIncompleteDetails when user details is incomplete
	ErrIncompleteDetails = errors.New("incomplete user details")

	// ErrIncompleteDetailsTier when user details is incomplete
	ErrIncompleteDetailsTier = errors.New("incomplete user details - tier")

	// ErrInvalidToken when user details is incomplete
	ErrInvalidToken = errors.New("invalid token")
)

// ErrDynamicInvalidUUID dynamic error for all UUID types.
func ErrDynamicInvalidUUID(uuidType string) error {
	return fmt.Errorf("the %s is not a valid uuid", uuidType)
}
