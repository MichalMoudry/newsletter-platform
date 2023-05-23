package query

import (
	_ "embed"
)

var (
	//go:embed scripts/commands/CreateNewsletter.sql
	CreateNewsletter string
	//go:embed scripts/commands/CreatePost.sql
	CreatePost string
	//go:embed scripts/commands/CreateUser.sql
	CreateUser string

	//go:embed scripts/commands/DeleteUser.sql
	DeleteUser string
	//go:embed scripts/commands/DeleteNewsletter.sql
	DeleteNewsletter string
	//go:embed scripts/commands/DeletePost.sql
	DeletePost string

	//go:embed scripts/commands/UpdateUser.sql
	UpdateUser string
	//go:embed scripts/commands/UpdateNewsletter.sql
	UpdateNewsletter string
	//go:embed scripts/commands/UpdatePost.sql
	UpdatePost string

	//go:embed scripts/commands/AddPassResetToken.sql
	AddPassResetToken string
	//go:embed scripts/commands/DeletePassResetToken.sql
	DeletePassResetToken string
	//go:embed scripts/commands/ResetPassword.sql
	ResetPassword string
	//go:embed scripts/commands/AddSubscription.sql
	AddSubscription string
	//go:embed scripts/commands/DeleteSubscription.sql
	DeleteSubscription string
	//go:embed scripts/commands/AddSubscriber.sql
	AddSubscriber string
	//go:embed scripts/commands/DeleteSubscriber.sql
	DeleteSubscriber string

	//go:embed scripts/queries/GetUser.sql
	GetUser string
	//go:embed scripts/queries/GetDataForLogin.sql
	GetDataForLogin string
	//go:embed scripts/queries/GetPassResetToken.sql
	GetPassResetToken string
	//go:embed scripts/queries/GetNewsletter.sql
	GetNewsletter string
	//go:embed scripts/queries/GetSubscriber.sql
	GetSubscriber string
	//go:embed scripts/queries/GetPost.sql
	GetPost string
)
