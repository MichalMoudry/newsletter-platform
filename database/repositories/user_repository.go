package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"

	"github.com/google/uuid"
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

// Method for obtaining data that is relevant to login process.
func (UserRepository) GetDataForLogin(email string) (model.LoginData, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.LoginData{}, err
	}

	var data model.LoginData
	if err = ctx.Get(&data, query.GetDataForLogin, email); err != nil {
		return model.LoginData{}, err
	}
	return data, nil
}

// Method for deleting a specific user in the database.
func (UserRepository) DeleteUser(email string, concurrencyStamp uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeleteUser, email, concurrencyStamp); err != nil {
		return err
	}
	return nil
}

// Method for updating user's general information in the database.
func (UserRepository) UpdateUser(data model.UserUpdateData) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.UpdateUser, data); err != nil {
		return err
	}
	return nil
}
