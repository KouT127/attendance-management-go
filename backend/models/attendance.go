package models

import "time"

type Attendance struct {
	Base
	UserId       string
	ClockedInId  uint
	ClockedOutId uint
}

type AttendanceTime struct {
	Base
	Remark   string
	PushedAt time.Time
}
