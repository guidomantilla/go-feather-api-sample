package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
	feather_security "github.com/guidomantilla/go-feather-security/pkg/security"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-api-sample/internal/repositories"
	"github.com/guidomantilla/go-feather-api-sample/internal/service"
	"github.com/guidomantilla/go-feather-api-sample/pkg/boot"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	appName, version := "go-feather-api-sample", "v0.3.0"

	authPrincipalRepository := repositories.NewDefaultAuthPrincipalRepository()

	builder := boot.NewBeanBuilder()
	builder.PrincipalManager = func(appCtx *boot.ApplicationContext) feather_security.PrincipalManager {
		return service.NewDBPrincipalManager(appCtx.TransactionHandler, authPrincipalRepository)
	}

	err := boot.Init(appName, version, args, builder, func(appCtx boot.ApplicationContext) {

		appCtx.SecureRouter.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"appName": appName})
		})
	})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}
