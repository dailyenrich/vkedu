package user

import (
	"time"
	"vke/model"
	"vke/pkg/hash"
	user2 "vke/repository/user"
)

type UserService struct {
	
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) Register(user *model.UserDto) error {
	hash := hash.NewHash()
	bytes, err := hash.Make([]byte(user.Password))
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return user2.InsertUser(user)
}

func (u UserService) Login(member *model.UserDto) (model.User, error) {
	panic("implement me")
}

