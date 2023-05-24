package boot

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
	feather_commons_properties "github.com/guidomantilla/go-feather-commons/pkg/properties"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	"go.uber.org/zap"
)

const (
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME"
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

type RelationalDatasourceContextBuilderFunc func() feather_sql_datasource.RelationalDatasourceContext

type PasswordGeneratorBuilderFunc func() feather_security.PasswordGenerator

type PasswordEncoderBuilderFunc func() feather_security.PasswordEncoder

type PrincipalManagerBuilderFunc func(passwordManager feather_security.PasswordManager) feather_security.PrincipalManager

type TokenManagerBuilderFunc func(secret string) feather_security.TokenManager

type AuthenticationServiceBuilderFunc func(passwordEncoder feather_security.PasswordEncoder, principalManager feather_security.PrincipalManager, tokenManager feather_security.TokenManager) feather_security.AuthenticationService

type AuthorizationServiceBuilderFunc func(tokenManager feather_security.TokenManager, principalManager feather_security.PrincipalManager) feather_security.AuthorizationService

type AuthenticationEndpointBuilderFunc func(authenticationService feather_security.AuthenticationService) feather_security.AuthenticationEndpoint

type AuthorizationFilterBuilderFunc func(authorizationService feather_security.AuthorizationService) feather_security.AuthorizationFilter

type BeanBuilder struct {
	RelationalDatasourceContext RelationalDatasourceContextBuilderFunc
	PasswordEncoder             PasswordEncoderBuilderFunc
	PasswordGenerator           PasswordGeneratorBuilderFunc
	PrincipalManager            PrincipalManagerBuilderFunc
	TokenManager                TokenManagerBuilderFunc
	AuthenticationService       AuthenticationServiceBuilderFunc
	AuthorizationService        AuthorizationServiceBuilderFunc
	AuthenticationEndpoint      AuthenticationEndpointBuilderFunc
	AuthorizationFilter         AuthorizationFilterBuilderFunc
}

func NewBeanBuilder() *BeanBuilder {
	return &BeanBuilder{
		RelationalDatasourceContext: func() feather_sql_datasource.RelationalDatasourceContext {
			return nil
		},
		PasswordEncoder: func() feather_security.PasswordEncoder {
			return nil
		},
		PasswordGenerator: func() feather_security.PasswordGenerator {
			return nil
		},
		PrincipalManager: func(passwordManager feather_security.PasswordManager) feather_security.PrincipalManager {
			return nil
		},
		TokenManager: func(secret string) feather_security.TokenManager {
			return nil
		},
		AuthenticationService: func(passwordEncoder feather_security.PasswordEncoder, principalManager feather_security.PrincipalManager, tokenManager feather_security.TokenManager) feather_security.AuthenticationService {
			return nil
		},
		AuthorizationService: func(tokenManager feather_security.TokenManager, principalManager feather_security.PrincipalManager) feather_security.AuthorizationService {
			return nil
		},
		AuthenticationEndpoint: func(authenticationService feather_security.AuthenticationService) feather_security.AuthenticationEndpoint {
			return nil
		},
		AuthorizationFilter: func(authorizationService feather_security.AuthorizationService) feather_security.AuthorizationFilter {
			return nil
		},
	}
}

type ApplicationContext struct {
	AppName                     string
	Environment                 feather_commons_environment.Environment
	RelationalDatasourceContext feather_sql_datasource.RelationalDatasourceContext
	RelationalDatasource        feather_sql_datasource.RelationalDatasource
	PasswordEncoder             feather_security.PasswordEncoder
	PasswordGenerator           feather_security.PasswordGenerator
	PrincipalManager            feather_security.PrincipalManager
	TokenManager                feather_security.TokenManager
	AuthenticationService       feather_security.AuthenticationService
	AuthenticationEndpoint      feather_security.AuthenticationEndpoint
	AuthorizationService        feather_security.AuthorizationService
	AuthorizationFilter         feather_security.AuthorizationFilter
	Router                      *gin.Engine
	SecureRouter                *gin.RouterGroup
}

func NewApplicationContext(appName string, args []string, builder *BeanBuilder) *ApplicationContext {

	if appName == "" {
		zap.L().Fatal("starting up - error setting up the ApplicationContext: appName is empty")
	}

	zap.L().Info(fmt.Sprintf("starting up - starting up ApplicationContext %s", appName))

	if args == nil {
		zap.L().Fatal("starting up - error setting up the ApplicationContext: args is nil")
	}

	if builder == nil {
		zap.L().Fatal("starting up - error setting up the ApplicationContext: builder is nil")
	}

	ctx := &ApplicationContext{}
	ctx.AppName = appName

	buildEnvironment(ctx, args)
	buildDatasource(ctx, builder)
	buildSecurity(ctx, builder)

	ctx.Router = gin.Default()

	return ctx
}

func buildEnvironment(ctx *ApplicationContext, args []string) {

	zap.L().Info("starting up - setting up environment variables")

	osArgs := os.Environ()
	osSource := feather_commons_properties.NewDefaultPropertySource(OsPropertySourceName, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(&osArgs)))
	cmdSource := feather_commons_properties.NewDefaultPropertySource(CmdPropertySourceName, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(&args)))
	ctx.Environment = feather_commons_environment.NewDefaultEnvironment(feather_commons_environment.WithPropertySources(osSource, cmdSource))
}

