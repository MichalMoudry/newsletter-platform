package model

import (
	"newsletter-platform/database/repositories"
	"newsletter-platform/service"
	"newsletter-platform/transport/model/ioc"
)

// Structure for encapsulating all services.
type ServiceCollection struct {
	UserService      ioc.IUserService
	PassResetService ioc.IPassResetService
}

func NewServiceCollection() ServiceCollection {
	return ServiceCollection{
		UserService: service.NewUserService(
			repositories.UserRepository{},
		),
	}
}
