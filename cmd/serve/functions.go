package serve

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	feather_boot "github.com/guidomantilla/go-feather-boot/pkg/boot"
	feather_commons_config "github.com/guidomantilla/go-feather-commons/pkg/config"
	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/guidomantilla/go-feather-api-sample/pkg/config"
	"github.com/guidomantilla/go-feather-api-sample/pkg/endpoint/rpc"
	"github.com/guidomantilla/go-feather-api-sample/pkg/repositories"
	"github.com/guidomantilla/go-feather-api-sample/pkg/service"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	logger := feather_commons_log.Custom()
	ctx := context.Background()
	appName, version := "go-feather-api-sample", "v0.3.0"

	authPrincipalRepository := repositories.NewDefaultAuthPrincipalRepository()

	builder := feather_boot.NewBeanBuilder(ctx)
	builder.Config = func(appCtx *feather_boot.ApplicationContext) {
		var cfg config.Config
		if err := feather_commons_config.Process(ctx, appCtx.Environment, &cfg); err != nil {
			feather_commons_log.Fatal("starting up - error setting up configuration.", "message", err.Error())
		}

		appCtx.HttpConfig = &feather_boot.HttpConfig{
			Host: cfg.Host,
			Port: cfg.HttpPort,
		}

		appCtx.GrpcConfig = &feather_boot.GrpcConfig{
			Host: cfg.Host,
			Port: cfg.GrpcPort,
		}

		appCtx.SecurityConfig = &feather_boot.SecurityConfig{
			TokenSignatureKey: cfg.TokenSignatureKey,
		}

		appCtx.DatabaseConfig = &feather_boot.DatabaseConfig{
			ParamHolder:        feather_sql.UndefinedParamHolder.ValueFromName(*cfg.ParamHolder),
			Driver:             feather_sql.UndefinedDriverName.ValueFromName(*cfg.DatasourceDriver),
			DatasourceUrl:      cfg.DatasourceUrl,
			DatasourceServer:   cfg.DatasourceServer,
			DatasourceService:  cfg.DatasourceService,
			DatasourceUsername: cfg.DatasourceUsername,
			DatasourcePassword: cfg.DatasourcePassword,
		}
	}
	builder.PrincipalManager = func(appCtx *feather_boot.ApplicationContext) feather_security.PrincipalManager {
		return service.NewDBPrincipalManager(appCtx.TransactionHandler, authPrincipalRepository)
	}
	builder.GrpcServer = func(appCtx *feather_boot.ApplicationContext) (*grpc.ServiceDesc, any) {
		grpcServer := rpc.NewApiSampleGrpcServer(appCtx.AuthenticationService, appCtx.AuthorizationService, appCtx.PrincipalManager)
		return &rpc.ApiSample_ServiceDesc, grpcServer
	}
	err := feather_boot.Init(appName, version, args, logger, builder, func(appCtx feather_boot.ApplicationContext) error {

		appCtx.PrivateRouter.GET("/principal", func(ctx *gin.Context) {

			var principal any
			var exists bool
			if principal, exists = ctx.Get("principal"); !exists {
				ex := feather_web_rest.NotFoundException("principal not found in context")
				ctx.AbortWithStatusJSON(ex.Code, ex)
				return
			}

			username := principal.(*feather_security.Principal).Username

			var err error
			var user *feather_security.Principal
			if user, err = appCtx.PrincipalManager.Find(ctx.Request.Context(), *username); err != nil {
				ex := feather_web_rest.UnauthorizedException(err.Error())
				ctx.AbortWithStatusJSON(ex.Code, ex)
				return
			}

			user.Password, user.Passphrase = nil, nil
			ctx.JSON(http.StatusOK, user)
		})

		return nil
	})
	if err != nil {
		feather_commons_log.Fatal(err.Error())
	}
}
