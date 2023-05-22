package ioc

import (
	"newsletter-platform/database/model"

	"github.com/google/uuid"
)

type IUserRepository interface {
	// Method for storing a new user in a database.
	AddUser(user *model.User) error

	// Method for obtaining user's information from a database.
	GetUser(email string) (model.UserInfo, error)

	// Method for obtaining data that is relevant to login process.
	GetDataForLogin(email string) (model.LoginData, error)

	// Method for deleting a specific user in the database.
	DeleteUser(email string, concurrencyStamp uuid.UUID) error

	// Method for updating user's general information in the database.
	UpdateUser(data model.UserUpdateData) error

	// Method for updating user's password.
	UpdatePassword(email, passwordHash string) error
}

type IPassResetRepository interface {
	// Method for inserting a new token for password reset.
	AddToken(token *model.PasswordResetToken) error

	// Method for obtaining last/newest user's token.
	GetPassResetToken(tokenId uuid.UUID) (model.PassResetTokenData, error)

	// Method for deleting one or more password reset tokens in the database.
	DeleteTokens(tokens []model.PassResetTokenData) error
}

type INewsletterRepository interface {
	// Method for adding a new newsletter to the database.
	AddNewsletter(data *model.Newsletter) error

	// Method for obtaining data about a specific newsletter in the database.
	GetNewsletter(newsletterId uuid.UUID) (model.NewsletterData, error)

	// Method for updating specific newsletter data in the database.
	UpdateNewsletter(data model.UpdateNewsletterData) error

	// Method for deleting a specific newsletter in the database.
	DeleteNewsletter(newsletterId uuid.UUID) error
}
