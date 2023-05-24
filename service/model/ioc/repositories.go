package ioc

import (
	"newsletter-platform/database/model"

	"github.com/google/uuid"
)

type IUserRepository interface {
	// Method for storing a new user in a database.
	AddUser(user *model.User) error

	// Method for obtaining user's information from a database.
	GetUser(email string) (model.UserInfo, error)

	// Method for obtaining data that is relevant to login process.
	GetDataForLogin(email string) (model.LoginData, error)

	// Method for deleting a specific user in the database.
	DeleteUser(email string, concurrencyStamp uuid.UUID) error

	// Method for updating user's general information in the database.
	UpdateUser(data model.UserUpdateData) error

	// Method for updating user's password.
	UpdatePassword(email, passwordHash string) error
}

type IPassResetRepository interface {
	// Method for inserting a new token for password reset.
	AddToken(token *model.PasswordResetToken) error

	// Method for obtaining last/newest user's token.
	GetPassResetToken(tokenId uuid.UUID) (model.PassResetTokenData, error)

	// Method for deleting one or more password reset tokens in the database.
	DeleteTokens(tokens []model.PassResetTokenData) error
}

type INewsletterRepository interface {
	// Method for adding a new newsletter to the database.
	AddNewsletter(data *model.Newsletter) (uuid.UUID, error)

	// Method for obtaining data about a specific newsletter in the database.
	GetNewsletter(newsletterId uuid.UUID) (model.NewsletterData, error)

	// Method for updating specific newsletter data in the database.
	UpdateNewsletter(data model.UpdateNewsletterData) error

	// Method for deleting a specific newsletter in the database.
	DeleteNewsletter(newsletterId uuid.UUID, authorId uuid.UUID) error
}

type ISubscriberRepository interface {
	// Method for adding a new subscriber to the database.
	AddSubscriber(sub *model.Subscriber) error

	// Method for obtaining information about a subscriber from the database.
	GetSubscriber(email string) (model.SubscriberInformation, error)

	// Method for deleting a subscriber from the database.
	DeleteSubscriber(id uuid.UUID) error
}

type ISubscriptionRepository interface {
	// Method for adding a new subscription to the database.
	AddSubscription(sub *model.NewsletterSubscription) error

	// Method for deleting a specific subscription in the database.
	DeleteSubscription(email string, newsletterId uuid.UUID) error

	// Method for obtaining list of newsletter subscribers.
	GetNewsletterSubscriptions(newsletterId uuid.UUID) (model.NewsletterSubscriptions, error)
}

// Function for deleting a specific post in the system.
type IPostRepository interface {
	// Method for adding a new post to the database.
	AddPost(data *model.Post) (uuid.UUID, error)

	// Method for obtaining data about a specific post in the database.
	GetPost(postId uuid.UUID) (model.PostData, error)

	// Method for updating specific Post data in the database.
	UpdatePost(data model.UpdatePostData) error

	// Method for deleting a specific Post in the database.
	DeletePost(postId, authorId uuid.UUID) error
}
