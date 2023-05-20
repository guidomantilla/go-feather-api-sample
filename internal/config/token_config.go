package config

import (
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
)

func InitToken(environment environment.Environment) feather_security.TokenManager {

	secret := environment.GetValueOrDefault(TokenSignatureKey, EnvVarDefaultValuesMap[TokenSignatureKey]).AsString()
	tokenTokenManager := feather_security.NewJwtTokenManager([]byte(secret), feather_security.WithIssuer(AppName))
	return tokenTokenManager
}
