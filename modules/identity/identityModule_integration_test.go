package identity

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestInitIdentityModule(t *testing.T) {
	g := NewGomegaWithT(t)

	app := UseIdentityModule()

	g.Expect(app).NotTo(BeNil())

}
