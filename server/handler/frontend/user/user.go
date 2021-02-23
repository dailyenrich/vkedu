package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"vke/pkg/param"
	"vke/pkg/response"
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
		},
	}
}

func UserRegisterHandler(c *gin.Context)  {
	params := new(UserRegisterParams)
	err := c.ShouldBind(params)
	err = param.Validate(err.(validator.ValidationErrors), params)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(response.Err, response.Message[response.Err], err.Error()))
	}
}