package identity

import (
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestInitIdentityModule(t *testing.T) {
	g := NewGomegaWithT(t)
	app := fxtest.New(t,
		InitIdentityModule(),
		fx.Invoke(func(im *Identity) {
			g.Expect(im).NotTo(BeNil())
		}),
	)

	defer app.RequireStop()
	app.RequireStart()
}
