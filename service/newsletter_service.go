package service

import (
	"context"
	"fmt"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
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
		model.NewNewsletter(name, description, author),
	)
	if err != nil {
		return err
	}
	return nil
}

// Method for obtaining a specific newsletter that is stored in the system.
func (srvc NewsletterService) GetNewsletter(_ context.Context, newsletterId string) (model.NewsletterData, error) {
	id, err := uuid.Parse(newsletterId)
	if err != nil {
		return model.NewsletterData{}, err
	}

	data, err := srvc.NewsletterRepo.GetNewsletter(id)
	if err != nil {
		return model.NewsletterData{}, err
	}

	return data, nil
}

// Method for updating a newsletter in the system.
func UpdateNewsletter(ctx context.Context) (err error) {
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
