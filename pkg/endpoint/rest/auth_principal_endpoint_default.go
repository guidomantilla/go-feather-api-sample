package rest

import (
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

func (endpoint *DefaultAuthPrincipalEndpoint) GetCurrentPrincipal(ctx *gin.Context) {

	var principal any
	var exists bool
	if principal, exists = ctx.Get("principal"); !exists {
		ex := feather_web_rest.NotFoundException("principal not found in context")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	username := principal.(*feather_security.Principal).Username

	var err error
	var user *feather_security.Principal
	if user, err = endpoint.principalManager.Find(ctx.Request.Context(), *username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	user.Password, user.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, user)
}
