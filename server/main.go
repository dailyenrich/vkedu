package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vke/pkg/config"
	"vke/pkg/log"
)

func main()  {
	r := gin.Default()
	log.Init("")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
		log.Info("ojbk")
	})
	r.Run(fmt.Sprintf(":%s", config.Get().Port))
}
