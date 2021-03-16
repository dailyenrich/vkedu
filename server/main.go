package main

import (
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
		app.GormDB = model.InitGormDB()
	})
	app.Run()
}
