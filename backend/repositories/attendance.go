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
	ClockedInId  *int64
	ClockedOutId *int64
	CreatedAt    time.Time `xorm:"created_at"`
	UpdatedAt    time.Time `xorm:"updated_at"`
}

func (a attendance) toAttendanceTime(cit attendanceTime, cot attendanceTime) models.Attendance {
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

func (t attendanceTime) build() models.AttendanceTime {
	return models.AttendanceTime{
		Id:        t.Id,
		Remark:    t.Remark,
		PushedAt:  t.PushedAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type attendanceDetail struct {
	attendance     `xorm:"extends"`
	ClockedInTime  attendanceTime `xorm:"extends"`
	ClockedOutTime attendanceTime `xorm:"extends"`
}

func (d attendanceDetail) build() *models.Attendance {
	a := d.attendance
	i := d.ClockedInTime
	o := d.ClockedOutTime
	attendance := &models.Attendance{
		Id:         a.Id,
		UserId:     a.UserId,
		ClockedIn:  i.build(),
		ClockedOut: o.build(),
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
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
	CreateAttendanceTime(t *models.AttendanceTime) error
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
	return attendance.build(), nil

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
			a := d.attendance.toAttendanceTime(d.ClockedOutTime, d.ClockedOutTime)
			attendances = append(attendances, &a)
			return nil
		})
	return attendances, err
}

func (r attendanceRepository) CreateAttendance(a *models.Attendance) (int64, error) {
	attendance := attendance{
		UserId:       a.UserId,
		ClockedOutId: nil,
		ClockedInId:  nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if a.ClockedIn.Id != 0 {
		attendance.ClockedInId = &a.ClockedIn.Id
	}
	if a.ClockedOut.Id != 0 {
		attendance.ClockedOutId = &a.ClockedOut.Id
	}
	return r.engine.Table(AttendanceTable).Insert(attendance)
}

func (r attendanceRepository) CreateAttendanceTime(t *models.AttendanceTime) error {
	if _, err := r.engine.Table(AttendanceTimeTable).InsertOne(t); err != nil {
		return err
	}
	return nil
}
