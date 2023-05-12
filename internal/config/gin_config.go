package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-api-sample/pkg/server"
)

func InitGinServer(environment environment.Environment, router *gin.Engine) lifecycle.Server {

	hostAddress := environment.GetValueOrDefault(HostPort, EnvVarDefaultValuesMap[HostPort]).AsString()
	httpServer := &http.Server{
		Addr:              hostAddress,
		Handler:           router,
		ReadHeaderTimeout: 60000,
	}

	ginServer := server.BuildHttpServer(httpServer)
	return ginServer
}
