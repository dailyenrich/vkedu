package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vke/pkg/config"
	"vke/pkg/log"
	"vke/pkg/model"
	"vke/router"
	"vke/router/middleware"
)

var (
	Engine = gin.Default()
	GormDB *gorm.DB
)

func init()  {
	log.Init("")
	router.Init()
	middleware.Init()
	GormDB = model.InitGormDB()
}

func Run() {
	Engine.Run(fmt.Sprintf(":%s", config.Get().Port))
}
