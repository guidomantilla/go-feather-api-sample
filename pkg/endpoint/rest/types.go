package rest

import "github.com/gin-gonic/gin"

var (
	_ AuthPrincipalEndpoint = (*DefaultAuthPrincipalEndpoint)(nil)
)

type AuthPrincipalEndpoint interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindByUsername(ctx *gin.Context)
	FindCurrent(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}
