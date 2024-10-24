package presentation

import (
	"StellaRP/modules/identity/presentation/handlers"
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
