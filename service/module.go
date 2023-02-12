package service

import "go.uber.org/fx"

var Module = fx.Module("service",
	fx.Provide(
		fx.Private,
		NewHelloService,
	),
	fx.Provide(func(helloService *HelloService) Hello {
		return helloService
	}),
)
