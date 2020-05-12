package models

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	ImageURL  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (User) TableName() string {
	return "users"
}

type GetOrCreateUserParams struct {
	UserID string
}

type GetOrCreateUserResults struct {
	User *User
}
