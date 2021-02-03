package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vke/pkg/config"
)

var (
	Engine = gin.Default()
	lifeCycle = make(map[string]func(), 3)
)

func Init(fn func())  {
	lifeCycle["init"] = fn
}

func Update(fn func())  {
	lifeCycle["update"] = fn
}

func After(fn func())  {
	lifeCycle["after"] = fn
}

func Run() {
	lifeCycle["init"]()
	Engine.Run(fmt.Sprintf(":%s", config.Get().Port))
}
