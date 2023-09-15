package boot

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
	feather_commons_properties "github.com/guidomantilla/go-feather-commons/pkg/properties"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	feather_sql_transaction "github.com/guidomantilla/go-feather-sql/pkg/transaction"
)

const (
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME" //nolint:gosec
	HostPort              = "HOST_PORT"
	TokenSignatureKey     = "TOKEN_SIGNATURE_KEY"
	ParamHolder           = "PARAM_HOLDER"
	DatasourceDriver      = "DATASOURCE_DRIVER"
	DatasourceUsername    = "DATASOURCE_USERNAME"
	DatasourcePassword    = "DATASOURCE_PASSWORD"
	DatasourceServer      = "DATASOURCE_SERVER"
	DatasourceService     = "DATASOURCE_SERVICE"
	DatasourceUrl         = "DATASOURCE_URL"
)

var (
	EnvVarDefaultValuesMap = map[string]string{
		HostPort:          ":8080",
		TokenSignatureKey: "SecretYouShouldHide",
		ParamHolder:       "named",
	}
)

type EnvironmentBuilderFunc func(appCtx *ApplicationContext) feather_commons_environment.Environment

type DatasourceContextBuilderFunc func(appCtx *ApplicationContext) feather_sql_datasource.DatasourceContext

type DatasourceBuilderFunc func(appCtx *ApplicationContext) feather_sql_datasource.Datasource

type TransactionHandlerBuilderFunc func(appCtx *ApplicationContext) feather_sql_transaction.TransactionHandler

type PasswordGeneratorBuilderFunc func(appCtx *ApplicationContext) feather_security.PasswordGenerator

type PasswordEncoderBuilderFunc func(appCtx *ApplicationContext) feather_security.PasswordEncoder

type PasswordManagerBuilderFunc func(appCtx *ApplicationContext) feather_security.PasswordManager

type PrincipalManagerBuilderFunc func(appCtx *ApplicationContext) feather_security.PrincipalManager

type TokenManagerBuilderFunc func(appCtx *ApplicationContext) feather_security.TokenManager

type AuthenticationServiceBuilderFunc func(appCtx *ApplicationContext) feather_security.AuthenticationService

type AuthorizationServiceBuilderFunc func(appCtx *ApplicationContext) feather_security.AuthorizationService

type AuthenticationEndpointBuilderFunc func(appCtx *ApplicationContext) feather_security.AuthenticationEndpoint

type AuthorizationFilterBuilderFunc func(appCtx *ApplicationContext) feather_security.AuthorizationFilter

type BeanBuilder struct {
	Environment            EnvironmentBuilderFunc
	DatasourceContext      DatasourceContextBuilderFunc
	Datasource             DatasourceBuilderFunc
	TransactionHandler     TransactionHandlerBuilderFunc
	PasswordEncoder        PasswordEncoderBuilderFunc
	PasswordGenerator      PasswordGeneratorBuilderFunc
	PasswordManager        PasswordManagerBuilderFunc
	PrincipalManager       PrincipalManagerBuilderFunc
	TokenManager           TokenManagerBuilderFunc
	AuthenticationService  AuthenticationServiceBuilderFunc
	AuthorizationService   AuthorizationServiceBuilderFunc
	AuthenticationEndpoint AuthenticationEndpointBuilderFunc
	AuthorizationFilter    AuthorizationFilterBuilderFunc
}

