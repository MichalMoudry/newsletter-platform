package contracts

type RegisterUserRequestData struct {
	Email    string `json:"email" validate:"email,required"`
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequestData struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=8"`
}

type DeleteUserRequestData struct {
	ConcurrencyStamp string `json:"concurrency_stamp" validate:"required"`
}

type UpdateUserRequestData struct {
	UserName         string `json:"user_name" validate:"required"`
	ConcurrencyStamp string `json:"concurrency_stamp" validate:"required"`
}

type PassResetTokenRequestData struct {
	Email string `json:"email" validate:"email,required"`
}

type PassResetRequestData struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"new_password" validate:"required,min=8"`
	TokenId  string `json:"token_id" validate:"required"`
}
