package boot

import (
	"net/http"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	feather_sql_config "github.com/guidomantilla/go-feather-sql/pkg/config"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
	feather_web_server "github.com/guidomantilla/go-feather-web/pkg/server"
	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

func Init(appName string, version string, args []string, builder *BeanBuilder, fn func(ctx ApplicationContext)) error {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

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
		lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	ctx := NewApplicationContext(strings.Join([]string{appName, version}, " - "), args, builder)
	app.Cleanup(feather_sql_config.Stop)

	ctx.Router.POST("/login", ctx.AuthenticationEndpoint.Authenticate)
	ctx.Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "alive"})
	})
	ctx.Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, feather_web_rest.NotFoundException("resource not found"))
	})
	ctx.SecureRouter = ctx.Router.Group("/api", ctx.AuthorizationFilter.Authorize)

	fn(*ctx)

	hostAddress := ctx.Environment.GetValueOrDefault(HostPort, EnvVarDefaultValuesMap[HostPort]).AsString()
	httpServer := &http.Server{
		Addr:              hostAddress,
		Handler:           ctx.Router,
		ReadHeaderTimeout: 60000,
	}

	app.Attach("GinServer", feather_web_server.BuildHttpServer(httpServer))
	return app.Run()
}
