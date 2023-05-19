package errors

import "errors"

var (
	ErrDbContextNotInitialized = errors.New("database context was not initialized")
)
