package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	packHandler "packs-per-order/api/pack"
)

// SetupRouter initializes and returns a configured Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong!",
		})
	})

	router.GET("/pack", packHandler.HandlePackRequest)

	return router
}

// RunServer starts the Gin server
func RunServer() {
	router := SetupRouter()
	router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
