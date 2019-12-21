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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type clockedInTime struct {
	Id uint
}
type clockedOutTime struct {
	Id uint
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
			attendance := d.attendance
			cit := d.clockedInTime
			cot := d.clockedOutTime
			print("in ", cit.Id,"out ", cot.Id)
			a := models.Attendance{
				Id:        attendance.Id,
				UserId:    attendance.UserId,
				CreatedAt: attendance.CreatedAt,
				UpdatedAt: attendance.UpdatedAt,
			}
			attendances = append(attendances, &a)
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
