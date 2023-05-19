package service

import (
	"context"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"
	"newsletter-platform/transport/model/dto"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepo ioc.IUserRepository
}

// Function for obtaining an instance of Editor service struct.
func NewUserService(storage ioc.IUserRepository) UserService {
	return UserService{
		UserRepo: storage,
	}
}

// Method for creating a new user in the system.
func (srvc UserService) CreateUser(ctx context.Context, data dto.NewUserData) error {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	err = srvc.UserRepo.AddUser(
		model.NewUser(
			data.Email,
			data.UserName,
			util.HashPassword(data.Password, uuid.NewString()),
		),
	)
	if err != nil {
		return err
	}
	return nil
}
