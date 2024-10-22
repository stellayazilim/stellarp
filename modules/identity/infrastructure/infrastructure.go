package infrastructure

import (
	"StellaRP/modules/identity/infrastructure/persistence/stores"
	"go.uber.org/fx"
)

func UseInfrastructure() fx.Option {
	return fx.Options(
		stores.UseStores(),
	)
}
