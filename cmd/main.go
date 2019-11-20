package main

import (
	"flag"
	"go-microframework/internal"
	"go-microframework/internal/provider"
	"go-microframework/pkg/cache"
	"go-microframework/pkg/database"
	"go-microframework/pkg/filesystem"
	"go-microframework/pkg/log"
	"go-microframework/pkg/mail"
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

		cache.NewService(),
		database.NewService(),
		filesystem.NewService(),
		mail.NewService(),
		log.NewService(),
		queue.NewService(),
	)

	app.Run()
}