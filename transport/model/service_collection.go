package model

import (
	"newsletter-platform/database/repositories"
	"newsletter-platform/service"
	"newsletter-platform/transport/model/ioc"

	"github.com/go-chi/jwtauth/v5"
)

// Structure for encapsulating all services.
type ServiceCollection struct {
	UserService         ioc.IUserService
	PassResetService    ioc.IPassResetService
	NewsletterService   ioc.INewsletterService
	SubscriptionService ioc.ISubscriptionService
	PostService         ioc.IPostService
}

func NewServiceCollection(tokenAuth *jwtauth.JWTAuth) ServiceCollection {
	userRepo := &repositories.UserRepository{}
	subscriptionRepo := &repositories.SubscriptionRepository{}
	newsletterRepo := &repositories.NewsletterRepository{}
	emailSrvc := service.NewEmailService()

	return ServiceCollection{
		UserService: service.NewUserService(
			userRepo,
			tokenAuth,
		),
		PassResetService: service.NewPassResetService(
			repositories.PassResetRepository{},
			userRepo,
			emailSrvc,
		),
		NewsletterService: service.NewNewsletterService(newsletterRepo),
		SubscriptionService: service.NewSubscriptionService(
			subscriptionRepo,
			repositories.SubscriberRepository{},
			newsletterRepo,
			emailSrvc,
		),
		PostService: service.NewPostService(
			repositories.PostRepository{},
			emailSrvc,
			subscriptionRepo,
		),
	}
}
