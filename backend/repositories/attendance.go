package repositories

import (
	"errors"
	"github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/validators"
	"github.com/go-xorm/xorm"
	"time"
)

const (
	AttendanceTable     = "attendances"
	AttendanceTimeTable = "attendances_time"
)

type attendance struct {
	Id           uint
	UserId       string
	ClockedInId  *uint
	ClockedOutId *uint
	CreatedAt    time.Time `xorm:"created_at"`
	UpdatedAt    time.Time `xorm:"updated_at"`
}

func (a attendance) toAttendanceTime(cit clockedInTime, cot clockedOutTime) models.Attendance {
	return models.Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  cit.toAttendanceTime(),
		ClockedOut: cot.toAttendanceTime(),
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

func (t clockedInTime) toAttendanceTime() models.AttendanceTime {
	return models.AttendanceTime{
		Id:        t.Id,
		Remark:    t.Remark,
		PushedAt:  t.PushedAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type clockedInTime struct {
	Id        uint
	PushedAt  time.Time
	Remark    string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

func (t clockedOutTime) toAttendanceTime() models.AttendanceTime {
	return models.AttendanceTime{
		Id:        t.Id,
		Remark:    t.Remark,
		PushedAt:  t.PushedAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type clockedOutTime struct {
	Id        uint
	PushedAt  time.Time
	Remark    string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type attendanceDetail struct {
	attendance     `xorm:"extends"`
	clockedInTime  `xorm:"extends"`
	clockedOutTime `xorm:"extends"`
}

func (d attendanceDetail) toAttendance() *models.Attendance {
	a := d.attendance
	i := d.clockedInTime
	o := d.clockedOutTime
	attendance := &models.Attendance{
		Id:     a.Id,
		UserId: a.UserId,
		ClockedIn: models.AttendanceTime{
			Id:        i.Id,
			Remark:    i.Remark,
			PushedAt:  i.PushedAt,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		},
		ClockedOut: models.AttendanceTime{
			Id:        o.Id,
			Remark:    o.Remark,
			PushedAt:  o.PushedAt,
			CreatedAt: o.CreatedAt,
			UpdatedAt: o.UpdatedAt,
		},
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
	return attendance
}

func NewAttendanceRepository(e xorm.Engine) *attendanceRepository {
	return &attendanceRepository{
		engine: e,
	}
}

type AttendanceRepository interface {
	FetchAttendancesCount(a *models.Attendance) (int64, error)
	FetchLatestAttendance(a *models.Attendance) (*models.Attendance, error)
	FetchAttendances(a *models.Attendance, p *Pagination) ([]*models.Attendance, error)
	CreateAttendance(a *models.Attendance) (int64, error)
	CreateAttendanceTime(t *models.AttendanceTime) (int64, error)
}

type attendanceRepository struct {
	engine xorm.Engine
}

func (r attendanceRepository) FetchAttendancesCount(a *models.Attendance) (int64, error) {
	attendance := &attendance{Id: a.Id}
	return r.engine.Table(AttendanceTable).Count(attendance)
}

func (r attendanceRepository) FetchLatestAttendance(a *models.Attendance) (*models.Attendance, error) {
	attendance := &attendanceDetail{}
	has, err := r.engine.Select("attendances.*, clockedInTime.*, clockedOutTime.*").
		Table(AttendanceTable).
		Join("left", "attendances_time clockedInTime", "attendances.clocked_in_id = clockedInTime.id").
		Join("left", "attendances_time clockedOutTime", "attendances.clocked_out_id = clockedOutTime.id").
		Where("attendances.user_id = ?", a.UserId).
		Limit(1).
		OrderBy("-attendances.id").
		Get(attendance)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("attendance not exists")
	}
	return attendance.toAttendance(), nil

}

func (r attendanceRepository) FetchAttendances(a *models.Attendance, p *Pagination) ([]*models.Attendance, error) {
	attendances := make([]*models.Attendance, 0)
	page := p.CalculatePage()
	err := r.engine.
		Select("attendances.*, clockedInTime.*, clockedOutTime.*").
		Table(AttendanceTable).
		Join("left", "attendances_time clockedInTime", "attendances.clocked_in_id = clockedInTime.id").
		Join("left", "attendances_time clockedOutTime", "attendances.clocked_out_id = clockedOutTime.id").
		Limit(int(p.Limit), int(page)).
		OrderBy("-attendances.id").
		Iterate(&attendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*attendanceDetail)
			a := d.attendance.toAttendanceTime(d.clockedInTime, d.clockedOutTime)
			attendances = append(attendances, &a)
			return nil
		})
	return attendances, err
}

func (r attendanceRepository) CreateAttendance(a *models.Attendance) (int64, error) {
	attendance := attendance{
		UserId:    a.UserId,
		ClockedOutId: nil,
		ClockedInId: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if a.ClockedIn.Id != 0 {
		attendance.ClockedInId = &a.ClockedIn.Id
	}
	if a.ClockedOut.Id != 0 {
		attendance.ClockedOutId = &a.ClockedOut.Id
	}
	return r.engine.Table(AttendanceTable).Insert(attendance)
}

func (r attendanceRepository) CreateAttendanceTime(t *models.AttendanceTime) (int64, error) {
	return r.engine.Table(AttendanceTimeTable).Insert(t)
}
