package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type file struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

var files = make([]file, 0)

func create(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		panic(err)
	}

	log.Println("File:", file.Filename)
	src, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer src.Close()

	context.String(200, "Yay!")
}

func get(context *gin.Context) {

}

// Mount will mount a set of routes into the given router
func Mount(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	{
		api.POST("/files", createFileHandler)
		api.GET("/files/:id", indexHandler)
		// api.PUT("/files/:id")
		// api.DELETE("/files/:id")
	}
	return router
}
