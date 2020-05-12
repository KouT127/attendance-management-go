package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/utilities/timeutil"
	"github.com/KouT127/attendance-management/utilities/timezone"
	"github.com/Songmu/flextime"
	"golang.org/x/xerrors"
	"time"
)

type Attendance interface {
	GetAttendancesCount(ctx context.Context, query *models.GetAttendancesParameters) (int64, error)
	GetLatestAttendance(ctx context.Context, userID string) (*models.Attendance, error)
	GetAttendances(ctx context.Context, userID string, month int) (models.Attendances, error)
	UpdateOldAttendanceTime(ctx context.Context, id int64, kindID uint8) error
	CreateAttendance(ctx context.Context, attendance *models.Attendance) error
	CreateAttendanceTime(ctx context.Context, attendanceTime *models.AttendanceTime) error
}

func (sqlStore) GetAttendancesCount(ctx context.Context, query *models.GetAttendancesParameters) (int64, error) {
	var count int64

	sess, err := getDBSession(ctx)
	if err != nil {
		return 0, err
	}

	if query.Month == 0 {
		return 0, xerrors.New("month is empty")
	}

	start, end, err := timeutil.GetMonthRange(query.Month)
	if err != nil {
		return 0, err
	}
	attendance := &models.Attendance{}
	attendance.UserID = query.UserID
	count, err = sess.Where("attendances.attended_at Between ? and ? ", start, end).Count(attendance)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (sqlStore) GetLatestAttendance(ctx context.Context, userID string) (*models.Attendance, error) {
	var (
		attendance models.AttendanceDetail
		has        bool
	)

	sess, err := getDBSession(ctx)
	if err != nil {
		return nil, err
	}
	now := flextime.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, timezone.JSTLocation())
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, timezone.JSTLocation())

	has, err = sess.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
		Table(AttendanceTable).
		Join("left outer",
			"attendances_time clocked_in_time",
			"attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1 and clocked_in_time.is_modified = false").
		Join("left outer",
			"attendances_time clocked_out_time",
			"attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2 and clocked_out_time.is_modified = false").
		Where("attendances.user_id = ?", userID).
		Where("attendances.attended_at Between ? and ? ", start, end).
		Limit(1).
		OrderBy("-attendances.id").
		Get(&attendance)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, nil
	}
	return attendance.ToAttendance(), nil
}

func (sqlStore) GetAttendances(ctx context.Context, userID string, month int) (models.Attendances, error) {
	attendances := make(models.Attendances, 0)
	eng.NoAutoTime()
	sess, err := getDBSession(ctx)
	if err != nil {
		return nil, err
	}
	if month == 0 {
		return nil, xerrors.New("month is empty")
	}

	start, end, err := timeutil.GetMonthRange(month)
	if err != nil {
		return nil, err
	}

	err = sess.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
		Table(AttendanceTable).
		Join("left outer",
			"attendances_time clocked_in_time",
			"attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1 and clocked_in_time.is_modified = false").
		Join("left outer",
			"attendances_time clocked_out_time",
			"attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2 and clocked_out_time.is_modified = false").
		Where("attendances.attended_at Between ? and ? ", start, end).
		Where("attendances.user_id = ?", userID).
		OrderBy("-attendances.id").
		Iterate(&models.AttendanceDetail{}, func(idx int, bean interface{}) error {
			d := bean.(*models.AttendanceDetail)
			a := d.ToAttendance()
			attendances = append(attendances, a)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return attendances, nil
}

func (sqlStore) UpdateOldAttendanceTime(ctx context.Context, id int64, kindID uint8) error {
	sess, err := getDBSession(ctx)
	if err != nil {
		return err
	}
	query := &models.AttendanceTime{
		AttendanceID:     id,
		AttendanceKindID: kindID,
		IsModified:       false,
	}
	_, err = sess.UseBool("is_modified").
		Update(&models.AttendanceTime{IsModified: true}, query)

	if err != nil {
		return err
	}
	return nil
}

func (sqlStore) CreateAttendance(ctx context.Context, attendance *models.Attendance) error {
	sess, err := getDBSession(ctx)
	if err != nil {
		return err
	}
	if _, err := sess.Insert(attendance); err != nil {
		return err
	}
	return nil
}

func (sqlStore) CreateAttendanceTime(ctx context.Context, attendanceTime *models.AttendanceTime) error {
	sess, err := getDBSession(ctx)
	if err != nil {
		return err
	}
	if _, err := sess.Insert(attendanceTime); err != nil {
		return err
	}
	return nil

}
