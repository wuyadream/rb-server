package main

import (

	"github.com/gin-gonic/gin"

	"github.com/wuya.dream/rb-server/handlers"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/rubbish/v1/type", handlers.GetType)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:6789
	r.Run(":6789")
}
