package user

import "vke/model"

type UserContract interface {
	Login(user *model.UserDto) (model.User, error)
	Register(user *model.UserDto) error
}