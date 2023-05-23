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

// Method for deleting a specific subscription.
func (SubscriptionRepository) DeleteSubscription(id uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}
	if _, err = ctx.NamedExec(query.DeleteNewsletter, id); err != nil {
		return err
	}
	return nil
}
