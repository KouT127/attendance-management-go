package repositories

import (
	. "github.com/KouT127/Attendance-management/backend/domains"
	. "github.com/KouT127/Attendance-management/backend/validators"
	"github.com/go-xorm/xorm"
)

const (
	AttendanceTable = "attendances"
)

func NewAttendanceRepository(e xorm.Engine) *attendanceRepository {
	return &attendanceRepository{
		engine: e,
	}
}

type AttendanceRepository interface {
	FetchAttendancesCount(a *Attendance) (int64, error)
	FetchAttendances(a *Attendance, p *Pagination) ([]*Attendance, error)
	CreateAttendance(a *Attendance) (int64, error)
}

type attendanceRepository struct {
	engine xorm.Engine
}

func (r attendanceRepository) FetchAttendancesCount(a *Attendance) (int64, error) {
	return r.engine.Table(AttendanceTable).Count(a)
}

func (r attendanceRepository) FetchAttendances(a *Attendance, p *Pagination) ([]*Attendance, error) {
	attendances := make([]*Attendance, 0)
	page := p.CalculatePage()
	err := r.engine.Table("attendances").
		Limit(int(p.Limit), int(page)).
		OrderBy("-id").
		Iterate(a, func(idx int, bean interface{}) error {
			attendance := bean.(*Attendance)
			attendances = append(attendances, attendance)
			return nil
		})
	return attendances, err
}

func (r attendanceRepository) CreateAttendance(a *Attendance) (int64, error) {
	return r.engine.Table("attendances").Insert(a)
}
