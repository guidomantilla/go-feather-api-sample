package serve

import (
	"syscall"

	"github.com/gin-gonic/gin"
	feather_commons_config "github.com/guidomantilla/go-feather-commons/pkg/config"
	"github.com/qmdx00/lifecycle"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-api-sample/internal/config"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	orchestrator := lifecycle.NewApp(
		lifecycle.WithName("go-feather-api-sample"),
		lifecycle.WithVersion("1.0"),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	//orchestrator.Cleanup(config.Init)

	//

	environment := feather_commons_config.Init(args)

	router := gin.Default()

	//

	orchestrator.Attach("GinServer", config.InitGinServer(environment, router))
	if err := orchestrator.Run(); err != nil {
		zap.L().Fatal(err.Error())
	}

	_ = logger.Sync()
}
