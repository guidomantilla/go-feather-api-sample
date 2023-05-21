package boot

import (
	"net/http"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"github.com/guidomantilla/go-feather-web/pkg/server"
	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

func Init(appName string, args []string, builder *BeanBuilder, fn func(ctx ApplicationContext)) error {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			zap.L().Error(err.Error())
		}
	}(logger)

	if appName == "" {
		zap.L().Fatal("starting up - error setting up the application: appName is empty")
	}

	if args == nil {
		zap.L().Fatal("starting up - error setting up the application: args is nil")
	}

	if builder == nil {
		zap.L().Fatal("starting up - error setting up the application: builder is nil")
	}

	if fn == nil {
		zap.L().Fatal("starting up - error setting up the application: fn is nil")
	}

	app := lifecycle.NewApp(
		lifecycle.WithName(appName),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	app.Cleanup() //TODO

	ctx := NewApplicationContext(appName, args, builder)
	ctx.Router.POST("/login", ctx.AuthenticationEndpoint.Authenticate)
	ctx.Router.NoRoute(ctx.AuthorizationFilter.Authorize, func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiGroup := ctx.Router.Group("/api")
	apiGroup.Use(ctx.AuthorizationFilter.Authorize)
	apiGroup.GET("/info", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "we are ok"})
	})

	fn(*ctx)

	app.Attach("GinServer", initGinServer(ctx.Environment, ctx.Router))
	return app.Run()
}

func initGinServer(environment environment.Environment, router *gin.Engine) lifecycle.Server {

	hostAddress := environment.GetValueOrDefault(HostPort, EnvVarDefaultValuesMap[HostPort]).AsString()
	httpServer := &http.Server{
		Addr:              hostAddress,
		Handler:           router,
		ReadHeaderTimeout: 60000,
	}

	ginServer := server.BuildHttpServer(httpServer)
	return ginServer
}
