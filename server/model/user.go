package model

import "time"

type User struct {
	ID uint32
	Username string
	Password string
	CreatedUser string
	CreatedAt time.Duration
	UpdatedAt time.Duration
	DeletedAt time.Duration
}

type UserDto struct {
	ID uint32
	Username string
	Password string
	CreatedUser string
	CreatedAt time.Duration
	UpdatedAt time.Duration
	DeletedAt time.Duration
}

type UserExt struct {
	ID uint32
	Username string
	Nickname string
	Score float32
	LearnTime uint32
	Description string
	LastLoginedAt time.Duration
	LastLoginedIp string
	Revision uint32
	CreatedUser string
	CreatedAt time.Duration
	UpdatedAt time.Duration
	DeletedAt time.Duration
}
type UserExtDto struct {
	ID uint32
	Username string
	Nickname string
	Score float32
	LearnTime uint32
	Description string
	LastLoginedAt time.Duration
	LastLoginedIp string
	Revision uint32
	CreatedUser string
	CreatedAt time.Duration
	UpdatedAt time.Duration
	DeletedAt time.Duration
}
