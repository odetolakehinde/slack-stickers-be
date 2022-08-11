package helper

import "errors"

var (
	// ErrRecordNotFound if the record is not found in the database
	ErrRecordNotFound = errors.New("record not found")
	// ErrRecordCreatingFailed if error occurred while trying to insert record into the database
	ErrRecordCreatingFailed = errors.New("record failed to insert")
	// ErrRecordUpdateFailed if error occurred in attempt to update the row
	ErrRecordUpdateFailed = errors.New("record update failed")
	// ErrDeleteFailed if error occurred in an attempt to delete a record from the database
	ErrDeleteFailed = errors.New("failed to delete record")
	// ErrInvalidResponse when the response cannot be interpreted
	ErrInvalidResponse = errors.New("invalid response")
	// ErrServiceUnsupported when service is currently unsupported by the provider
	ErrServiceUnsupported = errors.New("service currently unsupported")
	// ErrProviderMisConfigured when 3rd party provider is mis configured
	ErrProviderMisConfigured = errors.New("provider to process request mis configured")
	// ErrGinContextRetrieveFailed could not retrieve gin.Context
	ErrGinContextRetrieveFailed = errors.New("could not retrieve gin.Context")
	// ErrGinContextWrongType gin.Context has wrong type
	ErrGinContextWrongType = errors.New("gin.Context has wrong type")
)
