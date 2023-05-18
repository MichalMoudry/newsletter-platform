package service

import (
	"os"

	"github.com/go-chi/jwtauth/v5"
)

// Variable pointing to a structure for working with JWT.
var tokenAuth *jwtauth.JWTAuth

type AuthenticationService struct {
	TokenAuth *jwtauth.JWTAuth
}

func CreateAutenticationService() AuthenticationService {
	return AuthenticationService{
		TokenAuth: tokenAuth,
	}
}

// Method for obtaining a pointer to JWTAuth structure.
func (authService AuthenticationService) GetJwtAuth() *jwtauth.JWTAuth {
	return authService.TokenAuth
}

func init() {
	securityString := os.Getenv("SECURITY_STRING")
	if securityString == "" {
		securityString = "132456789-/ef-/eqw-f"
	}
	tokenAuth = jwtauth.New("HS512", []byte(securityString), nil)
}
