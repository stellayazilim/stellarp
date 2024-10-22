package presentation

import (
	"StellaRP/modules/identity/presentation/handlers"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

func UseGin() *gin.Engine {
	router := gin.Default()

	return router
}

func StartServer(lc fx.Lifecycle, router *gin.Engine) error {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", s.Addr)
			gr, ctx := errgroup.WithContext(context.Background())
			if err != nil {
				return err
			}

			gr.Go(func() error {
				return s.Serve(ln)
			})

			return ctx.Err()

		},

		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})
	return nil

}

func UsePresentation() fx.Option {
	return fx.Options(
		fx.Provide(UseGin),
		fx.Invoke(handlers.UseHandlers),
		fx.Invoke(StartServer),
	)
}
