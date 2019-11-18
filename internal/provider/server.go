package provider

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"go-microframework/internal/config"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, cf config.Configuration) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				host := cf.Server.Host
				port := cf.Server.Port

				e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port))) // curl localhost/ping
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			return e.Close()
		},
	})

	return e
}