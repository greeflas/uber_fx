package server

import (
	"net"

	"go.uber.org/fx"
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
