package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"
)

type UserRepository struct{}

// Method for storing a new user in a database.
func (UserRepository) AddUser(user *model.User) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.CreateUser, user); err != nil {
		return err
	}
	return nil
}

// Method for obtaining user's information from a database.
func (UserRepository) GetUser(email string) (model.UserInfo, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.UserInfo{}, err
	}

	var userInfo model.UserInfo
	if err = ctx.Get(&userInfo, query.GetUser, email); err != nil {
		return model.UserInfo{}, err
	}
	return userInfo, nil
}

func (UserRepository) DeleteUser() error {
	return nil
}

func (UserRepository) UpdateUser() error {
	return nil
}
