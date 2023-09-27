package rest

import "github.com/gin-gonic/gin"

type DefaultAuthResourceEndpoint struct {
}

func NewDefaultAuthResourceEndpoint() *DefaultAuthResourceEndpoint {
	return &DefaultAuthResourceEndpoint{}
}

func (endpoint *DefaultAuthResourceEndpoint) Create(ctx *gin.Context) {

}

func (endpoint *DefaultAuthResourceEndpoint) Update(ctx *gin.Context) {

}

func (endpoint *DefaultAuthResourceEndpoint) Delete(ctx *gin.Context) {

}

func (endpoint *DefaultAuthResourceEndpoint) FindById(ctx *gin.Context) {

}

func (endpoint *DefaultAuthResourceEndpoint) FindAll(ctx *gin.Context) {

}
