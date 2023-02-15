package app

import (
	"io"
	"os"

	"github.com/greeflas/uber_fx/handler"
	"github.com/greeflas/uber_fx/server"
	"github.com/greeflas/uber_fx/service"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Run() {
	fx.New(getAllOptions()...).Run()
}

func getAllOptions() []fx.Option {
	dependencies := getDependencies()

	options := append(dependencies, []fx.Option{
		fx.WithLogger(func(log *logrus.Logger) fxevent.Logger {
			return NewLogrusFxEventLogger(log)
		}),
	}...)

	return options
}

func getDependencies() []fx.Option {
	return []fx.Option{
		fx.Provide(newLogger),
		fx.Decorate(func(hello service.Hello) service.Hello {
			return service.NewHelloJSONDecorator(hello)
		}),
		server.Module,
		handler.Module,
		service.Module,
	}
}

func newLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "02-01-2006 15:04:05"})
	//log.SetLevel(logrus.WarnLevel)

	log.SetOutput(io.Discard) // Send all logs to nowhere by default

	log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})
	log.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})

	return log
}
