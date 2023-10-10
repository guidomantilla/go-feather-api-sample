package rest

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
)

type DefaultAuthPrincipalEndpoint struct {
	principalManager feather_security.PrincipalManager
}

func NewDefaultAuthPrincipalEndpoint(principalManager feather_security.PrincipalManager) *DefaultAuthPrincipalEndpoint {
	return &DefaultAuthPrincipalEndpoint{
		principalManager: principalManager,
	}
}

func (endpoint *DefaultAuthPrincipalEndpoint) Create(ctx *gin.Context) {

}

func (endpoint *DefaultAuthPrincipalEndpoint) Update(ctx *gin.Context) {

}

func (endpoint *DefaultAuthPrincipalEndpoint) Delete(ctx *gin.Context) {

}

func (endpoint *DefaultAuthPrincipalEndpoint) FindByUsername(ctx *gin.Context) {

	var err error
	var body []byte
	if body, err = io.ReadAll(ctx.Request.Body); err != nil {
		ex := feather_web_rest.BadRequestException("error reading body")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if len(body) != 0 {
		ex := feather_web_rest.BadRequestException("body is not empty")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var user *feather_security.Principal
	username := ctx.Param("username")
	if user, err = endpoint.principalManager.Find(ctx.Request.Context(), username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	user.Password, user.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, user)
}

func (endpoint *DefaultAuthPrincipalEndpoint) FindCurrent(ctx *gin.Context) {

	var err error
	var body []byte
	if body, err = io.ReadAll(ctx.Request.Body); err != nil {
		ex := feather_web_rest.BadRequestException("error reading body")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if len(body) != 0 {
		ex := feather_web_rest.BadRequestException("body is not empty")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var exists bool
	var principal *feather_security.Principal
	if principal, exists = feather_security.GetPrincipalFromContext(ctx); !exists {
		ex := feather_web_rest.NotFoundException("principal not found in context")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var user *feather_security.Principal
	if user, err = endpoint.principalManager.Find(ctx.Request.Context(), *principal.Username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	user.Password, user.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, user)
}

func (endpoint *DefaultAuthPrincipalEndpoint) ChangePassword(ctx *gin.Context) {

}

func (endpoint *DefaultAuthPrincipalEndpoint) VerifyResource(ctx *gin.Context) {

}
