package service

import (
	"context"
	"fmt"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service/model/email"
	"newsletter-platform/service/model/ioc"

	"github.com/google/uuid"
)

type SubscriptionService struct {
	SubscriptionRepo ioc.ISubscriptionRepository
	SubscriberRepo   ioc.ISubscriberRepository
	NewsletterRepo   ioc.INewsletterRepository
	EmailService     ioc.IEmailService
}

func NewSubscriptionService(
	subscriptionRepo ioc.ISubscriptionRepository,
	subscriberRepo ioc.ISubscriberRepository,
	newletterRepo ioc.INewsletterRepository,
	emailSrvc ioc.IEmailService,
) SubscriptionService {
	return SubscriptionService{
		SubscriptionRepo: subscriptionRepo,
		SubscriberRepo:   subscriberRepo,
		NewsletterRepo:   newletterRepo,
		EmailService:     emailSrvc,
	}
}

// Method for creating a new subscription in the system.
func (srvc SubscriptionService) CreateSubscription(ctx context.Context, subEmail string, newsletterId uuid.UUID) (err error) {
	newsletterData, err := srvc.NewsletterRepo.GetNewsletter(newsletterId)
	if err != nil {
		return err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = database.EndTransaction(tx, err)
		if err == nil {
			err = srvc.EmailService.Send(
				subEmail,
				email.NewSubscriptionEmailSubject,
				fmt.Sprintf(email.NewSubscriptionEmailContent, newsletterData.Nltr_Name),
				fmt.Sprintf(email.NewSubscriptionEmailHtmlContent, newsletterData.Nltr_Name, newsletterId),
			)
		}
	}()

	subscriberInfo, err := srvc.SubscriberRepo.GetSubscriber(subEmail)
	if err != nil {
		return err
	}

	if subscriberInfo.Email == "" {
		subscriber := model.NewSubscriber(subEmail)
		if err = srvc.SubscriberRepo.AddSubscriber(subscriber); err != nil {
			return err
		}
	}

	err = srvc.SubscriptionRepo.AddSubscription(model.NewNewsletterSubscription(newsletterId, subEmail))
	if err != nil {
		return err
	}
	return nil
}

// Method for cancelling a subscription in the system.
func (srvc SubscriptionService) CancelSubscription(ctx context.Context, email string, newsletterId uuid.UUID) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = database.EndTransaction(tx, err) }()

	if err = srvc.SubscriptionRepo.DeleteSubscription(email, newsletterId); err != nil {
		return err
	}
	// TODO: add case when subscriber is subscribed to another newsletter.
	return nil
}
