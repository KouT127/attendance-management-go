package models

import (
	"github.com/KouT127/attendance-management/database"
	"time"
)

type AttendanceTime struct {
	ID               int64 `xorm:"id"`
	Remark           string
	AttendanceID     int64 `xorm:"attendance_id"`
	AttendanceKindID uint8 `xorm:"attendance_kind_id"`
	IsModified       bool
	PushedAt         time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (AttendanceTime) TableName() string {
	return database.AttendanceTimeTable
}

type Attendance struct {
	ID        int64  `xorm:"id"`
	UserID    string `xorm:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ClockedIn  *AttendanceTime `xorm:"-"`
	ClockedOut *AttendanceTime `xorm:"-"`
}

func (Attendance) TableName() string {
	return database.AttendanceTable
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
		ClockedIn:  in,
		ClockedOut: out,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
	return attendance
}

type AttendanceKind uint8

const (
	AttendanceKindNone AttendanceKind = iota
	AttendanceKindClockIn
	AttendanceKindClockOut
)

func (k AttendanceKind) String() string {
	switch k {
	case AttendanceKindClockIn:
		return "出勤"
	case AttendanceKindClockOut:
		return "退勤"
	}
	return "不明"
}
