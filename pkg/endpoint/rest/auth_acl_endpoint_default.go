package rest

import "github.com/gin-gonic/gin"

type DefaultAuthAclEndpoint struct {
}

func NewDefaultAuthAclEndpoint() *DefaultAuthAclEndpoint {
	return &DefaultAuthAclEndpoint{}
}

func (endpoint *DefaultAuthAclEndpoint) Create(ctx *gin.Context) {

}

func (endpoint *DefaultAuthAclEndpoint) Update(ctx *gin.Context) {

}

func (endpoint *DefaultAuthAclEndpoint) Delete(ctx *gin.Context) {

}

func (endpoint *DefaultAuthAclEndpoint) FindById(ctx *gin.Context) {

}

func (endpoint *DefaultAuthAclEndpoint) FindAll(ctx *gin.Context) {

}
