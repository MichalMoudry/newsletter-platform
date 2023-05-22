package service

import (
	"context"
	"newsletter-platform/database"
	"newsletter-platform/service/model/ioc"
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
func (srvc NewsletterService) CreateNewsletter(ctx context.Context) error {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	return nil
}
