package query

import (
	_ "embed"
)

var (
	//go:embed scripts/commands/CreateNewsletter.sql
	CreateNewsletter string
	//go:embed scripts/commands/CreateUser.sql
	CreateUser string

	//go:embed scripts/queries/GetUser.sql
	GetUser string
	//go:embed scripts/queries/GetDataForLogin.sql
	GetDataForLogin string
)
