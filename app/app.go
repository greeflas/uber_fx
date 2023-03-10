package app

import (
	"github.com/greeflas/uber_fx/handler"
	"github.com/greeflas/uber_fx/server"
	"github.com/greeflas/uber_fx/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Run() {
	fx.New(getAllOptions()...).Run()
}

func getAllOptions() []fx.Option {
	dependencies := getDependencies()

	options := append(dependencies, []fx.Option{
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	}...)

	return options
}

func getDependencies() []fx.Option {
	return []fx.Option{
		fx.Provide(zap.NewExample),
		fx.Decorate(func(hello service.Hello) service.Hello {
			return service.NewHelloJSONDecorator(hello)
		}),
		server.Module,
		handler.Module,
		service.Module,
	}
}
