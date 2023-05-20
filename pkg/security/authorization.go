package security

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthorizationFilter interface {
	Authorize(ctx *gin.Context)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}

type AuthorizationDelegate interface {
	Authorize(ctx context.Context, principal *Principal) error
}