func NewBeanBuilder() *BeanBuilder {

	return &BeanBuilder{

		Environment: func(appCtx *ApplicationContext) feather_commons_environment.Environment {
			osArgs := os.Environ()
			osSource := feather_commons_properties.NewDefaultPropertySource(OsPropertySourceName, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(&osArgs)))
			cmdSource := feather_commons_properties.NewDefaultPropertySource(CmdPropertySourceName, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(&appCtx.CmdArgs)))
			return feather_commons_environment.NewDefaultEnvironment(feather_commons_environment.WithPropertySources(osSource, cmdSource))
		},

		DatasourceContext: func(appCtx *ApplicationContext) feather_sql_datasource.DatasourceContext {

			paramHolderName := appCtx.Environment.GetValueOrDefault(ParamHolder, EnvVarDefaultValuesMap[ParamHolder]).AsString()
			var paramHolder feather_sql.ParamHolder
			if paramHolder = feather_sql.UndefinedParamHolder.ValueFromName(paramHolderName); paramHolder == feather_sql.UndefinedParamHolder {
				slog.Error("starting up - error setting up DB config: invalid param holder")
				os.Exit(1)
			}

			driverName := appCtx.Environment.GetValue(DatasourceDriver).AsString()
			var driver feather_sql.DriverName
			if driver = feather_sql.UndefinedDriverName.ValueFromName(driverName); driver == feather_sql.UndefinedDriverName {
				slog.Error("starting up - error setting up DB config: invalid driver name")
				os.Exit(1)
			}

			var url string
			if url = appCtx.Environment.GetValue(DatasourceUrl).AsString(); strings.TrimSpace(url) == "" {
				slog.Error("starting up - error setting up DB config: url is empty")
				os.Exit(1)
			}

			var username string
			if username = appCtx.Environment.GetValue(DatasourceUsername).AsString(); strings.TrimSpace(username) == "" {
				slog.Error("starting up - error setting up DB config: username is empty")
				os.Exit(1)
			}

			var password string
			if password = appCtx.Environment.GetValue(DatasourcePassword).AsString(); strings.TrimSpace(password) == "" {
				slog.Error("starting up - error setting up DB config: password is empty")
				os.Exit(1)
			}

			var server string
			if server = appCtx.Environment.GetValue(DatasourceServer).AsString(); strings.TrimSpace(server) == "" {
				slog.Error("starting up - error setting up DB config: server is empty")
				os.Exit(1)
			}

			var service string
			if service = appCtx.Environment.GetValue(DatasourceService).AsString(); strings.TrimSpace(service) == "" {
				slog.Error("starting up - error setting up DB config: service is empty")
				os.Exit(1)
			}

			return feather_sql_datasource.NewDefaultDatasourceContext(driver, paramHolder, url, username, password, server, service)
		},

		Datasource: func(appCtx *ApplicationContext) feather_sql_datasource.Datasource {
			return feather_sql_datasource.NewDefaultDatasource(appCtx.DatasourceContext, sql.Open)
		},

		TransactionHandler: func(appCtx *ApplicationContext) feather_sql_transaction.TransactionHandler {
			return feather_sql_transaction.NewTransactionHandler(appCtx.Datasource)
		},

		PasswordEncoder: func(appCtx *ApplicationContext) feather_security.PasswordEncoder {
			return feather_security.NewBcryptPasswordEncoder()
		},

		PasswordGenerator: func(appCtx *ApplicationContext) feather_security.PasswordGenerator {
			return feather_security.NewDefaultPasswordGenerator()
		},

		PasswordManager: func(appCtx *ApplicationContext) feather_security.PasswordManager {
			return feather_security.NewDefaultPasswordManager(appCtx.PasswordEncoder, appCtx.PasswordGenerator)
		},

		PrincipalManager: func(appCtx *ApplicationContext) feather_security.PrincipalManager {
			return feather_security.NewInMemoryPrincipalManager(appCtx.PasswordManager)
		},

		TokenManager: func(appCtx *ApplicationContext) feather_security.TokenManager {

			secret := appCtx.Environment.GetValueOrDefault(TokenSignatureKey, EnvVarDefaultValuesMap[TokenSignatureKey]).AsString()
			return feather_security.NewJwtTokenManager([]byte(secret), feather_security.WithIssuer(appCtx.AppName))
		},

		AuthenticationService: func(appCtx *ApplicationContext) feather_security.AuthenticationService {
			return feather_security.NewDefaultAuthenticationService(appCtx.PasswordManager, appCtx.PrincipalManager, appCtx.TokenManager)
		},
		AuthorizationService: func(appCtx *ApplicationContext) feather_security.AuthorizationService {
			return feather_security.NewDefaultAuthorizationService(appCtx.TokenManager, appCtx.PrincipalManager)
		},
		AuthenticationEndpoint: func(appCtx *ApplicationContext) feather_security.AuthenticationEndpoint {
			return feather_security.NewDefaultAuthenticationEndpoint(appCtx.AuthenticationService)
		},
		AuthorizationFilter: func(appCtx *ApplicationContext) feather_security.AuthorizationFilter {
			return feather_security.NewDefaultAuthorizationFilter(appCtx.AuthorizationService)
		},
	}
}

