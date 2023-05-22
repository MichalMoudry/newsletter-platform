package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"

	"github.com/google/uuid"
)

type NewsletterRepository struct{}

// Method for adding a new newsletter to the database.
func (NewsletterRepository) AddNewsletter(data *model.Newsletter) (uuid.UUID, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return uuid.Nil, err
	}

	_, err = ctx.NamedExec(query.CreateNewsletter, data)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Nil, nil
}

// Method for obtaining data about a specific newsletter in the database.
func (NewsletterRepository) GetNewsletter(newsletterId uuid.UUID) (model.NewsletterData, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.NewsletterData{}, err
	}

	var newsletter model.NewsletterData
	if err = ctx.Select(&newsletter, query.GetNewsletter, newsletterId); err != nil {
		return model.NewsletterData{}, err
	}
	return newsletter, err
}

// Method for updating specific newsletter data in the database.
func (NewsletterRepository) UpdateNewsletter(data model.UpdateNewsletterData) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.UpdateNewsletter, data); err != nil {
		return err
	}
	return nil
}

// Method for deleting a specific newsletter in the database.
func (NewsletterRepository) DeleteNewsletter(newsletterId, authorId uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeleteNewsletter, newsletterId, authorId); err != nil {
		return err
	}
	return nil
}
