package postgres

import (
	"StellaRP/modules/config"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestUsePostgres(t *testing.T) {
	g := NewGomegaWithT(t)
	app := fxtest.New(t,
		config.UseConfigModule("./../../.test.env"),
		fx.Provide(UsePostgres),
		fx.Invoke(func(conn IConnection) {
			g.Expect(conn).NotTo(BeNil())
		}),
	)

	defer app.RequireStop()
	app.RequireStop()

}
