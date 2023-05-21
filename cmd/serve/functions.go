package serve

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-api-sample/pkg/boot"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	appName := "go-feather-api-sample"
	builder := boot.NewBeanBuilder()

	builder.AuthenticationDelegate = func() feather_security.AuthenticationService {

		return nil
	}

	builder.AuthorizationDelegate = func() feather_security.AuthorizationDelegate {

		return nil
	}

	err := boot.Init(appName, "1.0", args, builder, func(ctx boot.ApplicationContext) {

		root := &feather_security.Principal{
			Username:           feather_commons_util.ValueToPtr("raven"),
			Password:           feather_commons_util.ValueToPtr("RaveN123qweasd*+"),
			Role:               feather_commons_util.ValueToPtr("ROOT"),
			AccountNonExpired:  feather_commons_util.ValueToPtr(true),
			AccountNonLocked:   feather_commons_util.ValueToPtr(true),
			PasswordNonExpired: feather_commons_util.ValueToPtr(true),
			Enabled:            feather_commons_util.ValueToPtr(true),
			SignUpDone:         feather_commons_util.ValueToPtr(true),
			Authorities: []feather_security.GrantedAuthority{
				{
					Name:      feather_commons_util.ValueToPtr("name"),
					Resources: []string{"GET /api/info"},
				},
				{
					Name:      feather_commons_util.ValueToPtr("name"),
					Resources: []string{"GET /api/info", "GET /api/xxx"},
				},
				{
					Name:      feather_commons_util.ValueToPtr("name"),
					Resources: []string{"GET /api/zzz", "GET /api/yyy"},
				},
			},
		}
		_ = ctx.PrincipalManager.Create(context.Background(), root)

		ctx.SecureRouter.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"appName": appName})
		})
	})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}

/*
func fullSample(_ *cobra.Command, args []string) {
	appName := "go-feather-api-sample"
	var authenticationDelegate feather_security.AuthenticationDelegate
	var authorizationDelegate feather_security.AuthorizationDelegate
	builder := &boot.BeanBuilder{
		PasswordEncoder: func() feather_security.PasswordEncoder {
			return feather_security.NewBcryptPasswordEncoder()
		},
		PasswordGenerator: func() feather_security.PasswordGenerator {
			return feather_security.NewDefaultPasswordGenerator()
		},
		PrincipalManager: func(passwordManager feather_security.PasswordManager) feather_security.PrincipalManager {

			principalManager := feather_security.NewInMemoryPrincipalManager(passwordManager)
			authenticationDelegate, authorizationDelegate = principalManager, principalManager

			root := &feather_security.Principal{
				Username:           feather_commons_util.ValueToPtr("root"),
				Password:           feather_commons_util.ValueToPtr("RaveN123qweasd*+"),
				AccountNonExpired:  feather_commons_util.ValueToPtr(true),
				AccountNonLocked:   feather_commons_util.ValueToPtr(true),
				PasswordNonExpired: feather_commons_util.ValueToPtr(true),
				Enabled:            feather_commons_util.ValueToPtr(true),
				SignUpDone:         feather_commons_util.ValueToPtr(true),
				Authorities: feather_commons_util.ValueToPtr([]feather_security.GrantedAuthority{
					{
						Role: feather_commons_util.ValueToPtr("rol01"),
					},
					{
						Role: feather_commons_util.ValueToPtr("rol02"),
					},
					{
						Role: feather_commons_util.ValueToPtr("rol03"),
					},
				}),
			}
			_ = principalManager.Create(context.Background(), root)
			return principalManager
		},
		TokenManager: func(secret string) feather_security.TokenManager {
			return feather_security.NewJwtTokenManager([]byte(secret), feather_security.WithIssuer(appName))
		},
		AuthenticationDelegate: func() feather_security.AuthenticationDelegate {
			return authenticationDelegate
		},
		AuthenticationService: func(tokenManager feather_security.TokenManager, authenticationDelegate feather_security.AuthenticationDelegate) feather_security.AuthenticationService {
			return feather_security.NewDefaultAuthenticationService(tokenManager, authenticationDelegate)
		},
		AuthorizationDelegate: func() feather_security.AuthorizationDelegate {
			return authorizationDelegate
		},
		AuthorizationService: func(tokenManager feather_security.TokenManager, authorizationDelegate feather_security.AuthorizationDelegate) feather_security.AuthorizationService {
			return feather_security.NewDefaultAuthorizationService(tokenManager, authorizationDelegate)
		},
		AuthenticationEndpoint: func(authenticationService feather_security.AuthenticationService) feather_security.AuthenticationEndpoint {
			return feather_security.NewDefaultAuthenticationEndpoint(authenticationService)
		},
		AuthorizationFilter: func(authorizationService feather_security.AuthorizationService) feather_security.AuthorizationFilter {
			return feather_security.NewDefaultAuthorizationFilter(authorizationService)
		},
	}
	err := boot.Init(appName, args, builder, func(ctx boot.ApplicationContext) {

	})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}
*/
