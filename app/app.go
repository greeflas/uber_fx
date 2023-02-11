package app

import (
	"github.com/greeflas/uber_fx/handler"
	"github.com/greeflas/uber_fx/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net"
)

func Run() {
	dependencies := getDependencies()

	options := append(dependencies, []fx.Option{
		fx.Invoke(server.Run),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	}...)

	fx.New(options...).Run()
}

func getDependencies() []fx.Option {
	return []fx.Option{
		fx.Provide(
			server.NewHTTPServer,
			zap.NewExample,
		),
		fx.Provide(func() (net.Listener, error) {
			return server.NewTCPListener(":8080")
		}),
		fx.Provide(fx.Annotate(
			server.NewServerMux,
			fx.ParamTags(`group:"routes"`),
		)),
		fx.Provide(
			asRoute(handler.NewEchoHandler),
			asRoute(handler.NewHelloHandler),
		),
	}
}

func asRoute(r any) any {
	return fx.Annotate(
		r,
		fx.As(new(server.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
