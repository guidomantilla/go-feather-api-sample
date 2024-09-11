package serve

import (
	"context"
	"fmt"

	feather_boot "github.com/guidomantilla/go-feather-lib/pkg/boot"
	feather_commons_config "github.com/guidomantilla/go-feather-lib/pkg/config"
	feather_commons_log "github.com/guidomantilla/go-feather-lib/pkg/log"
	feather_security "github.com/guidomantilla/go-feather-lib/pkg/security"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/guidomantilla/go-feather-api-sample/pkg/endpoint/rest"
	"github.com/guidomantilla/go-feather-api-sample/pkg/endpoint/rpc"
	"github.com/guidomantilla/go-feather-api-sample/pkg/service"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	ctx := context.Background()
	logger := feather_commons_log.Custom()
	appName, version := feather_commons_config.Application, feather_commons_config.Version
	enablers := &feather_boot.Enablers{
		HttpServerEnabled: true,
		GrpcServerEnabled: true,
		DatabaseEnabled:   true,
	}

	builder := feather_boot.NewBeanBuilder(ctx)
	builder.Config = func(appCtx *feather_boot.ApplicationContext) {
		var cfg feather_commons_config.Config
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
			DatasourceUrl:      cfg.DatasourceUrl,
			DatasourceServer:   cfg.DatasourceServer,
			DatasourceService:  cfg.DatasourceService,
			DatasourceUsername: cfg.DatasourceUsername,
			DatasourcePassword: cfg.DatasourcePassword,
		}

	}
	builder.PrincipalManager = func(appCtx *feather_boot.ApplicationContext) feather_security.PrincipalManager {
		return service.NewDBPrincipalManager(appCtx.TransactionHandler, appCtx.PasswordManager)
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
		appCtx.PrivateRouter.DELETE("/principals/:username", authPrincipalEndpoint.Delete)
		appCtx.PrivateRouter.PATCH("/principals/change-password", authPrincipalEndpoint.ChangePassword)

		return nil
	})
	if err != nil {
		feather_commons_log.Fatal(err.Error())
	}
}
