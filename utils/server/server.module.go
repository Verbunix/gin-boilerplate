package server

import (
	"context"
	"docsign/utils/di"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Module() di.Module {
	return di.Module{
		Provides: di.Deps{
			func(lc fx.Lifecycle) *gin.Engine {
				var r *gin.Engine

				if os.Getenv("GIN_MODE") == "debug" {
					gin.SetMode(gin.DebugMode)
					r = gin.Default()
				} else {
					gin.SetMode(gin.ReleaseMode)
					r = gin.New()
				}

				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Default().Println("Server has been started at port :8080")
						r.Run(":8080")
						return nil
					},
				})

				return r
			},
		},
	}
}
