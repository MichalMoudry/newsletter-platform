package query

import (
	_ "embed"
)

var (
	//go:embed scripts/commands/CreateNewsletter.sql
	CreateNewsletter string
)
