package entities

import "errors"

// API errors
var (
	ErrNotFound      = errors.New("not found")
	ErrForbidden     = errors.New("forbidden")
	ErrBadRequest    = errors.New("bad request")
	ErrInternalError = errors.New("internal server error")
)

// Repository errors
var (
	ErrExecuteQuery  = errors.New("failed to execute query")
	ErrQueryRow      = errors.New("failed to query row")
	ErrScan          = errors.New("failed to scan row")
	ErrExecuteOrScan = errors.New("failed to execute or scan query")
)
