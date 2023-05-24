package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"

	"github.com/google/uuid"
)

type SubscriptionRepository struct{}

// Method for adding a new subscription to the database.
func (SubscriptionRepository) AddSubscription(sub *model.NewsletterSubscription) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.AddSubscription, sub); err != nil {
		return err
	}
	return nil
}

// Method for deleting a specific subscription in the database.
func (SubscriptionRepository) DeleteSubscription(email string, newsletterId uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}
	if _, err = ctx.Exec(query.DeleteSubscription, email, newsletterId); err != nil {
		return err
	}
	return nil
}

// Method for obtaining list of newsletter subscribers.
func (SubscriptionRepository) GetNewsletterSubscriptions(newsletterId uuid.UUID) ([]string, error) {
	return nil, nil
}
