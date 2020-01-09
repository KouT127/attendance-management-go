package repositories

import (
	"github.com/KouT127/attendance-management/backend/models"
	"github.com/go-xorm/xorm"
	"time"
)

const (
	AttendanceTable     = "attendances"
	AttendanceTimeTable = "attendances_time"
)

type attendance struct {
	Id           int64
	UserId       string
	ClockedInId  *int64
	ClockedOutId *int64
	CreatedAt    time.Time `xorm:"created_at"`
	UpdatedAt    time.Time `xorm:"updated_at"`
}

func (a *attendance) build(cit *attendanceTime, cot *attendanceTime) models.Attendance {
	return models.Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  cit.build(),
		ClockedOut: cot.build(),
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

type attendanceTime struct {
	Id        int64
	PushedAt  time.Time
	Remark    string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

func NewTime(at *models.AttendanceTime) *attendanceTime {
	t := new(attendanceTime)
	t.Id = at.Id
	t.Remark = at.Remark
	t.CreatedAt = at.CreatedAt
	t.UpdatedAt = at.UpdatedAt
	t.PushedAt = at.PushedAt
	return t
}

func (t attendanceTime) build() *models.AttendanceTime {
	return &models.AttendanceTime{
		Id:        t.Id,
		Remark:    t.Remark,
		PushedAt:  t.PushedAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type attendanceDetail struct {
	attendance     `xorm:"extends"`
	ClockedInTime  *attendanceTime `xorm:"extends"`
	ClockedOutTime *attendanceTime `xorm:"extends"`
}

func (d attendanceDetail) build() *models.Attendance {
	var (
		in  *models.AttendanceTime
		out *models.AttendanceTime
	)
	a := d.attendance
	if d.ClockedInTime.Id != 0 {
		in = d.ClockedInTime.build()
	}
	if d.ClockedOutTime.Id != 0 {
		out = d.ClockedOutTime.build()
	}

	attendance := &models.Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  in,
		ClockedOut: out,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
	return attendance
}

func NewAttendanceRepository() *attendanceRepository {
	return &attendanceRepository{}
}

type AttendanceRepository interface {
	FetchAttendancesCount(eng *xorm.Engine, a *models.Attendance) (int64, error)
	FetchAttendances(eng *xorm.Engine, a *models.Attendance, p *Paginator) ([]*models.Attendance, error)
	FetchLatestAttendance(eng *xorm.Engine, a *models.Attendance) (*models.Attendance, error)
	CreateAttendance(sess *xorm.Session, a *models.Attendance) (int64, error)
	UpdateAttendance(sess *xorm.Session, a *models.Attendance) (int64, error)
	CreateAttendanceTime(sess *xorm.Session, t *models.AttendanceTime) error
	Transaction
}

type attendanceRepository struct {
	transaction
}

func (r attendanceRepository) FetchAttendancesCount(eng *xorm.Engine, a *models.Attendance) (int64, error) {
	attendance := &attendance{}
	attendance.Id = a.Id
	return eng.Table(AttendanceTable).Count(attendance)
}

func (r attendanceRepository) FetchLatestAttendance(eng *xorm.Engine, a *models.Attendance) (*models.Attendance, error) {
	attendance := &attendanceDetail{}
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, time.Local)

	has, err := eng.Select("attendances.*, clockedInTime.*, clockedOutTime.*").
		Table(AttendanceTable).
		Join("left", "attendances_time clockedInTime", "attendances.clocked_in_id = clockedInTime.id").
		Join("left", "attendances_time clockedOutTime", "attendances.clocked_out_id = clockedOutTime.id").
		Where("attendances.user_id = ?", a.UserId).
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

func (r attendanceRepository) FetchAttendances(eng *xorm.Engine, a *models.Attendance, p *Paginator) ([]*models.Attendance, error) {
	attendances := make([]*models.Attendance, 0)
	page := p.CalculatePage()
	err := eng.
		Select("attendances.*, clockedInTime.*, clockedOutTime.*").
		Table(AttendanceTable).
		Join("left", "attendances_time clockedInTime", "attendances.clocked_in_id = clockedInTime.id").
		Join("left", "attendances_time clockedOutTime", "attendances.clocked_out_id = clockedOutTime.id").
		Limit(int(p.Limit), int(page)).
		OrderBy("-attendances.id").
		Iterate(&attendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*attendanceDetail)
			a := d.build()
			attendances = append(attendances, a)
			return nil
		})
	return attendances, err
}

func (r attendanceRepository) CreateAttendance(sess *xorm.Session, a *models.Attendance) (int64, error) {
	attendance := attendance{
		UserId:       a.UserId,
		ClockedOutId: nil,
		ClockedInId:  nil,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	if a.ClockedIn.Id != 0 {
		attendance.ClockedInId = &a.ClockedIn.Id
	}
	return sess.Table(AttendanceTable).Insert(attendance)
}

func (r attendanceRepository) UpdateAttendance(sess *xorm.Session, a *models.Attendance) (int64, error) {
	attendance := attendance{
		ClockedOutId: &a.ClockedOut.Id,
		UpdatedAt:    time.Now(),
	}
	if a.ClockedOut.Id != 0 {
		attendance.ClockedOutId = &a.ClockedOut.Id
	}
	return sess.Table(AttendanceTable).ID(a.Id).Cols("clocked_out_id", "updated_at").Update(&attendance)
}

func (r attendanceRepository) CreateAttendanceTime(sess *xorm.Session, t *models.AttendanceTime) error {
	at := NewTime(t)
	if _, err := sess.Table(AttendanceTimeTable).Insert(at); err != nil {
		return err
	}
	t.Id = at.Id
	return nil
}
