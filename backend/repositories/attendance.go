package repositories

import (
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
	Id        uint
	UserId    string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type clockedInTime struct {
	Id        uint
	PushedAt  time.Time
	Remark    string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
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

func NewAttendanceRepository(e xorm.Engine) *attendanceRepository {
	return &attendanceRepository{
		engine: e,
	}
}

type AttendanceRepository interface {
	FetchAttendancesCount(a *models.Attendance) (int64, error)
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

func (r attendanceRepository) FetchAttendances(a *models.Attendance, p *Pagination) ([]*models.Attendance, error) {
	attendances := make([]*models.Attendance, 0)

	page := p.CalculatePage()
	err := r.engine.
		Select("attendances.*, clockedInTime.*, clockedOutTime.*").
		Table("attendances").
		Join("left", "attendances_time clockedInTime", "attendances.clocked_in_id = clockedInTime.id").
		Join("left", "attendances_time clockedOutTime", "attendances.clocked_out_id = clockedOutTime.id").
		Limit(int(p.Limit), int(page)).
		OrderBy("-attendances.id").
		Iterate(&attendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*attendanceDetail)
			a := d.attendance
			cit := d.clockedInTime
			cot := d.clockedOutTime
			attendance := models.Attendance{
				Id:     a.Id,
				UserId: a.UserId,
				ClockedIn: models.AttendanceTime{Id:
				cit.Id, PushedAt:
				cit.PushedAt,
					Remark: cit.Remark,
				},
				ClockedOut: models.AttendanceTime{
					Id:       cot.Id,
					PushedAt: cot.PushedAt,
					Remark:   cot.Remark,
				},
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.UpdatedAt,
			}
			attendances = append(attendances, &attendance)
			return nil
		})
	return attendances, err
}

func (r attendanceRepository) CreateAttendance(a *models.Attendance) (int64, error) {
	return r.engine.Table("attendances").Insert(a)
}

func (r attendanceRepository) CreateAttendanceTime(t *models.AttendanceTime) (int64, error) {
	return r.engine.Table(AttendanceTimeTable).Insert(t)
}
