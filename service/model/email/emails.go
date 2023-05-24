package email

import (
	_ "embed"
)

var (
	//go:embed content/PasswordResetEmail.html
	PassResetEmailHtmlContent string
	//go:embed content/PostEmail.html
	PostEmailContent string
	//go:embed content/NewSubscriptionEmail.html
	NewSubscriptionEmailHtmlContent string
)

const (
	PassResetEmailSubject       string = "Newsplatform - password reset"
	PassResetEmailContent       string = "Hello, We have received a request to reset your password. To do so, click the link below. This link is valid for twenty-four hours and can only be used once."
	NewSubscriptionEmailContent string = "You have successfully subscribed to %s newsletter. Unsubscribe from the newsletter"
	NewSubscriptionEmailSubject string = "Newsplatform - new subscription"
)
