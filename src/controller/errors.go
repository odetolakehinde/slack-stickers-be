package controller

import (
	"errors"
)

var (
	// ErrIncorrectLoginDetails when user auth details are incorrect
	ErrIncorrectLoginDetails = errors.New("incorrect login details")
	// ErrIncorrectPassword when password is incorrect
	ErrIncorrectPassword = errors.New("incorrect password")
	// ErrUnauthorizedUser when a user tries to perform an unauthorized action
	ErrUnauthorizedUser = errors.New("you are not authorized")
	//ErrNetworkUnreachable when network is unreachable
	ErrNetworkUnreachable = errors.New("network currently unreachable")
	// ErrFieldsNotComplete when some fields are not added
	ErrFieldsNotComplete = errors.New("some required fields are missing")
)
