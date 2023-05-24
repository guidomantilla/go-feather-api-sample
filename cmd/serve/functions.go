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

	appName, version := "go-feather-api-sample", "v0.3.0"
	builder := boot.NewBeanBuilder()

	err := boot.Init(appName, version, args, builder, func(ctx boot.ApplicationContext) {

		root := &feather_security.Principal{
			Username: feather_commons_util.ValueToPtr("raven"),
			Role:     feather_commons_util.ValueToPtr("ROOT"),
			Password: feather_commons_util.ValueToPtr("Raven123Qweasd*+"),
			Resources: feather_commons_util.ValueToPtr(
				[]string{
					"GET /api/info", "GET /api/xxx", "GET /api/yyy", "GET /api/zzz",
				},
			),
		}

		var err error
		if err = ctx.PrincipalManager.Create(context.Background(), root); err != nil {
			zap.L().Error(err.Error())
		}

		ctx.SecureRouter.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"appName": appName})
		})
	})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}
