package models

import (
	"github.com/KouT127/attendance-management/database"
	"time"
)

type AttendanceTime struct {
	Id               int64
	Remark           string
	AttendanceId     int64
	AttendanceKindId int64
	IsModified       bool
	PushedAt         time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AttendanceTime) TableName() string {
	return database.AttendanceTimeTable
}

type Attendance struct {
	Id        int64
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time

	ClockedIn  *AttendanceTime `xorm:"-"`
	ClockedOut *AttendanceTime `xorm:"-"`
}

func (Attendance) TableName() string {
	return database.AttendanceTable
}

func (a *Attendance) isClockedOut() bool {
	return a.ClockedOut != nil
}

func (a *Attendance) nextKind() AttendanceKind {
	if !a.isClockedOut() {
		return AttendanceKindClockIn
	} else {
		return AttendanceKindClockOut
	}
}

func (a *Attendance) setTimes(cit *AttendanceTime, cot *AttendanceTime) Attendance {
	return Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  cit,
		ClockedOut: cot,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
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
	if d.ClockedInTime.Id != 0 {
		in = d.ClockedInTime
	}
	if d.ClockedOutTime.Id != 0 {
		out = d.ClockedOutTime
	}

	attendance := &Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  in,
		ClockedOut: out,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
	return attendance
}

type AttendanceKind int

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

func fetchAttendancesCount(eng Engine, a *Attendance) (int64, error) {
	attendance := &Attendance{}
	attendance.Id = a.Id
	return eng.Count(attendance)
}

func fetchLatestAttendance(eng Engine, userId string) (*Attendance, error) {
	attendance := &AttendanceDetail{}
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, time.Local)

	has, err := eng.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
		Table(database.AttendanceTable).
		Join("left", "attendances_time clocked_in_time", "attendances.id = clocked_in_time.attendance_id").
		Join("left", "attendances_time clocked_out_time", "attendances.id = clocked_out_time.attendance_id").
		Where("attendances.user_id = ?", userId).
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

func fetchAttendances(eng Engine, a *Attendance, p *Paginator) ([]*Attendance, error) {
	attendances := make([]*Attendance, 0)
	page := p.CalculatePage()
	err := eng.
		Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
		Table(database.AttendanceTable).
		Join("left", "attendances_time clocked_in_time", "attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1").
		Join("left", "attendances_time clocked_out_time", "attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2").
		Where("attendances.user_id = ?", a.UserId).
		And("clocked_in_time.is_modified = false").
		And("clocked_out_time.is_modified = false").
		Limit(int(p.Limit), int(page)).
		OrderBy("-attendances.id").
		Iterate(&AttendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*AttendanceDetail)
			a := d.toAttendance()
			attendances = append(attendances, a)
			return nil
		})
	return attendances, err
}

func updateOldAttendanceTime(eng Engine, attendanceTime *AttendanceTime) error {
	query := &AttendanceTime{
		AttendanceId:     attendanceTime.AttendanceId,
		AttendanceKindId: attendanceTime.AttendanceKindId,
		IsModified:       false,
	}
	_, err := eng.UseBool("is_modified").
		Update(&AttendanceTime{IsModified: true, UpdatedAt: time.Now()}, query)
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

func FetchAttendancesCount(a *Attendance) (int64, error) {
	return fetchAttendancesCount(engine, a)
}

func FetchLatestAttendance(userId string) (*Attendance, error) {
	return fetchLatestAttendance(engine, userId)
}

func FetchAttendances(a *Attendance, p *Paginator) ([]*Attendance, error) {
	return fetchAttendances(engine, a, p)
}

func CreateOrUpdateAttendance(attendance *Attendance, attendanceTime *AttendanceTime, userId string) error {
	sess := engine.NewSession()
	attendance, err := fetchLatestAttendance(sess, userId)
	if err != nil {
		return err
	}

	if attendance == nil {
		attendance = new(Attendance)
		attendance.UserId = userId
		if err := createAttendance(sess, attendance); err != nil {
			return err
		}
	}

	attendanceTime.AttendanceId = attendance.Id
	attendanceTime.AttendanceKindId = int64(attendance.nextKind())
	if err = updateOldAttendanceTime(sess, attendanceTime); err != nil {
		return err
	}

	if err := createAttendanceTime(sess, attendanceTime); err != nil {
		return err
	}
	return nil
}
