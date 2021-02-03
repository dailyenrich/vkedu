package main

import (
	"vke/pkg/app"
	"vke/pkg/log"
	"vke/router"
	"vke/router/middleware"
)

func main()  {
	app.Init(func() {
		log.Init("")
		router.Init()
		middleware.Init()
	})
	app.Run()
}
