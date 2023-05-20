package security

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
)

type DefaultClaims struct {
	jwt.RegisteredClaims
	Principal
}

type DefaultTokenTokenManager struct {
	issuer        string
	timeout       time.Duration
	secretKey     any
	signingMethod jwt.SigningMethod
}

func NewDefaultTokenTokenManager(issuer string, timeout time.Duration, secretKey any, signingMethod jwt.SigningMethod) *DefaultTokenTokenManager {
	return &DefaultTokenTokenManager{
		issuer:        issuer,
		timeout:       timeout,
		secretKey:     secretKey,
		signingMethod: signingMethod,
	}
}

func (manager *DefaultTokenTokenManager) Generate(principal *Principal) (*string, error) {

	principal.Password = nil

	claims := &DefaultClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    manager.issuer,
			Subject:   *principal.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.timeout)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Principal: *principal,
	}

	token := jwt.NewWithClaims(manager.signingMethod, claims)

	var err error
	var tokenString string
	if tokenString, err = token.SignedString(manager.secretKey); err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (manager *DefaultTokenTokenManager) Validate(tokenString string) (*Principal, error) {

	getKeyFunc := func(token *jwt.Token) (any, error) {
		return manager.secretKey, nil
	}

	parserOptions := []jwt.ParserOption{
		jwt.WithIssuer(manager.issuer),
		jwt.WithValidMethods([]string{manager.signingMethod.Alg()}),
	}

	var err error
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, getKeyFunc, parserOptions...); err != nil {
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	var ok bool
	var mapClaims jwt.MapClaims
	if mapClaims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("invalid token")
	}

	var value any
	if value, ok = mapClaims["username"]; !ok {
		return nil, errors.New("invalid token")
	}

	var username string
	if username, ok = value.(string); !ok {
		return nil, errors.New("invalid token")
	}

	principal := &Principal{
		Username: feather_commons_util.ValueToPtr(username),
	}

	return principal, err
}
