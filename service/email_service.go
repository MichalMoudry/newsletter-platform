package service

import (
	err_lib "errors"
	"log"
	"newsletter-platform/service/errors"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	emailClient *sendgrid.Client
}

func NewEmailService() *EmailService {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Println("No API key for email service was set.")
		return &EmailService{}
	}
	return &EmailService{
		emailClient: sendgrid.NewSendClient(apiKey),
	}
}

// Method for sending an email (as Newsletter platform) to someone.
func (srvc EmailService) Send(to, subject, content, htmlContent string) error {
	senderId, err := srvc.getSenderId()
	if err != nil {
		return err
	}

	from := mail.NewEmail("Newsletter platform", senderId)
	recipient := mail.NewEmail("Test recipient", to)
	message := mail.NewSingleEmail(from, subject, recipient, content, htmlContent)
	if response, err := srvc.emailClient.Send(message); err != nil {
		return err
	} else {
		log.Print(response)
		return nil
	}
}

// Method for sending an email to multiple recipents.
func (srvc EmailService) SendBatch(to []string, subject, content, htmlContent string) (err error) {
	var errs []error
	for _, recipent := range to {
		errs = append(errs, srvc.Send(recipent, subject, content, htmlContent))
	}

	err = err_lib.Join(errs...)
	if err != nil {
		return err
	}
	return nil
}

// Method for guarding send methods.
func (srvc EmailService) getSenderId() (string, error) {
	if srvc.emailClient == nil {
		return "", errors.ErrEmailClientNotInitialized
	}
	identity := os.Getenv("EMAIL_SENDER_IDENTITY")
	if identity == "" {
		return "", errors.ErrSenderIdentityNotSet
	}

	return identity, nil
}
