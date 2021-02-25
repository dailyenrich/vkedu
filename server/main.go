package main

import (
	"gorm.io/gorm"
	"vke/pkg/app"
	"vke/pkg/log"
	"vke/pkg/model"
	"vke/router"
	"vke/router/middleware"
)

func main()  {
	app.Init(func() {
		log.Init("")
		router.Init()
		middleware.Init()
		model.InitGormDB(func(db *gorm.DB) {
			app.GormDB = db
		})
	})
	app.Run()
}
