package service

import (
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
	if srvc.emailClient == nil {
		return errors.ErrEmailClientNotInitialized
	}
	identity := os.Getenv("EMAIL_SENDER_IDENTITY")
	if identity == "" {
		return errors.ErrEmailClientNotInitialized
	}

	from := mail.NewEmail("Newsletter platform", identity)
	recipient := mail.NewEmail("Test recipient", to)
	message := mail.NewSingleEmail(from, subject, recipient, content, htmlContent)
	if response, err := srvc.emailClient.Send(message); err != nil {
		return err
	} else {
		log.Print(response)
		return nil
	}
}
