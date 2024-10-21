package config

import "go.uber.org/fx"

func InitConfigModule() fx.Option {
	loadEnv()

	module := fx.Module("ConfigModule", fx.Provide(
		func() *Config {
			return newConfig()
		},
	))

	return module
}
