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
