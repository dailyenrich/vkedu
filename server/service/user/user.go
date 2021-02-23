package user

import (
	"vke/model"
)

type UserService struct {
	
}

func (u UserService) Register(user model.UserDto) error {
	panic("implement me")
}

func (u UserService) Login(member model.UserDto) (model.User, error) {
	panic("implement me")
}

