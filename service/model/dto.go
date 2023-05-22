package model

import "github.com/google/uuid"

type NewsletterUpdateModel struct {
	Name                string
	Description         string
	OldConcurrencyStamp uuid.UUID
	NewConcurrencyStamp uuid.UUID
}
