package api

import (
	"github.com/gin-gonic/gin"
	status "gopkg.in/appleboy/gin-status-api.v1"
)

func indexHandler(context *gin.Context) {
	context.String(200, "Hello!")
}

func createFileHandler(context *gin.Context) {
	context.String(200, "Yay!")
}

// Router creates a new API http router for handling requests
func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		api.GET("/", indexHandler)
		api.GET("/status", status.StatusHandler)

		api.POST("/files", createFileHandler)
		api.GET("/files/{id}", indexHandler)
		// api.PUT("/files/{id}")
		// api.DELETE("/files/{id}")
	}

	return router
}
