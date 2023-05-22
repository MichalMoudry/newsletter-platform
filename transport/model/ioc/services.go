package ioc

import (
	"context"
	"newsletter-platform/database/model"
	service_model "newsletter-platform/service/model"
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

	// Method for updating user's general information in the system.
	UpdateUsersInfo(ctx context.Context, email, username, concurrencyStamp string) error
}

// Service for working with password reset tokens.
type IPassResetService interface {
	// Method for generating a new password reset token.
	GenerateNewToken(ctx context.Context, email string) error

	// Method for resetting user's password.
	ResetPassword(ctx context.Context, email, password, tokenId string) error
}

// Service for working with newsletters.
type INewsletterService interface {
	// Method for creating a new newsletter in the system.
	CreateNewsletter(ctx context.Context, name, description string) error

	// Method for obtaining a specific newsletter that is stored in the system.
	GetNewsletter(ctx context.Context, newsletterId string) (model.NewsletterData, error)

	// Method for updating a newsletter in the system.
	UpdateNewsletter(ctx context.Context, data service_model.NewsletterUpdateModel) error

	// Method for deleting a specific newsletter in the system.
	DeleteNewsletter(ctx context.Context, newsletterId string) error
}
