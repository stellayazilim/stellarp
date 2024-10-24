package application

import (
	"StellaRP/modules/identity/application/features"
	"go.uber.org/fx"
)

func UseApplication() fx.Option {

	return fx.Options(
		features.UseFeatures(),
	)
}
