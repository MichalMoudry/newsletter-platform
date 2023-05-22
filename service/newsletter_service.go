package service

import (
	"context"
	"fmt"
	"newsletter-platform/database"
	db_model "newsletter-platform/database/model"
	"newsletter-platform/service/model"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"

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
func (srvc NewsletterService) CreateNewsletter(ctx context.Context, name, description string) (err error) {
	claims, err := util.GetClaimsFromContext(ctx)
	if err != nil {
		return err
	}
	author, err := uuid.Parse(fmt.Sprint(claims["user_id"]))

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	err = srvc.NewsletterRepo.AddNewsletter(
		db_model.NewNewsletter(name, description, author),
	)
	if err != nil {
		return err
	}
	return nil
}

// Method for obtaining a specific newsletter that is stored in the system.
func (srvc NewsletterService) GetNewsletter(_ context.Context, newsletterId string) (db_model.NewsletterData, error) {
	id, err := uuid.Parse(newsletterId)
	if err != nil {
		return db_model.NewsletterData{}, err
	}

	data, err := srvc.NewsletterRepo.GetNewsletter(id)
	if err != nil {
		return db_model.NewsletterData{}, err
	}

	return data, nil
}

// Method for updating a newsletter in the system.
func (srvc NewsletterService) UpdateNewsletter(ctx context.Context, data model.NewsletterUpdateModel) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	err = srvc.NewsletterRepo.UpdateNewsletter(db_model.UpdateNewsletterData{
		Nltr_Name:           data.Name,
		Nltr_Description:    data.Description,
		OldConcurrencyStamp: data.OldConcurrencyStamp,
		NewConcurrencyStamp: data.NewConcurrencyStamp,
	})
	if err != nil {
		return err
	}
	return nil
}

// Method for deleting a specific newsletter in the system.
func (srvc NewsletterService) DeleteNewsletter(ctx context.Context, newsletterId string) (err error) {
	id, err := uuid.Parse(newsletterId)
	if err != nil {
		return err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	if err = srvc.NewsletterRepo.DeleteNewsletter(id); err != nil {
		return err
	}
	return nil
}
