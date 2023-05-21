package ioc

// Method for sending an email to someone.
type IEmailService interface {
	Send(to, subject, content, htmlContent string) error
}
