package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"

	"github.com/google/uuid"
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

// Method for obtaining information about a subscriber from the database.
func (SubscriberRepository) GetSubscriber(email string) (model.SubscriberInformation, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.SubscriberInformation{}, err
	}

	var data model.SubscriberInformation
	rows, err := ctx.Query(query.GetSubscriber, email)
	if err != nil {
		return model.SubscriberInformation{}, err
	}
	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Email)
	}
	if err != nil {
		return model.SubscriberInformation{}, err
	}

	return data, nil
}

// Method for deleting a subscriber from the database.
func (SubscriberRepository) DeleteSubscriber(id uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeleteSubscriber, id); err != nil {
		return err
	}
	return nil
}