type ApplicationContext struct {
	AppName                string
	CmdArgs                []string
	Environment            feather_commons_environment.Environment
	DatasourceContext      feather_sql_datasource.DatasourceContext
	Datasource             feather_sql_datasource.Datasource
	TransactionHandler     feather_sql_transaction.TransactionHandler
	PasswordEncoder        feather_security.PasswordEncoder
	PasswordGenerator      feather_security.PasswordGenerator
	PasswordManager        feather_security.PasswordManager
	PrincipalManager       feather_security.PrincipalManager
	TokenManager           feather_security.TokenManager
	AuthenticationService  feather_security.AuthenticationService
	AuthenticationEndpoint feather_security.AuthenticationEndpoint
	AuthorizationService   feather_security.AuthorizationService
	AuthorizationFilter    feather_security.AuthorizationFilter
	Router                 *gin.Engine
	SecureRouter           *gin.RouterGroup
}

func NewApplicationContext(appName string, args []string, builder *BeanBuilder) *ApplicationContext {

	if appName == "" {
		slog.Error("starting up - error setting up the ApplicationContext: appName is empty")
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("starting up - starting up ApplicationContext %s", appName))

	if args == nil {
		slog.Error("starting up - error setting up the ApplicationContext: args is nil")
		os.Exit(1)
	}

	if builder == nil {
		slog.Error("starting up - error setting up the ApplicationContext: builder is nil")
		os.Exit(1)
	}

	ctx := &ApplicationContext{}
	ctx.AppName, ctx.CmdArgs = appName, args

	slog.Info("starting up - setting up environment variables")
	ctx.Environment = builder.Environment(ctx)

	slog.Info("starting up - setting up DB connection")
	ctx.DatasourceContext = builder.DatasourceContext(ctx)
	ctx.Datasource = builder.Datasource(ctx)
	ctx.TransactionHandler = builder.TransactionHandler(ctx)

	slog.Info("starting up - setting up security")
	ctx.PasswordEncoder = builder.PasswordEncoder(ctx)
	ctx.PasswordGenerator = builder.PasswordGenerator(ctx)
	ctx.PasswordManager = builder.PasswordManager(ctx)
	ctx.PrincipalManager, ctx.TokenManager = builder.PrincipalManager(ctx), builder.TokenManager(ctx)
	ctx.AuthenticationService, ctx.AuthorizationService = builder.AuthenticationService(ctx), builder.AuthorizationService(ctx)
	ctx.AuthenticationEndpoint, ctx.AuthorizationFilter = builder.AuthenticationEndpoint(ctx), builder.AuthorizationFilter(ctx)

	ctx.Router = gin.Default()

	return ctx
}

func (ctx *ApplicationContext) Stop() {

	var err error

	if ctx.Datasource != nil && ctx.DatasourceContext != nil {

		var database *sql.DB
		slog.Info("shutting down - closing up db connection")

		if database, err = ctx.Datasource.GetDatabase(); err != nil {
			slog.Error(fmt.Sprintf("shutting down - error db connection: %s", err.Error()))
			return
		}

		if err = database.Close(); err != nil {
			slog.Error(fmt.Sprintf("shutting down - error closing db connection: %s", err.Error()))
			return
		}

		slog.Info("shutting down - db connection closed")
	}

	slog.Info(fmt.Sprintf("shutting down - ApplicationContext closed %s", ctx.AppName))
}
