package model

import "time"

type User struct {
	ID uint32 `gorm:"primaryKey"`
	Username string `gorm:"column:username;unique;type:varchar(32)"`
	Password string `gorm:"column:password;unique;type:varchar(64)"`
	CreatedUser string `gorm:"column:create_user;unique;type:varchar(32)"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;"`
	DeletedAt time.Time
}

func NewUser() *User {
	return &User{}
}

type UserDto struct {
	ID uint32
	Username string
	Password string
	CreatedUser string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserExt struct {
	ID uint32
	Username string
	Nickname string
	Score float32
	LearnTime uint32
	Description string
	LastLoginedAt time.Time
	LastLoginedIp string
	Revision uint32
	CreatedUser string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
type UserExtDto struct {
	ID uint32
	Username string
	Nickname string
	Score float32
	LearnTime uint32
	Description string
	LastLoginedAt time.Time
	LastLoginedIp string
	Revision uint32
	CreatedUser string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
