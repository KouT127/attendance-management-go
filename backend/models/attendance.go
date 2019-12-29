package models

import "time"

type Attendance struct {
	Id         int64
	UserId     string
	ClockedIn  AttendanceTime
	ClockedOut AttendanceTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type AttendanceTime struct {
	Id        int64
	Remark    string
	PushedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
