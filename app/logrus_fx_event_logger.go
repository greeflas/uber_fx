package app

import (
	"strings"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

type LogrusFxEventLogger struct {
	logger *logrus.Logger
}

func NewLogrusFxEventLogger(logger *logrus.Logger) *LogrusFxEventLogger {
	return &LogrusFxEventLogger{logger: logger}
}

func (l *LogrusFxEventLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logger.WithFields(logrus.Fields{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		}).Info("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err == nil {
			l.logger.WithFields(logrus.Fields{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			})
		} else {
			l.logger.WithFields(logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			}).WithError(e.Err).Error("OnStart hook failed")
		}
	case *fxevent.OnStopExecuting:
		l.logger.WithFields(logrus.Fields{
			"callee": e.FunctionName,
			"caller": e.CallerName,
		}).Info("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err == nil {
			l.logger.WithFields(logrus.Fields{
				"callee":  e.FunctionName,
				"caller":  e.CallerName,
				"runtime": e.Runtime.String(),
			}).Info("OnStop hook executed")
		} else {
			l.logger.WithFields(logrus.Fields{
				"callee": e.FunctionName,
				"caller": e.CallerName,
			}).WithError(e.Err).Error("OnStop hook failed")
		}
	case *fxevent.Supplied:
		fields := logrus.Fields{"type": e.TypeName}
		fields = addModuleIfPresent(fields, e.ModuleName)

		if e.Err == nil {
			l.logger.WithFields(fields).Info("supplied")
		} else {
			l.logger.WithFields(fields).WithError(e.Err).Error("error encountered while applying options")
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			fields := logrus.Fields{
				"constructor": e.ConstructorName,
				"type":        rtype,
			}
			fields = addModuleIfPresent(fields, e.ModuleName)

			if e.Private {
				fields["private"] = true
			}

			l.logger.WithFields(fields).Info("provided")
		}

		if e.Err != nil {
			fields := addModuleIfPresent(logrus.Fields{}, e.ModuleName)

			l.logger.WithFields(fields).WithError(e.Err).Error("error encountered while applying options")
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			fields := logrus.Fields{"type": rtype}
			fields = addModuleIfPresent(fields, e.ModuleName)

			l.logger.WithFields(fields).Info("replaced")
		}

		if e.Err != nil {
			fields := addModuleIfPresent(logrus.Fields{}, e.ModuleName)

			l.logger.WithFields(fields).WithError(e.Err).Error("error encountered while replacing")
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			fields := logrus.Fields{
				"decorator": e.DecoratorName,
				"type":      rtype,
			}
			fields = addModuleIfPresent(fields, e.ModuleName)

			l.logger.WithFields(fields).Info("decorated")
		}

		if e.Err != nil {
			fields := addModuleIfPresent(logrus.Fields{}, e.ModuleName)

			l.logger.WithFields(fields).WithError(e.Err).Error("error encountered while applying options")
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.x
		fields := logrus.Fields{"function": e.FunctionName}
		fields = addModuleIfPresent(fields, e.ModuleName)

		l.logger.WithFields(fields).Info("invoking")
	case *fxevent.Invoked:
		if e.Err == nil {
			return
		}

		fields := logrus.Fields{
			"stack":    e.Trace,
			"function": e.FunctionName,
		}
		fields = addModuleIfPresent(fields, e.ModuleName)

		l.logger.WithFields(fields).WithError(e.Err).Error("invoke failed")
	case *fxevent.Stopping:
		l.logger.WithField("signal", strings.ToUpper(e.Signal.String())).Info("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.logger.WithError(e.Err).Error("stop failed")
		}
	case *fxevent.RollingBack:
		l.logger.WithError(e.StartErr).Error("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.logger.WithError(e.Err).Error("rollback failed")
		}
	case *fxevent.Started:
		if e.Err == nil {
			l.logger.Info("started")
		} else {
			l.logger.WithError(e.Err).Error("start failed")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.logger.WithField("function", e.ConstructorName).Info("initialized custom fxevent.Logger")
		} else {
			l.logger.WithError(e.Err).Error("custom logger initialization failed")
		}
	}
}

func addModuleIfPresent(fields logrus.Fields, moduleName string) logrus.Fields {
	if len(moduleName) != 0 {
		fields["module"] = moduleName
	}

	return fields
}