func buildDatasource(ctx *ApplicationContext, builder *BeanBuilder) {

	zap.L().Info("starting up - setting up DB connection")

	paramHolderName := ctx.Environment.GetValueOrDefault(ParamHolder, EnvVarDefaultValuesMap[ParamHolder]).AsString()
	var paramHolder feather_sql.ParamHolder
	if paramHolder = feather_sql.UndefinedParamHolder.ValueFromName(paramHolderName); paramHolder == feather_sql.UndefinedParamHolder {
		zap.L().Fatal("starting up - error setting up DB config: invalid param holder")
	}

	driverName := ctx.Environment.GetValue(DatasourceDriver).AsString()
	var driver feather_sql.DriverName
	if driver = feather_sql.UndefinedDriverName.ValueFromName(driverName); driver == feather_sql.UndefinedDriverName {
		zap.L().Fatal("starting up - error setting up DB config: invalid driver name")
	}

	var url string
	if url = ctx.Environment.GetValue(DatasourceUrl).AsString(); strings.TrimSpace(url) == "" {
		zap.L().Fatal("starting up - error setting up DB config: url is empty")
	}

	var username string
	if username = ctx.Environment.GetValue(DatasourceUsername).AsString(); strings.TrimSpace(username) == "" {
		zap.L().Fatal("starting up - error setting up DB config: username is empty")
	}

	var password string
	if password = ctx.Environment.GetValue(DatasourcePassword).AsString(); strings.TrimSpace(password) == "" {
		zap.L().Fatal("starting up - error setting up DB config: password is empty")
	}

	var server string
	if server = ctx.Environment.GetValue(DatasourceServer).AsString(); strings.TrimSpace(server) == "" {
		zap.L().Fatal("starting up - error setting up DB config: server is empty")
	}

	var service string
	if service = ctx.Environment.GetValue(DatasourceService).AsString(); strings.TrimSpace(service) == "" {
		zap.L().Fatal("starting up - error setting up DB config: service is empty")
	}

	if ctx.RelationalDatasourceContext = builder.RelationalDatasourceContext(); ctx.RelationalDatasourceContext == nil {
		ctx.RelationalDatasourceContext = feather_sql_datasource.NewDefaultRelationalDatasourceContext(driver, paramHolder, url, username, password, server, service)
	}
	ctx.RelationalDatasource = feather_sql_datasource.NewDefaultRelationalDatasource(ctx.RelationalDatasourceContext, sql.Open)
}

func buildSecurity(ctx *ApplicationContext, builder *BeanBuilder) {

	zap.L().Info("starting up - setting up security")

	if ctx.PasswordEncoder = builder.PasswordEncoder(); ctx.PasswordEncoder == nil {
		ctx.PasswordEncoder = feather_security.NewBcryptPasswordEncoder()
	}
	if ctx.PasswordGenerator = builder.PasswordGenerator(); ctx.PasswordGenerator == nil {
		ctx.PasswordGenerator = feather_security.NewDefaultPasswordGenerator()
	}

	passwordEncoder := feather_security.NewDelegatingPasswordEncoder(ctx.PasswordEncoder)
	passwordManager := feather_security.NewDefaultPasswordManager(passwordEncoder, ctx.PasswordGenerator)
	if ctx.PrincipalManager = builder.PrincipalManager(passwordManager); ctx.PrincipalManager == nil {
		ctx.PrincipalManager = feather_security.NewInMemoryPrincipalManager(passwordManager)
	}

	secret := ctx.Environment.GetValueOrDefault(TokenSignatureKey, EnvVarDefaultValuesMap[TokenSignatureKey]).AsString()
	if ctx.TokenManager = builder.TokenManager(secret); ctx.TokenManager == nil {
		ctx.TokenManager = feather_security.NewJwtTokenManager([]byte(secret), feather_security.WithIssuer(ctx.AppName))
	}

	if ctx.AuthenticationService = builder.AuthenticationService(passwordManager, ctx.PrincipalManager, ctx.TokenManager); ctx.AuthenticationService == nil {
		ctx.AuthenticationService = feather_security.NewDefaultAuthenticationService(passwordManager, ctx.PrincipalManager, ctx.TokenManager)
	}
	if ctx.AuthenticationEndpoint = builder.AuthenticationEndpoint(ctx.AuthenticationService); ctx.AuthenticationEndpoint == nil {
		ctx.AuthenticationEndpoint = feather_security.NewDefaultAuthenticationEndpoint(ctx.AuthenticationService)
	}

	if ctx.AuthorizationService = builder.AuthorizationService(ctx.TokenManager, ctx.PrincipalManager); ctx.AuthorizationService == nil {
		ctx.AuthorizationService = feather_security.NewDefaultAuthorizationService(ctx.TokenManager, ctx.PrincipalManager)
	}
	if ctx.AuthorizationFilter = builder.AuthorizationFilter(ctx.AuthorizationService); ctx.AuthorizationFilter == nil {
		ctx.AuthorizationFilter = feather_security.NewDefaultAuthorizationFilter(ctx.AuthorizationService)
	}
}

func (ctx *ApplicationContext) Stop() {

	var err error

	if ctx.RelationalDatasource != nil && ctx.RelationalDatasourceContext != nil {

		var database *sql.DB
		zap.L().Info("shutting down - closing up db connection")

		if database, err = ctx.RelationalDatasource.GetDatabase(); err != nil {
			zap.L().Error(fmt.Sprintf("shutting down - error db connection: %s", err.Error()))
			return
		}

		if err = database.Close(); err != nil {
			zap.L().Error(fmt.Sprintf("shutting down - error closing db connection: %s", err.Error()))
			return
		}

		zap.L().Info("shutting down - db connection closed")
	}

	zap.L().Info(fmt.Sprintf("shutting down - ApplicationContext closed %s", ctx.AppName))
}
