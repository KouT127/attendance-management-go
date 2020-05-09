package models

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
