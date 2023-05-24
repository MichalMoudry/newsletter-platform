package service

import (
	"context"
	"fmt"
	"newsletter-platform/database"
	db_model "newsletter-platform/database/model"
	"newsletter-platform/service/errors"
	"newsletter-platform/service/model"
	"newsletter-platform/service/model/email"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"

	"github.com/google/uuid"
)

type PostService struct {
	PostRepository         ioc.IPostRepository
	SubscriptionRepository ioc.ISubscriptionRepository
	EmailService           ioc.IEmailService
}

func NewPostService(postRepo ioc.IPostRepository, emailSrvc ioc.IEmailService, subsRepo ioc.ISubscriptionRepository) PostService {
	return PostService{
		PostRepository:         postRepo,
		SubscriptionRepository: subsRepo,
		EmailService:           emailSrvc,
	}
}

// Method for creating a new post in the system.
func (srvc PostService) CreateNewPost(ctx context.Context, data model.PostCreateModel) (postId uuid.UUID, err error) {
	authorId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	subs, err := srvc.SubscriptionRepository.GetNewsletterSubscriptions(data.NewsletterId)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	defer func() {
		err = database.EndTransaction(tx, err)
		if err == nil && len(subs) > 0 {
			if subs[0].NewsletterAuthor != authorId.String() {
				err = errors.ErrInvalidEmail
				return
			}
			var emails []string
			for _, subscription := range subs {
				emails = append(emails, subscription.Subscriber)
			}
			err = srvc.EmailService.SendBatch(
				emails,
				fmt.Sprintf("%s - %s", subs[0].NewsletterName, data.Title),
				data.Content,
				fmt.Sprintf(email.PostEmailContent, data.Content),
			)
		}
	}()

	postId, err = srvc.PostRepository.AddPost(db_model.NewPost(
		data.Title,
		data.Content,
		authorId,
		data.NewsletterId,
	))
	if err != nil {
		return uuid.Nil, err
	}
	return postId, nil
}

// Method for obtaining post information from the system.
func (srvc PostService) GetPost(_ context.Context, postId uuid.UUID) (data db_model.PostData, err error) {
	data, err = srvc.PostRepository.GetPost(postId)
	if err != nil {
		return db_model.PostData{}, err
	}

	return data, nil
}

// Function for deleting a specific post in the system.
func (srvc PostService) DeletePost(ctx context.Context, postId uuid.UUID) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = database.EndTransaction(tx, err) }()
	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	err = srvc.PostRepository.DeletePost(postId, userId)
	if err != nil {
		return err
	}
	return nil
}
