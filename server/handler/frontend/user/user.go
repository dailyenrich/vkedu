package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"vke/model"
	"vke/pkg/param"
	"vke/pkg/e"
	"vke/service/user"
)

type UserRegisterParams struct {
	Username        string `json:"username" binding:"required,min=6,max=20" form:"username"`
	Password        string `json:"password" binding:"required,min=6,max=20" form:"password"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6,max=20" form:"confirm_password"`
	Code 			string `json:"code" binding:"required" form:"code"`
}

func (u UserRegisterParams) Rule() param.R {
	return param.R{
		"Username": {
			"required": "用户名不能为空",
			"min": "用户名最小6位",
			"max": "用户名最大20位",
		},
	}
}

func RegisterHandler(c *gin.Context) {
	params := new(UserRegisterParams)
	err := c.ShouldBind(params)
	if err != nil {
		err = param.Validate(err.(validator.ValidationErrors), params)
	}
	if err != nil {
		c.JSON(http.StatusOK, e.Resp(e.Err, e.Message[e.Err], err.Error()))
		return
	}

	dto := new(model.UserDto)
	dto.Username = params.Username
	dto.Password = params.Password

	service := user.NewUserService()
	err = service.Register(dto)
	if err != nil {
		c.JSON(http.StatusOK, e.Resp(e.Err, e.Message[e.Err], err.Error()))
		return
	}

	c.JSON(http.StatusOK, e.Resp(e.Ok, e.Message[e.Ok], "注册成功"))
}