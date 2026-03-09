package errors

func NewNotFoundError(module string, message string) *CustomError {
	return &CustomError{
		Module:  module,
		Message: message,
		Params:  nil,
		Status:  404,
	}
}

// IsNotFound returns true if the error is a *CustomError with status 404.
// For other error types (e.g. ent.NotFoundError), check in your application.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	customErr, ok := err.(*CustomError)
	if !ok {
		return false
	}
	return customErr.Status == 404
}
