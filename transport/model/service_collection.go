package model

import (
	"newsletter-platform/database/repositories"
	"newsletter-platform/service"
	"newsletter-platform/transport/model/ioc"

	"github.com/go-chi/jwtauth/v5"
)

// Structure for encapsulating all services.
type ServiceCollection struct {
	UserService       ioc.IUserService
	PassResetService  ioc.IPassResetService
	NewsletterService ioc.INewsletterService
}

func NewServiceCollection(tokenAuth *jwtauth.JWTAuth) ServiceCollection {
	userRepo := &repositories.UserRepository{}

	return ServiceCollection{
		UserService: service.NewUserService(
			userRepo,
			tokenAuth,
		),
		PassResetService: service.NewPassResetService(
			repositories.PassResetRepository{},
			userRepo,
			service.NewEmailService(),
		),
		NewsletterService: service.NewNewsletterService(
			repositories.NewsletterRepository{},
		),
	}
}
