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
	PushedAt         time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Attendance struct {
	Id        int64
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time

	ClockedIn  *AttendanceTime `xorm:"-"`
	ClockedOut *AttendanceTime `xorm:"-"`
}

func (a *Attendance) IsClockedOut() bool {
	return a.ClockedOut != nil
}

func (a *Attendance) build(cit *AttendanceTime, cot *AttendanceTime) Attendance {
	return Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  cit.build(),
		ClockedOut: cot.build(),
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

func NewTime(at *AttendanceTime) *AttendanceTime {
	t := new(AttendanceTime)
	t.Id = at.Id
	t.Remark = at.Remark
	t.AttendanceId = at.AttendanceId
	t.AttendanceKindId = at.AttendanceKindId
	t.CreatedAt = at.CreatedAt
	t.UpdatedAt = at.UpdatedAt
	t.PushedAt = at.PushedAt
	return t
}

func (t AttendanceTime) build() *AttendanceTime {
	return &AttendanceTime{
		Id:               t.Id,
		Remark:           t.Remark,
		AttendanceId:     t.AttendanceId,
		AttendanceKindId: t.AttendanceKindId,
		PushedAt:         t.PushedAt,
		CreatedAt:        t.CreatedAt,
		UpdatedAt:        t.UpdatedAt,
	}
}

type AttendanceDetail struct {
	Attendance     `xorm:"extends"`
	ClockedInTime  *AttendanceTime `xorm:"extends"`
	ClockedOutTime *AttendanceTime `xorm:"extends"`
}

func (d AttendanceDetail) build() *Attendance {
	var (
		in  *AttendanceTime
		out *AttendanceTime
	)
	a := d.Attendance
	if d.ClockedInTime.Id != 0 {
		in = d.ClockedInTime.build()
	}
	if d.ClockedOutTime.Id != 0 {
		out = d.ClockedOutTime.build()
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
	return eng.Table(database.AttendanceTable).Count(attendance)
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

	return attendance.build(), nil
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
		Limit(int(p.Limit), int(page)).
		OrderBy("-attendances.id").
		Iterate(&AttendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*AttendanceDetail)
			a := d.build()
			attendances = append(attendances, a)
			return nil
		})
	return attendances, err
}

func updateAttendanceTime(eng Engine, attendanceTime *AttendanceTime) error {
	_, err := eng.Table(database.AttendanceTimeTable).
		Where("attendance_id = ?", attendanceTime.AttendanceId).
		And("attendance_kind_id = ?", attendanceTime.AttendanceKindId).
		And("is_updated = ?", true).
		Update(attendanceTime)
	if err != nil {
		return err
	}
	return nil
}

func createAttendance(eng Engine, a *Attendance) error {
	attendance := new(Attendance)
	attendance.UserId = a.UserId
	attendance.CreatedAt = time.Now()
	attendance.UpdatedAt = time.Now()
	if _, err := eng.Table(database.AttendanceTable).Insert(attendance); err != nil {
		return err
	}
	a.Id = attendance.Id
	return nil
}

func createAttendanceTime(eng Engine, t *AttendanceTime) error {
	at := NewTime(t)
	if _, err := eng.Table(database.AttendanceTimeTable).Insert(at); err != nil {
		return err
	}
	t.Id = at.Id
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
	if !attendance.IsClockedOut() {
		attendanceTime.AttendanceKindId = int64(AttendanceKindClockIn)
	} else {
		attendanceTime.AttendanceKindId = int64(AttendanceKindClockOut)
	}

	// TODO: すでにあるデータに対して、変更済みフラグを更新する

	if err := createAttendanceTime(sess, attendanceTime); err != nil {
		return err
	}
	return nil
}
