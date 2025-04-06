package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("database server failed to execute query")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
	ErrUnauthorized        = errors.New("unauthorized")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrInvalidToken,
	ErrForbidden,
}
