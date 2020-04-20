package facades

import (
	"context"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/models"
	"github.com/Songmu/flextime"
	"golang.org/x/xerrors"
)

type AttendanceFacade interface {
	GetAttendances(query models.GetAttendancesParameters) (*models.GetAttendancesResults, error)
	CreateOrUpdateAttendance(attendanceTime *models.AttendanceTime, userID string) (*models.Attendance, error)
}

type attendanceFacade struct {
	ss sqlstore.SQLStore
}

func NewAttendanceFacade(ss sqlstore.SQLStore) AttendanceFacade {
	return attendanceFacade{
		ss: ss,
	}
}

func (f attendanceFacade) GetAttendances(params models.GetAttendancesParameters) (*models.GetAttendancesResults, error) {
	ctx := context.Background()
	maxCnt, err := sqlstore.FetchAttendancesCount(ctx, params.UserID)
	if err != nil {
		return nil, err
	}
	attendances, err := sqlstore.FetchAttendances(ctx, &params)
	if err != nil {
		return nil, err
	}

	res := models.GetAttendancesResults{
		MaxCnt:      maxCnt,
		Attendances: attendances,
	}
	return &res, nil
}

func (f attendanceFacade) CreateOrUpdateAttendance(attendanceTime *models.AttendanceTime, userID string) (*models.Attendance, error) {
	var (
		attendance *models.Attendance
		err        error
	)

	if userID == "" {
		return nil, xerrors.New("userID is empty")
	}

	if attendanceTime == nil {
		return nil, xerrors.New("attendance time is empty")
	}

	err = f.ss.InTransaction(context.Background(), func(ctx context.Context) error {
		attendance, err = sqlstore.FetchLatestAttendance(ctx, userID)
		if err != nil {
			return err
		}

		if attendance == nil {
			attendance = &models.Attendance{}
			attendance.UserID = userID
			attendance.ClockedIn = attendanceTime
			attendance.CreatedAt = flextime.Now()
			attendance.UpdatedAt = flextime.Now()
			if err := sqlstore.CreateAttendance(ctx, attendance); err != nil {
				return err
			}
			attendanceTime.AttendanceID = attendance.ID
			attendanceTime.AttendanceKindID = uint8(models.AttendanceKindClockIn)
		} else {
			if err := sqlstore.UpdateOldAttendanceTime(ctx, attendance.ID, uint8(models.AttendanceKindClockOut)); err != nil {
				return err
			}
			attendance.ClockedOut = attendanceTime
			attendanceTime.AttendanceID = attendance.ID
			attendanceTime.AttendanceKindID = uint8(models.AttendanceKindClockOut)
		}

		if err := sqlstore.CreateAttendanceTime(ctx, attendanceTime); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return attendance, nil
}
