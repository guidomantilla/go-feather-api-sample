package rest

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_security_validation "github.com/guidomantilla/go-feather-security/pkg/validation"
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

	var err error
	var principal *feather_security.Principal
	if err = ctx.ShouldBindJSON(&principal); err != nil {
		ex := feather_web_rest.BadRequestException("error unmarshalling request json to object")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if errs := endpoint.validate(principal); errs != nil {
		ex := feather_web_rest.BadRequestException("error validating the principal", errs...)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if err = endpoint.principalManager.Create(ctx.Request.Context(), principal); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (endpoint *DefaultAuthPrincipalEndpoint) Update(ctx *gin.Context) {

	var err error
	var principal *feather_security.Principal
	if err = ctx.ShouldBindJSON(&principal); err != nil {
		ex := feather_web_rest.BadRequestException("error unmarshalling request json to object")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if errs := endpoint.validate(principal); errs != nil {
		ex := feather_web_rest.BadRequestException("error validating the principal", errs...)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if err = endpoint.principalManager.Update(ctx.Request.Context(), principal); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	principal.Password, principal.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, principal)
}

func (endpoint *DefaultAuthPrincipalEndpoint) validate(principal *feather_security.Principal) []error {

	var errors []error

	if err := feather_security_validation.ValidateFieldIsRequired("this", "username", principal.Username); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldIsRequired("this", "role", principal.Role); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldIsRequired("this", "password", principal.Password); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldIsRequired("this", "passphrase", principal.Passphrase); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "enabled", principal.Enabled); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "non_locked", principal.NonLocked); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "non_expired", principal.NonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "password_non_expired", principal.PasswordNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "signup_done", principal.SignUpDone); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateStructIsRequired("this", "resources", principal.Resources); err != nil {
		errors = append(errors, err)
		return errors
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "token", principal.Token); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (endpoint *DefaultAuthPrincipalEndpoint) Delete(ctx *gin.Context) {

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

	username := ctx.Param("username")
	if err = endpoint.principalManager.Delete(ctx.Request.Context(), username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, nil)
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

	var principal *feather_security.Principal
	username := ctx.Param("username")
	if principal, err = endpoint.principalManager.Find(ctx.Request.Context(), username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	principal.Password, principal.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, principal)
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

	if principal, err = endpoint.principalManager.Find(ctx.Request.Context(), *principal.Username); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	principal.Password, principal.Passphrase = nil, nil
	ctx.JSON(http.StatusOK, principal)
}

func (endpoint *DefaultAuthPrincipalEndpoint) ChangePassword(ctx *gin.Context) {

}

func (endpoint *DefaultAuthPrincipalEndpoint) VerifyResource(ctx *gin.Context) {

}
