package ioc

import (
	"context"

	"github.com/go-chi/jwtauth/v5"
)

// Service for working with JWT authentication and authorization.
type IAuthService interface {
	GetJwtAuth() *jwtauth.JWTAuth
}

// Service for working with users.
type IUserService interface {
	CreateUser(ctx context.Context) error
	GetUser(ctx context.Context)
	DeleteUser(ctx context.Context)
	ValidateLogin(ctx context.Context)
}

// Service for working with password reset tokens.
type IPassResetService interface {
	CreateToken(ctx context.Context)
}
