package service

import (
	"context"
	"fmt"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service/errors"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"
	"newsletter-platform/transport/model/dto"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

// Service that provides functionality for dealing with users.
type UserService struct {
	UserRepo  ioc.IUserRepository
	tokenAuth *jwtauth.JWTAuth
}

// Function for obtaining an instance of Editor service struct.
func NewUserService(storage ioc.IUserRepository, auth *jwtauth.JWTAuth) UserService {
	return UserService{
		UserRepo:  storage,
		tokenAuth: auth,
	}
}

// Method for creating a new user in the system.
func (srvc UserService) CreateUser(ctx context.Context, data dto.NewUserData) (err error) {
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

// Method for obtaining a specific user in the system.
func (srvc UserService) GetUser(ctx context.Context, email string) (model.UserInfo, error) {
	err := validateEmailWithContext(ctx, email)
	if err != nil {
		return model.UserInfo{}, err
	}

	data, err := srvc.UserRepo.GetUser(email)
	if err != nil {
		return model.UserInfo{}, err
	}
	return data, nil
}

// Method for validating credentials provided by a user.
func (srvc UserService) ValidateLogin(_ context.Context, email, password string) (string, error) {
	data, err := srvc.UserRepo.GetDataForLogin(email)
	if err != nil {
		return "", err
	}

	match, err := util.ComparePasswordHash(password, data.PasswordHash)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.ErrInvalidLogin
	}

	_, token, err := srvc.tokenAuth.Encode(util.GetJwtTokenContent(email, data.UserId))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Method for handling of removal of a specific user in the system.
func (srvc UserService) DeleteUser(ctx context.Context, email, concurrencyStamp string) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	err = validateEmailWithContext(ctx, email)
	if err != nil {
		return err
	}

	stamp, err := uuid.Parse(concurrencyStamp)
	if err != nil {
		return err
	}

	err = srvc.UserRepo.DeleteUser(email, stamp)
	if err != nil {
		return err
	}
	return nil
}

// Method for updating user's general information in the system.
func (srvc UserService) UpdateUsersInfo(ctx context.Context, email, username, concurrencyStamp string) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	err = validateEmailWithContext(ctx, email)
	if err != nil {
		return err
	}

	claims, err := util.GetClaimsFromContext(ctx)
	if err != nil {
		return err
	}
	userId, err := uuid.Parse(fmt.Sprint(claims["user_id"]))
	if err != nil {
		return err
	}
	oldConcurrencyStamp, err := uuid.Parse(concurrencyStamp)
	if err != nil {
		return err
	}

	err = srvc.UserRepo.UpdateUser(model.UserUpdateData{
		UserId:              userId,
		Email:               email,
		UserName:            username,
		UpdateDate:          time.Now(),
		OldConcurrencyStamp: oldConcurrencyStamp,
		NewConcurrencyStamp: uuid.New(),
	})
	if err != nil {
		return err
	}
	return nil
}

// Function for validating if email in the context equals a specific email.
func validateEmailWithContext(ctx context.Context, email string) error {
	jwtClaims, err := util.GetClaimsFromContext(ctx)
	if err != nil {
		return errors.ErrJwtParsing
	}
	if jwtClaims["sub"] != email {
		return errors.ErrInvalidEmail
	}
	return nil
}
