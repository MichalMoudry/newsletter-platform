package model

import (
	"time"

	"github.com/google/uuid"
)

type Newsletter struct {
	Nltr_ID                uuid.UUID `db:"Nltr_ID"`
	Nltr_Name              string    `db:"Nltr_Name"`
	Nltr_Description       string    `db:"Nltr_Description"`
	Nltr_Inserted_Datetime time.Time `db:"Nltr_Inserted_Datetime"`
	Nltr_Updated_Datetime  time.Time `db:"Nltr_Updated_Datetime"`
	Nltr_Author            uuid.UUID `db:"Nltr_Author"`
	ConcurrencyStamp       uuid.UUID `db:"concurrency_stamp"`
}

func NewNewsletter(Nltr_Name, Nltr_Description string, Nltr_Author uuid.UUID) *Newsletter {
	now := time.Now()
	return &Newsletter{
		Nltr_ID:                uuid.New(),
		Nltr_Name:              Nltr_Name,
		Nltr_Description:       Nltr_Description,
		Nltr_Inserted_Datetime: now,
		Nltr_Updated_Datetime:  now,
		Nltr_Author:            Nltr_Author,
		ConcurrencyStamp:       uuid.New(),
	}
}

type Post struct {
	Post_ID                uuid.UUID `db:"Post_ID"`
	Post_Title             string    `db:"Post_Title"`
	Post_Content           string    `db:"Post_Content"`
	Post_Publishing_Date   time.Time `db:"Post_Publishing_Date"`
	Post_Inserted_Datetime time.Time `db:"Post_Inserted_Datetime"`
	Post_Updated_Datetime  time.Time `db:"Post_Updated_Datetime"`
	Post_Author            uuid.UUID `db:"Post_Author"`
	Nltr_ID                uuid.UUID `db:"Nltr_ID"`
	ConcurrencyStamp       uuid.UUID `db:"concurrency_stamp"`
}

func NewPost(title, content string, author, newsletterId uuid.UUID) *Post {
	now := time.Now()
	return &Post{
		Post_ID:                uuid.New(),
		Post_Title:             title,
		Post_Content:           content,
		Post_Publishing_Date:   now,
		Post_Inserted_Datetime: now,
		Post_Updated_Datetime:  now,
		Post_Author:            author,
		Nltr_ID:                newsletterId,
		ConcurrencyStamp:       uuid.New(),
	}
}

type Role string

type User struct {
	Id               uuid.UUID `db:"id"`
	Email            string    `db:"email"`
	UserName         string    `db:"user_name"`
	UserRole         Role      `db:"user_role"`
	PasswordHash     string    `db:"password_hash"`
	DateAdded        time.Time `db:"date_added"`
	DateUpdated      time.Time `db:"date_updated"`
	ConcurrencyStamp uuid.UUID `db:"concurrency_stamp"`
}

func NewUser(email, userName, password string) *User {
	now := time.Now()
	return &User{
		Id:               uuid.New(),
		Email:            email,
		UserName:         userName,
		UserRole:         "Editor",
		PasswordHash:     password,
		DateAdded:        now,
		DateUpdated:      now,
		ConcurrencyStamp: uuid.New(),
	}
}

type PasswordResetToken struct {
	Id             uuid.UUID `db:"id"`
	ExpirationDate time.Time `db:"expiration_date"`
	DateAdded      time.Time `db:"date_added"`
	Email          string    `db:"email"`
}

func NewPasswordResetToken(email string) *PasswordResetToken {
	now := time.Now()
	return &PasswordResetToken{
		Id:             uuid.New(),
		ExpirationDate: now.Add(time.Hour * 24),
		DateAdded:      now,
		Email:          email,
	}
}

type Subscriber struct {
	Id        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	DateAdded time.Time `db:"date_added"`
}

func NewSubscriber(email string) *Subscriber {
	return &Subscriber{
		Id:        uuid.New(),
		Email:     email,
		DateAdded: time.Now(),
	}
}

type NewsletterSubscription struct {
	Id           uuid.UUID `db:"id"`
	NewsletterId uuid.UUID `db:"newsletter_id"`
	SubscriberId string    `db:"subscriber_id"`
}

func NewNewsletterSubscription(newsletterId uuid.UUID, subscriberId string) *NewsletterSubscription {
	return &NewsletterSubscription{
		Id:           uuid.New(),
		NewsletterId: newsletterId,
		SubscriberId: subscriberId,
	}
}
