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
	lifeCycle = make(map[string]func(), 3)
)

func Init(fn func())  {
	fn()
}

func Update(fn func())  {
	lifeCycle["update"] = fn
}

func After(fn func())  {
	lifeCycle["after"] = fn
}

func Run() {
	Engine.Run(fmt.Sprintf(":%s", config.Get().Port))
}
