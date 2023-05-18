package errors

import "errors"

var (
	ErrEmailClientNotInitialized = errors.New("email client was not initialized")
)
