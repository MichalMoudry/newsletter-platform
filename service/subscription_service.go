package service

import (
	"context"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service/model/ioc"

	"github.com/google/uuid"
)

type SubscriptionService struct {
	SubscriptionRepo ioc.ISubscriptionRepository
	SubscriberRepo   ioc.ISubscriberRepository
}

func NewSubscriptionService(
	subscriptionRepo ioc.ISubscriptionRepository,
	subscriberRepo ioc.ISubscriberRepository,
) SubscriptionService {
	return SubscriptionService{
		SubscriptionRepo: subscriptionRepo,
		SubscriberRepo:   subscriberRepo,
	}
}

// Method for creating a new subscription in the system.
func (srvc SubscriptionService) CreateSubscription(ctx context.Context, email string, newsletterId uuid.UUID) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = database.EndTransaction(tx, err) }()

	subscriberInfo, err := srvc.SubscriberRepo.GetSubscriber(email)
	if err != nil {
		return err
	}

	if subscriberInfo.Email == "" {
		subscriber := model.NewSubscriber(email)
		if err = srvc.SubscriberRepo.AddSubscriber(subscriber); err != nil {
			return err
		}
	}

	err = srvc.SubscriptionRepo.AddSubscription(model.NewNewsletterSubscription(newsletterId, email))
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
