package errors

import "errors"

var (
	ErrEmailClientNotInitialized = errors.New("email client was not initialized")
	ErrInvalidHash               = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion       = errors.New("incompatible version of argon2")
	ErrJwtParsing                = errors.New("unable to parse JWT token")
	ErrInvalidEmail              = errors.New("invalid email in the URL")
	ErrInvalidLogin              = errors.New("invalid login information")
)
