package security

import (
	"context"
)

type DefaultAuthenticationService struct {
	tokenGenerator         TokenGenerator
	authenticationDelegate AuthenticationDelegate
}

func NewDefaultAuthenticationService(tokenGenerator TokenGenerator, authenticationDelegate AuthenticationDelegate) *DefaultAuthenticationService {
	return &DefaultAuthenticationService{
		tokenGenerator:         tokenGenerator,
		authenticationDelegate: authenticationDelegate,
	}
}

func (service *DefaultAuthenticationService) Authenticate(ctx context.Context, principal *Principal) (*string, error) {

	var err error
	if err = service.authenticationDelegate.Authenticate(ctx, principal); err != nil {
		return nil, err
	}

	var token *string
	if token, err = service.tokenGenerator.Generate(principal); err != nil {
		return nil, err
	}

	return token, nil
}
