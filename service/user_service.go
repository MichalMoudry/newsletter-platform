package service

import (
	"context"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service/errors"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"
	"newsletter-platform/transport/model/dto"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

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

// Method for obtaining a specific user in the system.
func (srvc UserService) GetUser(ctx context.Context, email string) (model.UserInfo, error) {
	jwtClaims, err := util.GetClaimsFromContext(ctx)
	if err != nil {
		return model.UserInfo{}, errors.ErrJwtParsing
	}
	if jwtClaims["sub"] != email {
		return model.UserInfo{}, errors.ErrInvalidEmail
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
		return "", nil
	}

	match, err := util.ComparePasswordHash(password, data.PasswordHash)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.ErrInvalidLogin
	}

	_, token, err := srvc.tokenAuth.Encode(util.GetJwtTokenContent(email))
	if err != nil {
		return "", err
	}

	return token, nil
}
