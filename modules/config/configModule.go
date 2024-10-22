package config

import "go.uber.org/fx"

func UseConfigModule() fx.Option {
	loadEnv()

	module := fx.Module("ConfigModule", fx.Provide(
		func() *Config {
			return newConfig()
		},
	))

	return module
}
