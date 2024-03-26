package router

import (
	"stealfiles-server/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/", controller.Msgdata{}.GetMsg)

	return r
}
