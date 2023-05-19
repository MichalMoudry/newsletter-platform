package ioc

import (
	"context"
	"newsletter-platform/transport/model/dto"
)

// Service for working with users.
type IUserService interface {
	// Method for creating a new user in the system.
	CreateUser(ctx context.Context, data dto.NewUserData) error
	GetUser(ctx context.Context)
	/*DeleteUser(ctx context.Context)
	ValidateLogin(ctx context.Context)*/
}

// Service for working with password reset tokens.
type IPassResetService interface {
	CreateToken(ctx context.Context)
}
