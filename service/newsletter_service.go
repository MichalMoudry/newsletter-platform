package service

import (
	"context"
	"fmt"
	"newsletter-platform/database"
	db_model "newsletter-platform/database/model"
	"newsletter-platform/service/model"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"
	"time"

	"github.com/google/uuid"
)

type NewsletterService struct {
	NewsletterRepo ioc.INewsletterRepository
}

func NewNewsletterService(newsletterRepo ioc.INewsletterRepository) NewsletterService {
	return NewsletterService{
		NewsletterRepo: newsletterRepo,
	}
}

// Method for creating a new newsletter in the system.
func (srvc NewsletterService) CreateNewsletter(ctx context.Context, name, description string) (newsletterId uuid.UUID, err error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() { database.EndTransaction(tx, err) }()

	newsletterId, err = srvc.NewsletterRepo.AddNewsletter(
		db_model.NewNewsletter(name, description, userId),
	)
	if err != nil {
		return uuid.Nil, err
	}
	return newsletterId, nil
}

// Method for obtaining a specific newsletter that is stored in the system.
func (srvc NewsletterService) GetNewsletter(_ context.Context, newsletterId uuid.UUID) (db_model.NewsletterData, error) {
	data, err := srvc.NewsletterRepo.GetNewsletter(newsletterId)
	if err != nil {
		return db_model.NewsletterData{}, err
	}

	return data, nil
}

// Method for updating a newsletter in the system.
func (srvc NewsletterService) UpdateNewsletter(ctx context.Context, data model.NewsletterUpdateModel) (err error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	err = srvc.NewsletterRepo.UpdateNewsletter(db_model.UpdateNewsletterData{
		Nltr_Name:           data.Name,
		Nltr_Description:    data.Description,
		OldConcurrencyStamp: data.OldConcurrencyStamp,
		NewConcurrencyStamp: uuid.New(),
		AuthorId:            userId,
		UpdateDate:          time.Now(),
		Id:                  data.NewsletterId,
	})
	if err != nil {
		return err
	}
	return nil
}

// Method for deleting a specific newsletter in the system.
func (srvc NewsletterService) DeleteNewsletter(ctx context.Context, newsletterId uuid.UUID) (err error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	if err = srvc.NewsletterRepo.DeleteNewsletter(newsletterId, userId); err != nil {
		return err
	}
	return nil
}

// Function for obtaining a user id from context.
func getUserIdFromContext(ctx context.Context) (uuid.UUID, error) {
	claims, err := util.GetClaimsFromContext(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Parse(fmt.Sprint(claims["user_id"]))
}
