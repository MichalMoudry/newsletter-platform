package service

import (
	"context"
	"newsletter-platform/database"
	"newsletter-platform/database/model"
	"newsletter-platform/service/errors"
	email_model "newsletter-platform/service/model/email"
	"newsletter-platform/service/model/ioc"
	"newsletter-platform/service/util"
	"time"

	"github.com/google/uuid"
)

// Service for working with password reset tokens.
type PassResetService struct {
	PassResetRepo ioc.IPassResetRepository
	UserRepo      ioc.IUserRepository
	EmailService  ioc.IEmailService
}

func NewPassResetService(
	passResetRepository ioc.IPassResetRepository,
	userRepository ioc.IUserRepository,
	emailSrvc ioc.IEmailService,
) PassResetService {
	return PassResetService{
		PassResetRepo: passResetRepository,
		UserRepo:      userRepository,
		EmailService:  emailSrvc,
	}
}

// Method for generating a new password reset token.
func (srvc PassResetService) GenerateNewToken(ctx context.Context, email string) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		database.EndTransaction(tx, err)
		if err == nil {
			err = srvc.EmailService.Send(
				email,
				email_model.PassResetEmailSubject,
				email_model.PassResetEmailContent,
				email_model.PassResetEmailHtmlContent,
			)
		}
	}()

	if err = srvc.PassResetRepo.AddToken(model.NewPasswordResetToken(email)); err != nil {
		return err
	}
	return nil
}

// Method for resetting user's password.
func (srvc PassResetService) ResetPassword(ctx context.Context, email, password, tokenId string) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	err = validateEmailWithContext(ctx, email)
	if err != nil {
		return err
	}
	parsedToken, err := uuid.Parse(tokenId)
	if err != nil {
		return err
	}
	defer func() { database.EndTransaction(tx, err) }()

	// Verify token
	token, err := srvc.PassResetRepo.GetUsersLastToken(parsedToken)
	if err != nil {
		return err
	}
	if token.ExpirationDate.Before(time.Now()) {
		return errors.ErrPassResetTokenExpired
	}

	// Delete token
	err = srvc.PassResetRepo.DeleteTokens([]model.PassResetTokenData{
		token,
	})
	if err != nil {
		return err
	}

	// Reset password
	err = srvc.UserRepo.UpdatePassword(email, util.HashPassword(password, uuid.NewString()))
	if err != nil {
		return err
	}
	return nil
}
