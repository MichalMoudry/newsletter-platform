package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"
)

type SubscriberRepository struct{}

// Method for adding a new subscriber to the database.
func (SubscriberRepository) AddSubscriber(sub *model.Subscriber) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.AddSubscriber, sub); err != nil {
		return err
	}
	return nil
}
