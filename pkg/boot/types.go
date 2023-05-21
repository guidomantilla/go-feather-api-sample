package boot

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"github.com/guidomantilla/go-feather-commons/pkg/properties"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	"go.uber.org/zap"
)

const (
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME"
	HostPort              = "HOST_PORT"
	TokenSignatureKey     = "TOKEN_SIGNATURE_KEY"
)

var (
	EnvVarDefaultValuesMap = map[string]string{
		HostPort:          ":8080",
		TokenSignatureKey: "SecretYouShouldHide",
	}
)

type PasswordGeneratorBuilderFunc func() feather_security.PasswordGenerator

type PasswordEncoderBuilderFunc func() feather_security.PasswordEncoder

type PrincipalManagerBuilderFunc func(passwordManager feather_security.PasswordManager) feather_security.PrincipalManager

type TokenManagerBuilderFunc func(secret string) feather_security.TokenManager

type AuthenticationDelegateBuilderFunc func() feather_security.AuthenticationService

type AuthenticationServiceBuilderFunc func(tokenManager feather_security.TokenManager, authenticationDelegate feather_security.AuthenticationService) feather_security.AuthenticationService

type AuthorizationDelegateBuilderFunc func() feather_security.AuthorizationDelegate

type AuthorizationServiceBuilderFunc func(tokenManager feather_security.TokenManager, authorizationDelegate feather_security.AuthorizationDelegate) feather_security.AuthorizationService

type AuthenticationEndpointBuilderFunc func(authenticationService feather_security.AuthenticationService) feather_security.AuthenticationEndpoint

type AuthorizationFilterBuilderFunc func(authorizationService feather_security.AuthorizationService) feather_security.AuthorizationFilter

type BeanBuilder struct {
	PasswordEncoder        PasswordEncoderBuilderFunc
	PasswordGenerator      PasswordGeneratorBuilderFunc
	PrincipalManager       PrincipalManagerBuilderFunc
	TokenManager           TokenManagerBuilderFunc
	AuthenticationDelegate AuthenticationDelegateBuilderFunc
	AuthenticationService  AuthenticationServiceBuilderFunc
	AuthorizationDelegate  AuthorizationDelegateBuilderFunc
	AuthorizationService   AuthorizationServiceBuilderFunc
	AuthenticationEndpoint AuthenticationEndpointBuilderFunc
	AuthorizationFilter    AuthorizationFilterBuilderFunc
}

func NewBeanBuilder() *BeanBuilder {
	return &BeanBuilder{
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
		AuthenticationDelegate: func() feather_security.AuthenticationService {
			return nil
		},
		AuthenticationService: func(tokenManager feather_security.TokenManager, authenticationDelegate feather_security.AuthenticationService) feather_security.AuthenticationService {
			return nil
		},
		AuthorizationDelegate: func() feather_security.AuthorizationDelegate {
			return nil
		},
		AuthorizationService: func(tokenManager feather_security.TokenManager, authorizationDelegate feather_security.AuthorizationDelegate) feather_security.AuthorizationService {
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
	AppName                string
	Environment            environment.Environment
	PasswordEncoder        feather_security.PasswordEncoder
	PasswordGenerator      feather_security.PasswordGenerator
	PrincipalManager       feather_security.PrincipalManager
	TokenManager           feather_security.TokenManager
	AuthenticationDelegate feather_security.AuthenticationService
	AuthenticationService  feather_security.AuthenticationService
	AuthenticationEndpoint feather_security.AuthenticationEndpoint
	AuthorizationDelegate  feather_security.AuthorizationDelegate
	AuthorizationService   feather_security.AuthorizationService
	AuthorizationFilter    feather_security.AuthorizationFilter
	Router                 *gin.Engine
	SecureRouter           *gin.RouterGroup
}

func NewApplicationContext(appName string, args []string, builder *BeanBuilder) *ApplicationContext {

	if appName == "" {
		zap.L().Fatal("starting up - error setting up the ApplicationContext: appName is empty")
	}

	if args == nil {
		zap.L().Fatal("starting up - error setting up the ApplicationContext: args is nil")
	}

	if builder == nil {
		zap.L().Fatal("starting up - error setting up the ApplicationContext: builder is nil")
	}

	ctx := ApplicationContext{}
	ctx.AppName = appName

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OsPropertySourceName, properties.NewDefaultProperties(properties.FromArray(&osArgs)))
	cmdSource := properties.NewDefaultPropertySource(CmdPropertySourceName, properties.NewDefaultProperties(properties.FromArray(&args)))
	ctx.Environment = environment.NewDefaultEnvironment(environment.WithPropertySources(osSource, cmdSource))

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

	if ctx.AuthenticationDelegate = builder.AuthenticationDelegate(); ctx.AuthenticationDelegate == nil {
		ctx.AuthenticationDelegate = feather_security.NewDefaultAuthenticationDelegate(passwordManager, ctx.PrincipalManager)
	}
	if ctx.AuthenticationService = builder.AuthenticationService(ctx.TokenManager, ctx.AuthenticationDelegate); ctx.AuthenticationService == nil {
		ctx.AuthenticationService = feather_security.NewDefaultAuthenticationService(ctx.TokenManager, ctx.AuthenticationDelegate)
	}
	if ctx.AuthenticationEndpoint = builder.AuthenticationEndpoint(ctx.AuthenticationService); ctx.AuthenticationEndpoint == nil {
		ctx.AuthenticationEndpoint = feather_security.NewDefaultAuthenticationEndpoint(ctx.AuthenticationService)
	}

	if ctx.AuthorizationDelegate = builder.AuthorizationDelegate(); ctx.AuthorizationDelegate == nil {
		ctx.AuthorizationDelegate = feather_security.NewDefaultAuthorizationDelegate(ctx.PrincipalManager)
	}
	if ctx.AuthorizationService = builder.AuthorizationService(ctx.TokenManager, ctx.AuthorizationDelegate); ctx.AuthorizationService == nil {
		ctx.AuthorizationService = feather_security.NewDefaultAuthorizationService(ctx.TokenManager, ctx.AuthorizationDelegate)
	}
	if ctx.AuthorizationFilter = builder.AuthorizationFilter(ctx.AuthorizationService); ctx.AuthorizationFilter == nil {
		ctx.AuthorizationFilter = feather_security.NewDefaultAuthorizationFilter(ctx.AuthorizationService)
	}

	ctx.Router = gin.Default()
	ctx.Router.POST("/login", ctx.AuthenticationEndpoint.Authenticate)
	ctx.Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "alive"})
	})
	ctx.Router.NoRoute(ctx.AuthorizationFilter.Authorize, func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	ctx.SecureRouter = ctx.Router.Group("/api", ctx.AuthorizationFilter.Authorize)

	return &ctx
}
