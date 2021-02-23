package router

import (
	"vke/handler/frontend/user"
	"vke/pkg/app"
)

func userApi()  {
	app.Engine.GET("/user", user.UserRegisterHandler)
}
