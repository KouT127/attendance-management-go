package models

import (
	"time"
)

type AttendanceKind uint8

const (
	AttendanceKindNone AttendanceKind = iota
	AttendanceKindClockIn
	AttendanceKindClockOut
)

type AttendanceTime struct {
	ID               int64
	Remark           string
	AttendanceID     int64
	AttendanceKindID uint8
	IsModified       bool
	PushedAt         time.Time
	CreatedAt        time.Time `xorm:"created"`
	UpdatedAt        time.Time `xorm:"updated"`
}

func (AttendanceTime) TableName() string {
	return "attendances_time"
}

type Attendance struct {
	ID         int64
	UserID     string
	AttendedAt time.Time
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`

	ClockedIn  *AttendanceTime `xorm:"-"`
	ClockedOut *AttendanceTime `xorm:"-"`
}

func (Attendance) TableName() string {
	return "attendances"
}

type Attendances []*Attendance

func (attendances Attendances) ManipulateTotalWorkHours() float64 {
	var total float64
	for _, attendance := range attendances {
		if attendance.ClockedOut == nil {
			continue
		}
		workTime := attendance.ClockedOut.PushedAt.Sub(attendance.ClockedIn.PushedAt)
		total += workTime.Hours()
	}
	return total
}

type AttendanceDetail struct {
	Attendance     `xorm:"extends"`
	ClockedInTime  *AttendanceTime `xorm:"extends"`
	ClockedOutTime *AttendanceTime `xorm:"extends"`
}

func (d AttendanceDetail) ToAttendance() *Attendance {
	var (
		in  *AttendanceTime
		out *AttendanceTime
	)
	a := d.Attendance
	if d.ClockedInTime.ID != 0 {
		in = d.ClockedInTime
	}
	if d.ClockedOutTime.ID != 0 {
		out = d.ClockedOutTime
	}

	attendance := &Attendance{
		ID:         a.ID,
		UserID:     a.UserID,
		AttendedAt: a.AttendedAt,
		ClockedIn:  in,
		ClockedOut: out,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
	return attendance
}

func (k AttendanceKind) String() string {
	switch k {
	case AttendanceKindClockIn:
		return "出勤"
	case AttendanceKindClockOut:
		return "退勤"
	}
	return "不明"
}
