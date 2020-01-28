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

type attendanceOption interface {
	apply(*Attendance)
}

type clockedIn struct {
	t *AttendanceTime
}

type clockedOut struct {
	t *AttendanceTime
}

func (c clockedIn) apply(a *Attendance) {
	a.ClockedIn = c.t
}

func attendanceWithClockedIn(t *AttendanceTime) clockedIn {
	return clockedIn{
		t: t,
	}
}

func (c clockedOut) apply(a *Attendance) {
	a.ClockedOut = c.t
}

func attendanceWithClockedOut(t *AttendanceTime) clockedOut {
	return clockedOut{
		t: t,
	}
}

func (a *Attendance) setValues(opts ...attendanceOption) {
	for _, opt := range opts {
		opt.apply(a)
	}
}

func (a *Attendance) ClockIn(time *AttendanceTime) {
	a.setValues(attendanceWithClockedIn(time))
}

func (a *Attendance) ClockOut(time *AttendanceTime) {
	a.setValues(attendanceWithClockedOut(time))
}

func (a *Attendance) IsClockedOut() bool {
	return a.ClockedOut != nil
}
