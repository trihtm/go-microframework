package main

import (
	"flag"
	"go-microframework/internal"
	"go-microframework/internal/provider"
	"go.uber.org/fx"
)

func main() {
	flag.String("config-path", "", "config path")

	app := fx.New(
		fx.Provide(
			provider.NewConfig,
			provider.NewEcho,
		),

		fx.Invoke(
			provider.InitConfig,
			internal.RegisterRoutes,
		),
	)

	app.Run()
}