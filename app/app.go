package app

import (
	"github.com/greeflas/uber_fx/handler"
	"github.com/greeflas/uber_fx/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Run() {
	dependencies := getDependencies()

	options := append(dependencies, []fx.Option{
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	}...)

	fx.New(options...).Run()
}

func getDependencies() []fx.Option {
	return []fx.Option{
		fx.Provide(zap.NewExample),
		server.Module,
		handler.Module,
	}
}
