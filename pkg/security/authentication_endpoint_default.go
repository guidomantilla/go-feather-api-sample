package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultAuthenticationEndpoint struct {
	authenticationService AuthenticationService
}

func NewDefaultAuthenticationEndpoint(authenticationService AuthenticationService) *DefaultAuthenticationEndpoint {
	return &DefaultAuthenticationEndpoint{
		authenticationService: authenticationService,
	}
}

func (endpoint *DefaultAuthenticationEndpoint) Login(ctx *gin.Context) {

	var err error
	var principal *Principal
	if err = ctx.ShouldBindJSON(&principal); err != nil {
		ex := BadRequestException("error unmarshalling request json to object")
		ctx.JSON(ex.Code, ex)
		return
	}

	if errs := principal.Validate(); errs != nil {
		ex := BadRequestException("error validating the principal", errs...)
		ctx.JSON(ex.Code, ex)
		return
	}

	var token *string
	if token, err = endpoint.authenticationService.Authenticate(ctx.Request.Context(), principal); err != nil {
		ex := UnauthorizedException(err.Error())
		ctx.JSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
