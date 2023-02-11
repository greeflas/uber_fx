package server

import (
	"go.uber.org/fx"
	"net"
)

var Module = fx.Module("server",
	fx.Provide(NewHTTPServer),
	fx.Provide(
		fx.Private,
		func() (net.Listener, error) {
			return NewTCPListener(":8080")
		},
	),
	fx.Provide(
		fx.Private,
		fx.Annotate(
			NewServerMux,
			fx.ParamTags(`group:"routes"`),
		),
	),
	fx.Invoke(Run),
)
