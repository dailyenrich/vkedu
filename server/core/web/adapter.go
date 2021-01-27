package web

import "github.com/gin-gonic/gin"

func NewWebServer() *gin.Engine {
	engine := gin.New()
	return engine
}