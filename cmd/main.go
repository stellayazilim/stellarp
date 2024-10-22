package main

import (
	"StellaRP/modules/config"
	"StellaRP/modules/identity"
	"go.uber.org/fx"
)

func main() {
	Bootstrap()
}

func Bootstrap() {
	fx.New(
		config.UseConfigModule(),
		identity.UseIdentityModule(),
	).Run()
}
