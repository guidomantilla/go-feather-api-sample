package rest

import "github.com/gin-gonic/gin"

type DefaultAuthUserEndpoint struct {
}

func NewDefaultAuthUserEndpoint() *DefaultAuthUserEndpoint {
	return &DefaultAuthUserEndpoint{}
}

func (endpoint *DefaultAuthUserEndpoint) Create(ctx *gin.Context) {

}

func (endpoint *DefaultAuthUserEndpoint) Update(ctx *gin.Context) {

}

func (endpoint *DefaultAuthUserEndpoint) Delete(ctx *gin.Context) {

}

func (endpoint *DefaultAuthUserEndpoint) FindById(ctx *gin.Context) {

}

func (endpoint *DefaultAuthUserEndpoint) FindAll(ctx *gin.Context) {

}
