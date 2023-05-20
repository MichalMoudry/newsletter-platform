package ioc

import "newsletter-platform/database/model"

type IUserRepository interface {
	// Method for storing a new user in a database.
	AddUser(user *model.User) error
	// Method for obtaining user's information from a database.
	GetUser(email string) (model.UserInfo, error)
	// Method for obtaining data that is relevant to login process.
	GetDataForLogin(email string) (model.LoginData, error)
}
