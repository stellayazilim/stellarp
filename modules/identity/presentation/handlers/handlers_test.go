package handlers

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRealmHandler_ShouldRegisterHandlers(t *testing.T) {
	g := NewGomegaWithT(t)
	app := gin.New()

	UseHandlers(app)

	h := app.Routes()

	g.Expect(len(h)).Should(BeNumerically(">=", 2))

}
