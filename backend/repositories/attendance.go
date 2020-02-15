package repositories

import (
	"context"
	"database/sql"
	database "github.com/KouT127/attendance-management/database/gen"
	"github.com/KouT127/attendance-management/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

func NewTime(at *models.AttendanceTime) *database.AttendancesTime {
	t := new(database.AttendancesTime)
	t.ID = at.Id
	t.Remark = null.StringFrom(at.Remark)
	t.CreatedAt = null.TimeFrom(at.CreatedAt)
	t.UpdatedAt = null.TimeFrom(at.UpdatedAt)
	t.PushedAt = at.PushedAt
	return t
}

type AttendanceDetail struct {
	database.Attendance `boil:",bind"`
	ClockedInTime       database.AttendancesTime `boil:",bind"`
	ClockedOutTime      database.AttendancesTime `boil:",bind"`
}

func (d AttendanceDetail) build() *models.Attendance {
	var (
		in  *models.AttendanceTime
		out *models.AttendanceTime
	)
	a := d.Attendance
	if d.ClockedInTime.ID != 0 {
		in = &models.AttendanceTime{
			Id:        d.ClockedInTime.ID,
			Remark:    d.ClockedInTime.Remark.String,
			PushedAt:  d.ClockedInTime.PushedAt,
			CreatedAt: d.ClockedInTime.CreatedAt.Time,
			UpdatedAt: d.ClockedInTime.UpdatedAt.Time,
		}
	}
	if d.ClockedOutTime.ID != 0 {
		out = &models.AttendanceTime{
			Id:        d.ClockedOutTime.ID,
			Remark:    d.ClockedOutTime.Remark.String,
			PushedAt:  d.ClockedOutTime.PushedAt,
			CreatedAt: d.ClockedOutTime.CreatedAt.Time,
			UpdatedAt: d.ClockedOutTime.UpdatedAt.Time,
		}
	}
	attendance := &models.Attendance{
		Id:         a.ID,
		UserId:     a.UserID.String,
		ClockedIn:  in,
		ClockedOut: out,
		CreatedAt:  a.CreatedAt.Time,
		UpdatedAt:  a.UpdatedAt.Time,
	}
	return attendance
}

func NewAttendanceRepository() *attendanceRepository {
	return &attendanceRepository{}
}

type AttendanceRepository interface {
	FetchAttendancesCount(ctx context.Context, db *sql.DB, a *models.Attendance) (int64, error)
	FetchAttendances(ctx context.Context, db *sql.DB, query *models.Attendance, p *Paginator) ([]*models.Attendance, error)
	FetchLatestAttendance(ctx context.Context, db *sql.DB, query *models.Attendance) (*models.Attendance, error)
	CreateAttendance(ctx context.Context, db *sql.DB, a *models.Attendance) error
	UpdateAttendance(ctx context.Context, db *sql.DB, a *models.Attendance) error
	CreateAttendanceTime(ctx context.Context, db *sql.DB, t *models.AttendanceTime) error
	Transaction
}

type attendanceRepository struct {
	transaction
}

func (r *attendanceRepository) FetchAttendancesCount(ctx context.Context, db *sql.DB, a *models.Attendance) (int64, error) {
	cnt, err := database.Attendances(
		Where("user_id = ?", a.UserId),
	).Count(ctx, db)
	return cnt, err
}

func (r *attendanceRepository) FetchLatestAttendance(ctx context.Context, db *sql.DB, query *models.Attendance) (*models.Attendance, error) {
	detail := new(AttendanceDetail)
	err := database.Attendances(
		Select("attendances.*, clocked_in_time.*, clocked_out_time.*"),
		InnerJoin("attendances_time clocked_in_time on clocked_in_time.id = attendances.clocked_in_id"),
		InnerJoin("attendances_time clocked_out_time on clocked_out_time.id = attendances.clocked_out_id"),
		Where("user_id = ?", query.UserId),
		Limit(1),
		OrderBy("-attendances.id"),
	).Bind(ctx, db, detail)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *attendanceRepository) FetchAttendances(ctx context.Context, db *sql.DB, query *models.Attendance, p *Paginator) ([]*models.Attendance, error) {
	attendances := make([]*models.Attendance, 0)
	page := p.CalculatePage()
	details := make([]AttendanceDetail, 0)
	err := database.Attendances(
		Select("attendances.*, clocked_in_time.*, clocked_out_time.*"),
		InnerJoin("attendances_time clocked_in_time on clocked_in_time.id = attendances.clocked_in_id"),
		InnerJoin("attendances_time clocked_out_time on clocked_out_time.id = attendances.clocked_out_id"),
		Where("user_id = ?", query.UserId),
		Limit(int(p.Limit)),
		Offset(int(page)),
		OrderBy("-attendances.id"),
	).Bind(ctx, db, &details)

	if err != nil {
		return nil, err
	}

	for _, detail := range details {
		attendances = append(attendances, detail.build())
	}
	return attendances, err
}

func (r *attendanceRepository) CreateAttendance(ctx context.Context, db *sql.DB, a *models.Attendance) error {
	attendance := new(database.Attendance)
	attendance.UserID = null.StringFrom(a.UserId)
	attendance.CreatedAt = null.TimeFrom(time.Now())
	attendance.UpdatedAt = null.TimeFrom(time.Now())

	if a.ClockedIn.Id != 0 {
		attendance.ClockedInID = null.UintFrom(a.ClockedIn.Id)
	}
	return attendance.Insert(ctx, db, boil.Infer())
}

func (r *attendanceRepository) UpdateAttendance(ctx context.Context, db *sql.DB, a *models.Attendance) error {
	attendance, err := database.FindAttendance(ctx, db, a.Id)
	if err != nil {
		return err
	}
	attendance.ClockedOutID = null.UintFrom(a.ClockedOut.Id)
	attendance.UpdatedAt = null.TimeFrom(time.Now())
	if a.ClockedOut.Id != 0 {
		attendance.ClockedOutID = null.UintFrom(a.ClockedOut.Id)
	}
	_, err = attendance.Update(ctx, db, boil.Whitelist("clocked_out_id", "updated_at"))
	return err
}

func (r *attendanceRepository) CreateAttendanceTime(ctx context.Context, db *sql.DB, t *models.AttendanceTime) error {
	attendanceTime := NewTime(t)
	if err := attendanceTime.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}
	t.Id = attendanceTime.ID
	return nil
}
