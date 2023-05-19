package model

import "github.com/google/uuid"

// Structure encapsulating information available to the user.
type UserInfo struct {
	Email            string    `db:"email"`
	UserName         string    `db:"user_name"`
	UserRole         string    `db:"user_role"`
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
}
