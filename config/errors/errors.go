package errors

import "errors"

var (
	ErrMissingConnectionString = errors.New("environment variable CONNECTION_STRING was not set")
)
