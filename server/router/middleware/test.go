package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vke/pkg/app"
	"vke/pkg/log"
)

func test()  {
	app.Engine.Use(func(c *gin.Context) {
		fmt.Println("okkkkkkkkkkkkkkkkkkkkkkkkk")
		log.Info(c.Request.URL)
		c.Next()
	})
}
