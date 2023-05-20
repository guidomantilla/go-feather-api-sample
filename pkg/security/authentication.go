package security

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthenticationEndpoint interface {
	Login(ctx *gin.Context)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) (*string, error)
}

type AuthenticationDelegate interface {
	Authenticate(ctx context.Context, principal *Principal) error
}
