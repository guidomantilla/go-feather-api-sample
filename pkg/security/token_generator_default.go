package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type DefaultClaims struct {
	jwt.RegisteredClaims
	Principal
}

type DefaultTokenGenerator struct {
	issuer        string
	timeout       time.Duration
	secretKey     any
	signingMethod jwt.SigningMethod
}

func NewDefaultTokenGenerator(issuer string, timeout time.Duration, secretKey any, signingMethod jwt.SigningMethod) *DefaultTokenGenerator {
	return &DefaultTokenGenerator{
		issuer:        issuer,
		timeout:       timeout,
		secretKey:     secretKey,
		signingMethod: signingMethod,
	}
}

func (generator *DefaultTokenGenerator) Generate(principal *Principal) (*string, error) {

	principal.Password = nil

	claims := &DefaultClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    generator.issuer,
			Subject:   *principal.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(generator.timeout)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Principal: *principal,
	}
	token := jwt.NewWithClaims(generator.signingMethod, claims)

	var err error
	var tokenString string
	if tokenString, err = token.SignedString(generator.secretKey); err != nil {
		return nil, err
	}

	return &tokenString, nil
}
