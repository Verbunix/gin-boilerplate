package test

import (
	"docsign/utils/di"

	"github.com/gin-gonic/gin"
)

func Module() di.Module {
	return di.Module{
		Imports: di.Modules{},
		Controllers: di.Deps{
			NewTestController,
		},
		Provides: di.Deps{
			NewTestService,
		},
		Routing: func(engine *gin.Engine, controller *TestController) {
			engine.GET("/", controller.Ping)
		},
	}
}
