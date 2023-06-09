package errors

import "errors"

var (
	ErrEmailClientNotInitialized = errors.New("email client was not initialized")
	ErrSenderIdentityNotSet      = errors.New("EMAIL_SENDER_IDENTITY environment variable has not been set")
	ErrInvalidHash               = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion       = errors.New("incompatible version of argon2")
	ErrJwtParsing                = errors.New("unable to parse JWT token")
	ErrInvalidEmail              = errors.New("invalid email in the URL")
	ErrInvalidLogin              = errors.New("invalid login information")
	ErrPassResetTokenExpired     = errors.New("password reset token expired")
)
