package config

import (
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
)

func InitAuthEndpoints(_ environment.Environment, tokenTokenManager feather_security.TokenManager,
	authenticationDelegate feather_security.AuthenticationDelegate, authorizationDelegate feather_security.AuthorizationDelegate) (feather_security.AuthenticationEndpoint, feather_security.AuthorizationFilter) {

	authenticationService := feather_security.NewDefaultAuthenticationService(tokenTokenManager, authenticationDelegate)
	authenticationEndpoint := feather_security.NewDefaultAuthenticationEndpoint(authenticationService)

	authorizationService := feather_security.NewDefaultAuthorizationService(tokenTokenManager, authorizationDelegate)
	authorizationFilter := feather_security.NewDefaultAuthorizationFilter(authorizationService)

	return authenticationEndpoint, authorizationFilter
}
