package bus

import "go.uber.org/fx"

func NewService() fx.Option {
	return fx.Options(
		fx.Provide(),
	)
}
