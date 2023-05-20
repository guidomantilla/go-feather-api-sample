package serve

import (
	"context"
	"net/http"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	feather_commons_config "github.com/guidomantilla/go-feather-commons/pkg/config"
	feather_commons_security "github.com/guidomantilla/go-feather-commons/pkg/security"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
	"github.com/qmdx00/lifecycle"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-api-sample/internal/config"
	"github.com/guidomantilla/go-feather-api-sample/pkg/security"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	appName := "go-feather-api-sample"

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	app := lifecycle.NewApp(
		lifecycle.WithName(appName),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	//orchestrator.Cleanup(config.Init)

	//

	environment := feather_commons_config.Init(args)

	passwordGenerator := feather_commons_security.NewDefaultPasswordGenerator()
	bcryptPasswordEncoder := feather_commons_security.NewBcryptPasswordEncoder()
	passwordEncoder := feather_commons_security.NewDelegatingPasswordEncoder(bcryptPasswordEncoder)
	passwordManager := feather_commons_security.NewDefaultPasswordManager(passwordEncoder, passwordGenerator)
	tokenTokenManager := security.NewDefaultTokenTokenManager(appName, time.Hour*24, []byte("SecretYouShouldHide"), jwt.SigningMethodHS512)

	root := &security.Principal{
		Username:           feather_commons_util.ValueToPtr("root"),
		Password:           feather_commons_util.ValueToPtr("RaveN123qweasd*+"),
		AccountNonExpired:  feather_commons_util.ValueToPtr(true),
		AccountNonLocked:   feather_commons_util.ValueToPtr(true),
		PasswordNonExpired: feather_commons_util.ValueToPtr(true),
		Enabled:            feather_commons_util.ValueToPtr(true),
		SignUpDone:         feather_commons_util.ValueToPtr(true),
		Authorities:        nil,
	}
	principalManager := security.NewInMemoryPrincipalManager(passwordManager)

	var err error
	if err = principalManager.Create(context.Background(), root); err != nil {
		zap.L().Fatal(err.Error())
		return
	}

	authenticationService := security.NewDefaultAuthenticationService(tokenTokenManager, principalManager)
	authenticationEndpoint := security.NewDefaultAuthenticationEndpoint(authenticationService)

	authorizationService := security.NewDefaultAuthorizationService(tokenTokenManager, principalManager)
	authorizationFilter := security.NewDefaultAuthorizationFilter(authorizationService)

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

	app.Attach("GinServer", config.InitGinServer(environment, router))
	if err = app.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}

	_ = logger.Sync()
}
