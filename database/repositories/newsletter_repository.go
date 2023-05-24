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

	rows, err := ctx.NamedQuery(query.CreateNewsletter, data)
	if err != nil {
		return uuid.Nil, err
	}

	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
	}
	if err != nil {
		return uuid.Nil, err
	}

	newsletterId, err := uuid.Parse(returnedId)
	if err != nil {
		return uuid.Nil, err
	}

	return newsletterId, nil
}

// Method for obtaining data about a specific newsletter in the database.
func (NewsletterRepository) GetNewsletter(newsletterId uuid.UUID) (model.NewsletterData, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.NewsletterData{}, err
	}

	var newsletter model.NewsletterData
	if err = ctx.Get(&newsletter, query.GetNewsletter, newsletterId); err != nil {
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

// Method for obtaining newsletter's posts from the database.
func (NewsletterRepository) GetPosts(newsletterId uuid.UUID) ([]model.PostData, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return make([]model.PostData, 0), err
	}

	var data []model.PostData
	if err = ctx.Select(&data, query.GetNewsletterPosts, newsletterId); err != nil {
		return make([]model.PostData, 0), err
	}

	return data, nil
}
