package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"

	"github.com/google/uuid"
)

type NewsletterRepository struct{}

// Method for adding a new newsletter to the database.
func (NewsletterRepository) AddNewsletter(data *model.Newsletter) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedQuery(query.CreateNewsletter, data); err != nil {
		return err
	}
	return nil
}

func GetNewsletter(newsletterId uuid.UUID) (model.Newsletter, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.Newsletter{}, err
	}

	var newsletter model.Newsletter
	return newsletter, err
}
