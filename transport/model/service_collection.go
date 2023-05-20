package model

import (
	"newsletter-platform/database/repositories"
	"newsletter-platform/service"
	"newsletter-platform/transport/model/ioc"

	"github.com/go-chi/jwtauth/v5"
)

// Structure for encapsulating all services.
type ServiceCollection struct {
	UserService      ioc.IUserService
	PassResetService ioc.IPassResetService
}

func NewServiceCollection(tokenAuth *jwtauth.JWTAuth) ServiceCollection {
	return ServiceCollection{
		UserService: service.NewUserService(
			repositories.UserRepository{},
			tokenAuth,
		),
	}
}
