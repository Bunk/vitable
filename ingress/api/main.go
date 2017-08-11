package api

import (
	"github.com/bunk/vitable/ingress/config"
	"github.com/gin-gonic/gin"
)

// Start will create an HTTP server and start listening for requests
func Start(c config.Configuration) error {
	if c.Mode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := Router()
	return router.Run(":" + c.Port)
}
