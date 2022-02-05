package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	service *TestService
}

func (controller *TestController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": controller.service.ping(),
	})
}

func NewTestController(service *TestService) *TestController {
	return &TestController{
		service: service,
	}
}
