package user

import (
	"vke/model"
	"vke/pkg/app"
)

func InsertUser(dto *model.UserDto) error {
	return app.GormDB.Exec("insert into user (username, password, created_at, updated_at) value (?,?,?,?)", dto.Username,dto.Password).Error
}
