package boot

import (
	"net/http"
	"strings"
	"syscall"

	"github.com/guidomantilla/go-feather-web/pkg/server"
	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

func Init(appName string, version string, args []string, builder *BeanBuilder, fn func(ctx ApplicationContext)) error {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	if appName == "" {
		zap.L().Fatal("starting up - error setting up the application: appName is empty")
	}

	if args == nil {
		zap.L().Fatal("starting up - error setting up the application: args is nil")
	}

	if builder == nil {
		zap.L().Fatal("starting up - error setting up the application: builder is nil")
	}

	if fn == nil {
		zap.L().Fatal("starting up - error setting up the application: fn is nil")
	}

	app := lifecycle.NewApp(
		lifecycle.WithName(appName),
		lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)
	app.Cleanup() //TODO

	ctx := NewApplicationContext(strings.Join([]string{appName, version}, " - "), args, builder)

	fn(*ctx)

	hostAddress := ctx.Environment.GetValueOrDefault(HostPort, EnvVarDefaultValuesMap[HostPort]).AsString()
	httpServer := &http.Server{
		Addr:              hostAddress,
		Handler:           ctx.Router,
		ReadHeaderTimeout: 60000,
	}

	app.Attach("GinServer", server.BuildHttpServer(httpServer))
	return app.Run()
}
