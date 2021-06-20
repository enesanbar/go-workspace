package services

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewScheduler,
		NewPusherClient,
		NewDispatcherConfig,
		NewTester,
		NewMonitoring,
		NewRenderer,
	),
)
