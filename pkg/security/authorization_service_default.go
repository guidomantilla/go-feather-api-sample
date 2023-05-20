package security

import (
	"context"
)

type DefaultAuthorizationService struct {
	tokenManager          TokenManager
	authorizationDelegate AuthorizationDelegate
}

func NewDefaultAuthorizationService(tokenManager TokenManager, authorizationDelegate AuthorizationDelegate) *DefaultAuthorizationService {
	return &DefaultAuthorizationService{
		tokenManager:          tokenManager,
		authorizationDelegate: authorizationDelegate,
	}
}

func (service *DefaultAuthorizationService) Authorize(ctx context.Context, tokenString string) (*Principal, error) {

	var err error
	var principal *Principal
	if principal, err = service.tokenManager.Validate(tokenString); err != nil {
		return nil, err
	}

	if err = service.authorizationDelegate.Authorize(ctx, principal); err != nil {
		return nil, err
	}

	return principal, nil
}
