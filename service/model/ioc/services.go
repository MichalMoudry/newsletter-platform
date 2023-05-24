package ioc

// Method for sending an email to someone.
type IEmailService interface {
	// Method for sending an email (as Newsletter platform) to someone.
	Send(to, subject, content, htmlContent string) error
	SendBatch(to []string, subject, content, htmlContent string) error
}
