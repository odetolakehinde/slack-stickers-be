package controller

import (
	"errors"
)

var (
	// ErrEmailAlreadyExists when an email already exists in the DB
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	// ErrAdminEmailAlreadyExists when an admin email already exists in the DB
	ErrAdminEmailAlreadyExists = errors.New("admin with this email already exists")
	// ErrEmailAlreadyVerified when an email already exists in the DB
	ErrEmailAlreadyVerified = errors.New("user email address has already been verified")
	// ErrUserDoesNotExist when an email already exists in the DB
	ErrUserDoesNotExist = errors.New("user with this email does not exist")
	// ErrIncorrectLoginDetails when user auth details are incorrect
	ErrIncorrectLoginDetails = errors.New("incorrect login details")
	// ErrIncorrectUserID  when user userID is incorrect
	ErrIncorrectUserID = errors.New("invalid user ID")
	// ErrCannotVerifyGoogleToken when the Google auth token cannot be verified
	ErrCannotVerifyGoogleToken = errors.New("cannot verify google auth token")
	// ErrUserSignedUpWithEmail when the user created an account with email but trying to access with sso
	ErrUserSignedUpWithEmail = errors.New("user created account with email and password")
	// ErrUserSignedUpWithGoogle when the user created an account with google but trying to access with email
	ErrUserSignedUpWithGoogle = errors.New("user created account with google sso")
	// ErrInvalidOtp when the otp code is not a valid one i.e check not passed
	ErrInvalidOtp = errors.New("otp code is not valid")
	// ErrEmailOtpIsExpired when the otp code has expired
	ErrEmailOtpIsExpired = errors.New("otp code has expired")
	// ErrIncorrectPassword when password is incorrect
	ErrIncorrectPassword = errors.New("incorrect password")
	// ErrUnauthorizedUser when a user tries to perform an unauthorized action
	ErrUnauthorizedUser = errors.New("you are not authorized")
	// ErrRatingIsEarly when a user tries to rate too early
	ErrRatingIsEarly = errors.New("you cannot rate this plan yet")
	// ErrCannotCreatePaymentCode when we were unable to create a payment plan
	ErrCannotCreatePaymentCode = errors.New("we were unable to create a payment plan code for the plan")
	//ErrNetworkUnreachable when network is unreachable
	ErrNetworkUnreachable = errors.New("network currently unreachable")
	// ErrFieldsNotComplete when some fields are not added
	ErrFieldsNotComplete = errors.New("some required fields are missing")
)
