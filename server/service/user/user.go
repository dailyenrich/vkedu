package user

import (
	"vke/model"
)

type UserService struct {
	
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) Register(user *model.UserDto) error {
	return nil
}

func (u UserService) Login(member *model.UserDto) (model.User, error) {
	panic("implement me")
}

