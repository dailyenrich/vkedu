package middleware

import (
	"github.com/gin-gonic/gin"
	"vke/pkg/app"
	"vke/pkg/log"
)

func test()  {
	app.Engine.Use(func(c *gin.Context) {
		log.Logger().Infoln(c.Request.URL)
		c.Next()
	})
}
