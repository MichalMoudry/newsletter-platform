package service

import "newsletter-platform/service/model/ioc"

// Service for working with password reset tokens.
type PassResetService struct {
	PassResetRepo ioc.IPassResetRepository
}

func NewPassResetService(repository ioc.IPassResetRepository) PassResetService {
	return PassResetService{
		PassResetRepo: repository,
	}
}
