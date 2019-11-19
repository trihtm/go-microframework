package main

import (
	"flag"
	"go-microframework/internal"
	"go-microframework/internal/provider"
	"go-microframework/pkg/database"
	"go-microframework/pkg/queue"
	"go.uber.org/fx"
)

func main() {
	flag.String("config-path", "", "config path")

	app := fx.New(
		fx.Provide(
			provider.InitializeConfig,
			provider.NewHttpServer,
		),

		fx.Invoke(
			internal.RegisterRoutes,
		),

		database.NewService(),
		queue.NewService(),
	)

	app.Run()
}