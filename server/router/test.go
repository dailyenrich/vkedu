package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vke/pkg/app"
	"vke/pkg/log"
)

func testApi()  {
	app.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
		log.Logger().Infoln("ping")
	})
}
