// Package error provides custom application error
package error

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	// ErrNotFound is not found error
	ErrNotFound = errors.New("error not found")

	// ErrInvalidRequest is bad request error
	ErrInvalidRequest = errors.New("error invalid request")

	// ErrUnauthorized is unauthorized request error
	ErrUnauthorized = errors.New("error unauthorized")

	// ErrPermission is permission error
	ErrPermission = errors.New("error permission")

	// ErrUnprocessable is unprocessable entity error
	ErrUnprocessable = errors.New("error unprocessable")

	// ErrInternalServer is internal server error
	ErrInternalServer = errors.New("error internal server")
)

// IsErrNotFound check err is ErrNotFound
func IsErrNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// NewErrNotFound returns new error with ErrNotFound error type
func NewErrNotFound(message string) error {
	return fmt.Errorf("%s:%w", message, ErrNotFound)
}

// NewErrInvalidRequest returns new error with ErrInvalidRequest error type
func NewErrInvalidRequest(message string) error {
	return fmt.Errorf("%s:%w", message, ErrInvalidRequest)
}

// NewErrUnauthorized returns new error with ErrUnauthorized error type
func NewErrUnauthorized(message string) error {
	return fmt.Errorf("%s:%w", message, ErrUnauthorized)
}

// NewErrPermission returns new error with ErrPermission error type
func NewErrPermission(message string) error {
	return fmt.Errorf("%s:%w", message, ErrPermission)
}

// NewErrUnprocessable returns new error with ErrUnprocessable error type
func NewErrUnprocessable(message string) error {
	return fmt.Errorf("%s:%w", message, ErrUnprocessable)
}

// NewErrInternalServer returns new error with ErrInternalServer error type
func NewErrInternalServer(message string) error {
	return fmt.Errorf("%s:%w", message, ErrInternalServer)
}

// ErrStatusCode returns http status code and error message from error
func ErrStatusCode(err error) (int, string) {
	errMessage := strings.Split(err.Error(), ":")[0]

	switch {

	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, errMessage

	case errors.Is(err, ErrInvalidRequest):
		return http.StatusBadRequest, errMessage

	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized, errMessage

	case errors.Is(err, ErrPermission):
		return http.StatusForbidden, errMessage

	case errors.Is(err, ErrUnprocessable):
		return http.StatusUnprocessableEntity, errMessage

	case errors.Is(err, ErrInternalServer):
		return http.StatusInternalServerError, errMessage

	default:
		return http.StatusInternalServerError, "oops"

	}
}
