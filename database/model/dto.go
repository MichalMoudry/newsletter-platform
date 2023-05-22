package model

import (
	"time"

	"github.com/google/uuid"
)

// Structure encapsulating information available to the user.
type UserInfo struct {
	Email            string    `db:"email"`
	UserName         string    `db:"user_name"`
	UserRole         string    `db:"user_role"`
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
}

// Structure encapsulating data used in login process.
type LoginData struct {
	UserId       uuid.UUID `db:"id"`
	PasswordHash string    `db:"password_hash"`
}

type UserUpdateData struct {
	UserId              uuid.UUID `db:"user_id"`
	Email               string    `db:"email"`
	UserName            string    `db:"user_name"`
	OldConcurrencyStamp uuid.UUID `db:"old_concurrency_stamp"`
	NewConcurrencyStamp uuid.UUID `db:"new_concurrency_stamp"`
}

type PassResetTokenData struct {
	Id             uuid.UUID `db:"id"`
	ExpirationDate time.Time `db:"expiration_date"`
	DateAdded      time.Time `db:"date_added"`
}

type NewsletterData struct {
	Nltr_Name        string `db:"Nltr_Name"`
	Nltr_Description string `db:"Nltr_Description"`
}
