package handler

import (
	"github.com/greeflas/uber_fx/server"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	fx.Provide(
		asRoute(NewEchoHandler),
		asRoute(NewHelloHandler),
	),
)

func asRoute(r any) any {
	return fx.Annotate(
		r,
		fx.As(new(server.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
