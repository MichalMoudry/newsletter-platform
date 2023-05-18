package contracts

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"email,required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
