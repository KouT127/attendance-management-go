package models

import "time"

type Attendance struct {
	Id           uint
	UserId       string
	ClockedIn  AttendanceTime
	ClockedOut AttendanceTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AttendanceTime struct {
	Id        uint
	Remark    string
	PushedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
