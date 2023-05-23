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

type CreateNewsletterRequestData struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateNewsletterRequestData struct {
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description" validate:"required"`
	ConcurrencyStamp string `json:"concurrency_stamp" validate:"required"`
}

type RegisterSubscriptionRequestData struct {
	Email        string `json:"email" validate:"required"`
	NewsletterId string `json:"newsletter_id" validate:"required"`
}

type CancelSubscriptionRequestData struct {
	Email string `json:"email" validate:"required"`
}
