package models

import "time"

type AttendanceTime struct {
	Id        int64
	Remark    string
	PushedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Attendance struct {
	Id         int64
	UserId     string
	ClockedIn  *AttendanceTime
	ClockedOut *AttendanceTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (a *Attendance) ClockIn(userId string, time *AttendanceTime) {
	a.UserId = userId
	a.ClockedIn = time
}

func (a *Attendance) ClockOut(time *AttendanceTime) {
	a.ClockedOut = time
}

func (a *Attendance) IsClockedOut() bool {
	return a.ClockedOut != nil
}
