package rest

import "github.com/gin-gonic/gin"

var (
	_ AuthResourceEndpoint          = (*DefaultAuthResourceEndpoint)(nil)
	_ AuthRoleEndpoint              = (*DefaultAuthRoleEndpoint)(nil)
	_ AuthAccessControlListEndpoint = (*DefaultAuthAccessControlListEndpoint)(nil)
	_ AuthUserEndpoint              = (*DefaultAuthUserEndpoint)(nil)
	_ AuthPrincipalEndpoint         = (*DefaultAuthPrincipalEndpoint)(nil)
)

type AuthResourceEndpoint interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type AuthRoleEndpoint interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type AuthAccessControlListEndpoint interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type AuthUserEndpoint interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type AuthPrincipalEndpoint interface {
	GetCurrentPrincipal(ctx *gin.Context)
}
