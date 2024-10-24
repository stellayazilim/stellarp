package main

import (
	"StellaRP/modules/config"
	"StellaRP/modules/identity"
	_ "github.com/flashlabs/rootpath"
	"go.uber.org/fx"
)

func main() {
	Bootstrap()
}

func Bootstrap() {
	fx.New(
		config.UseConfigModule(".env"),
		identity.UseIdentityModule(),
	).Run()
}
