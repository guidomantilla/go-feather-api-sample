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

	var exists bool
	var principal *feather_security.Principal
	if principal, exists = feather_security.GetPrincipalFROMContext(ctx); !exists {
		ex := feather_web_rest.NotFoundException("principal not found in context")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var err error
	var user *feather_security.Principal
	if user, err = endpoint.principalManager.Find(ctx.Request.Context(), *principal.Username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	user.Password, user.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, user)
}
