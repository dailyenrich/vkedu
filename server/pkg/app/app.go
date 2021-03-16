package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vke/pkg/config"
)

var (
	Engine = gin.Default()
	GormDB *gorm.DB
)

func Init(fn func())  {
	fn()
}

func Run() {
	Engine.Run(fmt.Sprintf(":%s", config.Get().Port))
}
