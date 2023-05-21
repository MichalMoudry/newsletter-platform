package dto

import "newsletter-platform/transport/model/contracts"

type NewUserData struct {
	Email    string
	UserName string
	Password string
}

// Function for transforming registration request to a DTO.
func ToNewUserData(r contracts.RegisterUserRequestData) NewUserData {
	return NewUserData{
		Email:    r.Email,
		UserName: r.UserName,
		Password: r.Password,
	}
}
