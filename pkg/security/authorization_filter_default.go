package security

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type DefaultAuthorizationFilter struct {
	authorizationService AuthorizationService
}

func NewDefaultAuthorizationFilter(authorizationService AuthorizationService) *DefaultAuthorizationFilter {
	return &DefaultAuthorizationFilter{
		authorizationService: authorizationService,
	}
}

func (filter *DefaultAuthorizationFilter) Authorize(ctx *gin.Context) {

	header := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		ex := UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	splits := strings.Split(header, " ")
	if len(splits) != 2 {
		ex := UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var err error
	var principal *Principal
	if principal, err = filter.authorizationService.Authorize(ctx.Request.Context(), splits[1]); err != nil {
		ex := UnauthorizedException("invalid authorization header", err)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.Set("principal", principal)
	ctx.Next()
}
