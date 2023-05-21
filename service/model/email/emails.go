package email

import (
	_ "embed"
)

var (
	//go:embed content/PasswordResetEmail.html
	PassResetEmailHtmlContent string
)

const (
	PassResetEmailSubject string = "Newsplatform - password reset"
	PassResetEmailContent string = "" //TODO: Add text version of password reset email
)
