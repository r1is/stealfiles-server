package main

import (
	"stealfiles-server/router"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := router.Router()
	r.Run(":9000")

}
