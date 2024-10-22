package identity

import (
	"StellaRP/modules/identity/infrastructure"
	"StellaRP/modules/identity/presentation"
	"go.uber.org/fx"
)

func UseIdentityModule() fx.Option {

	m := fx.Module("Identity",
		presentation.UsePresentation(),
		infrastructure.UseInfrastructure(),
	)

	return m
}
