package ioc

import (
	"context"
	"newsletter-platform/database/model"
	"newsletter-platform/transport/model/dto"
)

// Service for working with users.
type IUserService interface {
	// Method for creating a new user in the system.
	CreateUser(ctx context.Context, data dto.NewUserData) error
	// Method for obtaining a specific user in the system.
	GetUser(ctx context.Context, email string) (model.UserInfo, error)
	// Method for validating if user provided correct credentials.
	ValidateLogin(ctx context.Context, email, password string) (string, error)
	// Method for handling of removal of a specific user in the system.
	DeleteUser(ctx context.Context, email, concurrencyStamp string) error
}

// Service for working with password reset tokens.
type IPassResetService interface {
	CreateToken(ctx context.Context)
}
