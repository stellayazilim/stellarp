package features

import (
	realmfeature "StellaRP/modules/identity/application/features/realm"
	"go.uber.org/fx"
)

func UseFeatures() fx.Option {

	return fx.Options(
		fx.Invoke(realmfeature.RegisterCreateRealmHandler),
	)
}
