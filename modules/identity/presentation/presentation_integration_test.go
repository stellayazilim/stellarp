package presentation

import (
	"StellaRP/modules/config"
	"StellaRP/modules/identity/presentation/handlers"
	"context"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinIntegration(t *testing.T) {

	g := NewGomegaWithT(t)

	f := fxtest.New(t,
		fx.Provide(UseGin),
		fx.Invoke(handlers.UseHandlers),
		fx.Invoke(func(r *gin.Engine) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/login", nil)
			r.ServeHTTP(w, req)

			g.Expect(w.Code).To(Equal(http.StatusOK))
		}),
	)

	defer f.RequireStop()

	f.RequireStart()

}

func TestStartServer_ShouldStart(t *testing.T) {
	g := NewGomegaWithT(t)
	f := fxtest.New(t,
		fx.Provide(UseGin),
		fx.Invoke(handlers.UseHandlers),
		fx.Invoke(StartServer),
	)

	defer f.RequireStop()

	f.RequireStart()

	g.Expect(f.Err()).NotTo(HaveOccurred())
	g.Expect(f.App.Err()).NotTo(HaveOccurred())

	err := f.App.Err()

	g.Expect(err).To(BeNil())
}

func TestStartServer_ShouldFailOnNonEmptyPort(t *testing.T) {
	g := NewGomegaWithT(t)

	f := fxtest.New(t,
		config.UseConfigModule(),
		fx.Provide(UseGin),
		fx.Invoke(handlers.UseHandlers),
		fx.Invoke(StartServer),
		fx.Invoke(StartServer),
	)

	defer f.RequireStop()

	err := f.Start(context.Background())

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).NotTo(BeEmpty())
}
