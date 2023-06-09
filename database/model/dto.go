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

// Structure encapsulating data that is need for updating user's information.
type UserUpdateData struct {
	UserId              uuid.UUID `db:"user_id"`
	Email               string    `db:"email"`
	UserName            string    `db:"user_name"`
	UpdateDate          time.Time `db:"date_updated"`
	OldConcurrencyStamp uuid.UUID `db:"old_concurrency_stamp"`
	NewConcurrencyStamp uuid.UUID `db:"new_concurrency_stamp"`
}

type PassResetTokenData struct {
	Id             uuid.UUID `db:"id"`
	ExpirationDate time.Time `db:"expiration_date"`
	DateAdded      time.Time `db:"date_added"`
}

type NewsletterData struct {
	Nltr_ID          uuid.UUID `db:"nltr_id"`
	Nltr_Name        string    `db:"nltr_name"`
	Nltr_Description string    `db:"nltr_description"`
	UserName         string    `db:"user_name"`
}

type UpdateNewsletterData struct {
	Id                  uuid.UUID `db:"Nltr_ID"`
	Nltr_Name           string    `db:"Nltr_Name"`
	Nltr_Description    string    `db:"Nltr_Description"`
	AuthorId            uuid.UUID `db:"Nltr_Author"`
	OldConcurrencyStamp uuid.UUID `db:"old_concurrency_stamp"`
	NewConcurrencyStamp uuid.UUID `db:"new_concurrency_stamp"`
	UpdateDate          time.Time `db:"Nltr_Updated_Datetime"`
}

type SubscriberInformation struct {
	Id    uuid.UUID `db:"id"`
	Email string    `db:"email"`
}

type PostData struct {
	Title   string `db:"post_title"`
	Content string `db:"post_content"`
	Author  string `db:"user_name"`
}

type UpdatePostData struct {
	Id                   uuid.UUID `db:"Post_ID"`
	Post_Title           string    `db:"Post_Title"`
	Post_Content         string    `db:"Post_Content"`
	Post_Publishing_Date time.Time `db:"Post_Inserted_Datetime"`
	Inserted_Datetime    time.Time `db:"Post_Inserted_Datetime"`
	Updated_Datetime     time.Time `db:"Post_Updated_Datetime"`
	authorId             uuid.UUID `db:"Post_Author"`
	Nltr_ID              uuid.UUID `db:"Nltr_ID"`
	concurrency_stamp    uuid.UUID `db:"concurrency_stamp"`
}

type NewsletterInformation struct {
	NewsletterName   string `db:"nltr_name"`
	NewsletterAuthor string `db:"nltr_author"`
	Subscriber       string `db:"subscriber_id"`
}
