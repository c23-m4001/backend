package constant

import "errors"

var (
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrForbidden        = errors.New("forbidden")

	ErrNoData                = errors.New("no data")
	ErrDuplicateData         = errors.New("duplicate data")
	ErrForeignKeyViolation   = errors.New("data in use or id doesn't exist")
	ErrRequestIdNotFound     = errors.New("request id not found")
	ErrRequestActionNotFound = errors.New("request action not found")
)
