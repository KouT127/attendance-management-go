package models

import (
	"errors"
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/modules/timeutil"
	"github.com/KouT127/attendance-management/modules/timezone"
	"github.com/Songmu/flextime"
	"time"
)

type AttendanceTime struct {
	ID               int64
	Remark           string
	AttendanceID     int64
	AttendanceKindID uint8
	IsModified       bool
	PushedAt         time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (AttendanceTime) TableName() string {
	return database.AttendanceTimeTable
}

type Attendance struct {
	ID        int64
	UserID    string
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

func (d AttendanceDetail) toAttendance() *Attendance {
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

func fetchAttendancesCount(eng Engine, userID string) (int64, error) {
	attendance := &Attendance{}
	attendance.UserID = userID
	return eng.Count(attendance)
}

func fetchLatestAttendance(eng Engine, userID string) (*Attendance, error) {
	attendance := &AttendanceDetail{}
	now := flextime.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, timezone.JSTLocation())
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, timezone.JSTLocation())

	has, err := eng.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
		Table(database.AttendanceTable).
		Join("left outer",
			"attendances_time clocked_in_time",
			"attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1 and clocked_in_time.is_modified = false").
		Join("left outer",
			"attendances_time clocked_out_time",
			"attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2 and clocked_out_time.is_modified = false").
		Where("attendances.user_id = ?", userID).
		Where("attendances.created_at Between ? and ? ", start, end).
		Limit(1).
		OrderBy("-attendances.id").
		Get(attendance)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return attendance.toAttendance(), nil
}

func fetchAttendances(eng Engine, opt *AttendanceSearchOption) ([]*Attendance, error) {
	attendances := make([]*Attendance, 0)
	sess := eng.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
		Table(database.AttendanceTable).
		Join("left outer",
			"attendances_time clocked_in_time",
			"attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1 and clocked_in_time.is_modified = false").
		Join("left outer",
			"attendances_time clocked_out_time",
			"attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2 and clocked_out_time.is_modified = false")

	if opt.Date != nil {
		now := flextime.Now()
		start, end := timeutil.GetMonthRange(now)
		sess = sess.Where("attendances.created_at Between ? and ? ", start, end)
	}

	sess = opt.setPaginatedSession(sess)
	err := sess.
		Where("attendances.user_id = ?", opt.UserID).
		OrderBy("-attendances.id").
		Iterate(&AttendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*AttendanceDetail)
			a := d.toAttendance()
			attendances = append(attendances, a)
			return nil
		})
	return attendances, err
}

func updateOldAttendanceTime(eng Engine, id int64, kindID uint8) error {
	query := &AttendanceTime{
		AttendanceID:     id,
		AttendanceKindID: kindID,
		IsModified:       false,
	}

	_, err := eng.UseBool("is_modified").
		Update(&AttendanceTime{IsModified: true, UpdatedAt: flextime.Now()}, query)

	if err != nil {
		return err
	}
	return nil
}

func createAttendance(eng Engine, a *Attendance) error {
	if _, err := eng.Insert(a); err != nil {
		return err
	}
	return nil
}

func createAttendanceTime(eng Engine, t *AttendanceTime) error {
	if _, err := eng.Insert(t); err != nil {
		return err
	}
	return nil
}

func FetchAttendancesCount(userID string) (int64, error) {
	return fetchAttendancesCount(engine, userID)
}

func FetchLatestAttendance(userID string) (*Attendance, error) {
	return fetchLatestAttendance(engine, userID)
}

func FetchAttendances(opt *AttendanceSearchOption) ([]*Attendance, error) {
	return fetchAttendances(engine, opt)
}

func CreateOrUpdateAttendance(attendanceTime *AttendanceTime, userID string) (*Attendance, error) {
	sess := engine.NewSession()
	defer sess.Close()

	if userID == "" {
		return nil, errors.New("userID is empty")
	}

	if attendanceTime == nil {
		return nil, errors.New("attendance time is empty")
	}

	attendance, err := fetchLatestAttendance(sess, userID)
	if err != nil {
		return nil, err
	}

	if attendance == nil {
		attendance = new(Attendance)
		attendance.UserID = userID
		attendance.ClockedIn = attendanceTime
		attendance.CreatedAt = flextime.Now()
		attendance.UpdatedAt = flextime.Now()
		if err := createAttendance(sess, attendance); err != nil {
			return nil, err
		}
		attendanceTime.AttendanceID = attendance.ID
		attendanceTime.AttendanceKindID = uint8(AttendanceKindClockIn)
	} else {
		if err := updateOldAttendanceTime(sess, attendance.ID, uint8(AttendanceKindClockOut)); err != nil {
			return nil, err
		}
		attendance.ClockedOut = attendanceTime
		attendanceTime.AttendanceID = attendance.ID
		attendanceTime.AttendanceKindID = uint8(AttendanceKindClockOut)
	}

	if err := createAttendanceTime(sess, attendanceTime); err != nil {
		return nil, err
	}

	if err := sess.Commit(); err != nil {
		return nil, err
	}

	return attendance, nil
}
