package errors

import (
	"fmt"
	"net/http"
)

// Error returned from a M3O service handler, the status code is used by the M3O API to determine
// the status code of the response.
type Error struct {
	// Service which first returned the error
	Service string
	// StatusCode which indicates the type of error
	StatusCode int
	// Message provides additional information about an error
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v: %v", http.StatusText(e.StatusCode), e.Message)
}

// BadRequest error indicates that the server cannot or will not process the request due to something
// that is perceived to be a client error (e.g., malformed request syntax)
func BadRequest(format string, a ...interface{}) *Error {
	return &Error{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf(format, a...),
	}
}

// InternalServerError indicates that the server encountered an unexpected condition that prevented
// it from fulfilling the request.
func InternalServerError(format string, a ...interface{}) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintf(format, a...),
	}
}
