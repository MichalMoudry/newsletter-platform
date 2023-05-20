package contracts

type RegisterUserRequest struct {
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
