package ioc

import "newsletter-platform/database/model"

type IUserRepository interface {
	// Method for storing a new user in a database.
	AddUser(user *model.User) error
}
