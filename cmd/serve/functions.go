package serve

import (
	"context"

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

		authResourceEndpoint := rest.NewDefaultAuthResourceEndpoint()
		appCtx.PrivateRouter.GET("/resources", authResourceEndpoint.FindAll)
		appCtx.PrivateRouter.GET("/resources/:id", authResourceEndpoint.FindById)
		appCtx.PrivateRouter.POST("/resources", authResourceEndpoint.Create)
		appCtx.PrivateRouter.PUT("/resources/:id", authResourceEndpoint.Update)
		appCtx.PrivateRouter.DELETE("/resources/:id", authResourceEndpoint.Delete)

		authRoleEndpoint := rest.NewDefaultAuthRoleEndpoint()
		appCtx.PrivateRouter.GET("/roles", authRoleEndpoint.FindAll)
		appCtx.PrivateRouter.GET("/roles/:id", authRoleEndpoint.FindById)
		appCtx.PrivateRouter.POST("/roles", authRoleEndpoint.Create)
		appCtx.PrivateRouter.PUT("/roles/:id", authRoleEndpoint.Update)
		appCtx.PrivateRouter.DELETE("/roles/:id", authRoleEndpoint.Delete)

		authAclEndpoint := rest.NewDefaultAuthAclEndpoint()
		appCtx.PrivateRouter.GET("/acls", authAclEndpoint.FindAll)
		appCtx.PrivateRouter.GET("/acls/:id", authAclEndpoint.FindById)
		appCtx.PrivateRouter.POST("/acls", authAclEndpoint.Create)
		appCtx.PrivateRouter.PUT("/acls/:id", authAclEndpoint.Update)
		appCtx.PrivateRouter.DELETE("/acls/:id", authAclEndpoint.Delete)

		authUserEndpoint := rest.NewDefaultAuthUserEndpoint()
		appCtx.PrivateRouter.GET("/users", authUserEndpoint.FindAll)
		appCtx.PrivateRouter.GET("/users/:id", authUserEndpoint.FindById)
		appCtx.PrivateRouter.POST("/users", authUserEndpoint.Create)
		appCtx.PrivateRouter.PUT("/users/:id", authUserEndpoint.Update)
		appCtx.PrivateRouter.DELETE("/users/:id", authUserEndpoint.Delete)

		authPrincipalEndpoint := rest.NewDefaultAuthPrincipalEndpoint(appCtx.PrincipalManager)
		appCtx.PrivateRouter.GET("/principal", authPrincipalEndpoint.GetCurrentPrincipal)

		return nil
	})
	if err != nil {
		feather_commons_log.Fatal(err.Error())
	}
}
