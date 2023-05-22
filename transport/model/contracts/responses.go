package contracts

import "github.com/google/uuid"

type CreateNewsletterResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
