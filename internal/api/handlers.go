package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"app-name": "go-feather-api-sample"})
}
