package stores

import (
	"StellaRP/modules/identity/infrastructure/persistence/interfaces"
	"StellaRP/modules/identity/infrastructure/persistence/stores/postgres"
	"StellaRP/modules/identity/infrastructure/persistence/stores/postgres/repositories"
	"go.uber.org/fx"
)

type (
	Stores struct {
		RealmStore interfaces.IRealmStore
	}
)

func UseStores() fx.Option {

	return fx.Options(
		fx.Provide(postgres.UsePostgres),
		fx.Provide(repositories.UsePgRealmStore),
	)
}
