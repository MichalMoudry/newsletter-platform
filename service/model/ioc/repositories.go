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
}

type IPassResetRepository interface {
	// Method for inserting a new token for password reset.
	AddToken(token *model.PasswordResetToken) error

	// Method for obtaining last/newest user's token.
	GetUsersLastToken(email string) (model.PasswordResetToken, error)

	// Method for deleting one or more password reset tokens in the database.
	DeleteTokens(tokens []model.PasswordResetToken) error
}
