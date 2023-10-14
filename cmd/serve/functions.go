package serve

import (
	"context"
	"fmt"

	feather_boot "github.com/guidomantilla/go-feather-boot/pkg/boot"
	feather_commons_config "github.com/guidomantilla/go-feather-commons/pkg/config"
	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/guidomantilla/go-feather-api-sample/pkg/config"
	"github.com/guidomantilla/go-feather-api-sample/pkg/endpoint/rest"
	"github.com/guidomantilla/go-feather-api-sample/pkg/endpoint/rpc"
	"github.com/guidomantilla/go-feather-api-sample/pkg/repositories"
	"github.com/guidomantilla/go-feather-api-sample/pkg/service"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	ctx := context.Background()
	logger := feather_commons_log.Custom()
	appName, version := config.Application, config.Version
	enablers := &feather_boot.Enablers{
		HttpServerEnabled: true,
		GrpcServerEnabled: true,
		DatabaseEnabled:   true,
	}

	repository := repositories.NewDefaultRepository()

	builder := feather_boot.NewBeanBuilder(ctx)
	builder.Config = func(appCtx *feather_boot.ApplicationContext) {
		var cfg config.Config
		if err := feather_commons_config.Process(ctx, appCtx.Environment, &cfg); err != nil {
			feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up configuration: %s", err.Error()))
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
			TokenSignatureKey:    cfg.TokenSignatureKey,
			TokenVerificationKey: cfg.TokenSignatureKey,
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
		return service.NewDBPrincipalManager(appCtx.TransactionHandler, appCtx.PasswordManager, repository)
	}
	builder.GrpcServer = func(appCtx *feather_boot.ApplicationContext) (*grpc.ServiceDesc, any) {
		grpcServer := rpc.NewApiSampleGrpcServer(appCtx.AuthenticationService, appCtx.AuthorizationService, appCtx.PrincipalManager)
		return &rpc.ApiSample_ServiceDesc, grpcServer
	}
	err := feather_boot.Init(appName, version, args, logger, enablers, builder, func(appCtx feather_boot.ApplicationContext) error {

		authPrincipalEndpoint := rest.NewDefaultAuthPrincipalEndpoint(appCtx.PrincipalManager)
		appCtx.PrivateRouter.GET("/principals/current", authPrincipalEndpoint.FindCurrent)
		appCtx.PrivateRouter.GET("/principals/:username", authPrincipalEndpoint.FindByUsername)
		appCtx.PrivateRouter.POST("/principals", authPrincipalEndpoint.Create)
		appCtx.PrivateRouter.PUT("/principals", authPrincipalEndpoint.Update)
		appCtx.PrivateRouter.DELETE("/principals", authPrincipalEndpoint.Delete)
		appCtx.PrivateRouter.PATCH("/principals/change-password", authPrincipalEndpoint.ChangePassword)

		return nil
	})
	if err != nil {
		feather_commons_log.Fatal(err.Error())
	}
}
