package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/timeutil"
	"github.com/KouT127/attendance-management/modules/timezone"
	"github.com/Songmu/flextime"
	"time"
)

func FetchAttendancesCount(ctx context.Context, userID string) (int64, error) {
	var count int64

	dbErr := withDBSession(ctx, func(sess *DBSession) error {
		var err error
		attendance := &models.Attendance{}
		attendance.UserID = userID
		count, err = sess.Count(attendance)
		if err != nil {
			return err
		}
		return nil
	})

	if dbErr != nil {
		return 0, dbErr
	}
	return count, nil
}

func FetchLatestAttendance(ctx context.Context, userID string) (*models.Attendance, error) {
	var attendance models.AttendanceDetail

	dbErr := withDBSession(ctx, func(sess *DBSession) error {
		now := flextime.Now()
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, timezone.JSTLocation())
		end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 59, timezone.JSTLocation())

		_, err := sess.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
			Table(database.AttendanceTable).
			Join("left outer",
				"attendances_time clocked_in_time",
				"attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1 and clocked_in_time.is_modified = false").
			Join("left outer",
				"attendances_time clocked_out_time",
				"attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2 and clocked_out_time.is_modified = false").
			Where("attendances.user_id = ?", userID).
			Where("attendances.created_at Between ? and ? ", start, end).
			Limit(1).
			OrderBy("-attendances.id").
			Get(&attendance)

		if err != nil {
			return err
		}
		return nil
	})

	if dbErr != nil {
		return nil, dbErr
	}
	return attendance.ToAttendance(), nil
}

func FetchAttendances(ctx context.Context, query *models.GetAttendancesParameters) ([]*models.Attendance, error) {
	var attendances []*models.Attendance

	dbErr := withDBSession(ctx, func(sess *DBSession) error {
		dbSess := sess.Select("attendances.*, clocked_in_time.*, clocked_out_time.*").
			Table(database.AttendanceTable).
			Join("left outer",
				"attendances_time clocked_in_time",
				"attendances.id = clocked_in_time.attendance_id and clocked_in_time.attendance_kind_id = 1 and clocked_in_time.is_modified = false").
			Join("left outer",
				"attendances_time clocked_out_time",
				"attendances.id = clocked_out_time.attendance_id and clocked_out_time.attendance_kind_id = 2 and clocked_out_time.is_modified = false")

		if query.Date != nil {
			now := flextime.Now()
			start, end := timeutil.GetMonthRange(now)
			dbSess = dbSess.Where("attendances.created_at Between ? and ? ", start, end)
		}

		p := query.Paginator
		if query.Paginator == nil {
			p = &models.Paginator{}
		}
		if p.Limit == 0 {
			p.Limit = 15
		}
		page := p.CalculatePage()

		err := sess.Limit(int(p.Limit), int(page)).
			Where("attendances.user_id = ?", query.UserID).
			OrderBy("-attendances.id").
			Iterate(&models.AttendanceDetail{}, func(idx int, bean interface{}) error {
				d := bean.(*models.AttendanceDetail)
				a := d.ToAttendance()
				attendances = append(attendances, a)
				return nil
			})
		return err
	})

	if dbErr != nil {
		return nil, dbErr
	}
	return attendances, nil
}

func UpdateOldAttendanceTime(ctx context.Context, id int64, kindID uint8) error {
	return withDBSession(ctx, func(sess *DBSession) error {
		query := &models.AttendanceTime{
			AttendanceID:     id,
			AttendanceKindID: kindID,
			IsModified:       false,
		}
		_, err := sess.UseBool("is_modified").
			Update(&models.AttendanceTime{IsModified: true, UpdatedAt: flextime.Now()}, query)

		if err != nil {
			return err
		}
		return nil
	})
}

func CreateAttendance(ctx context.Context, a *models.Attendance) error {
	return withDBSession(ctx, func(sess *DBSession) error {
		if _, err := sess.Insert(a); err != nil {
			return err
		}
		return nil
	})
}

func CreateAttendanceTime(ctx context.Context, t *models.AttendanceTime) error {
	return withDBSession(ctx, func(sess *DBSession) error {
		if _, err := sess.Insert(t); err != nil {
			return err
		}
		return nil
	})
}
