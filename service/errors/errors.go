package errors

import "errors"

var (
	ErrEmailClientNotInitialized = errors.New("email client was not initialized")
	ErrInvalidHash               = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion       = errors.New("incompatible version of argon2")
)
