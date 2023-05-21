package repositories

import (
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/database/query"

	"github.com/google/uuid"
)

type PassResetRepository struct{}

// Method for inserting a new token for password reset.
func (PassResetRepository) AddToken(token *model.PasswordResetToken) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.NamedExec(query.AddPassResetToken, token); err != nil {
		return err
	}
	return nil
}

// Method for obtaining last/newest user's token.
func (PassResetRepository) GetPassResetToken(tokenId uuid.UUID) (model.PassResetTokenData, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.PassResetTokenData{}, err
	}

	var token model.PassResetTokenData
	if err = ctx.Get(&token, query.GetPassResetToken, tokenId); err != nil {
		return model.PassResetTokenData{}, err
	}

	return token, nil
}

// Method for deleting one or more password reset tokens in the database.
func (PassResetRepository) DeleteTokens(tokens []model.PassResetTokenData) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	var params []map[string]any
	for _, v := range tokens {
		params = append(params, map[string]any{
			"token_id": v.Id,
		})
	}
	if _, err = ctx.NamedExec(query.DeletePassResetToken, params); err != nil {
		return err
	}

	return nil
}
