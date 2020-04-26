package models

import (
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	ImageURL  string `xorm:"image_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

type GetOrCreateUserParams struct {
	UserId string
}

type GetOrCreateUserResults struct {
	User             *User
	LatestAttendance *Attendance
}
