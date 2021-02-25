package user

import (
	"vke/model"
	user2 "vke/repository/user"
)

type UserService struct {
	
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) Register(user *model.UserDto) error {
	return user2.InsertUser(user)
}

func (u UserService) Login(member *model.UserDto) (model.User, error) {
	panic("implement me")
}

