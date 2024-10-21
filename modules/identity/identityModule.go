package identity

import (
	"go.uber.org/fx"
)

type Identity struct {
}

func InitIdentityModule() fx.Option {

	m := fx.Module("Identity",
		fx.Provide(func() *Identity { return new(Identity) }),
	)

	return m
}
