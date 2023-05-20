package serve

import (
	"net/http"
	"syscall"

	"github.com/gin-gonic/gin"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	"github.com/qmdx00/lifecycle"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-api-sample/internal/config"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	app := lifecycle.NewApp(
		lifecycle.WithName(config.AppName),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	//orchestrator.Cleanup(config.Init)

	//

	environment := config.InitEnv(args)
	passwordManager := config.InitPwd(environment)
	tokenTokenManager := config.InitToken(environment)
	principalManager := config.InitPrincipal(environment, passwordManager)
	authenticationEndpoint, authorizationFilter := config.InitAuthEndpoints(environment, tokenTokenManager,
		principalManager.(feather_security.AuthenticationDelegate), principalManager.(feather_security.AuthorizationDelegate))

	router := gin.Default()
	router.POST("/login", authenticationEndpoint.Authenticate)
	router.NoRoute(authorizationFilter.Authorize, func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiGroup := router.Group("/api")
	apiGroup.Use(authorizationFilter.Authorize)
	apiGroup.GET("/info", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "we are ok"})
	})

	//

	var err error
	app.Attach("GinServer", config.InitGinServer(environment, router))
	if err = app.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}

	_ = logger.Sync()
}
