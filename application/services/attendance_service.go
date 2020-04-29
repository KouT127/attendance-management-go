package services

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/Songmu/flextime"
	"golang.org/x/xerrors"
)

type AttendanceService interface {
	GetAttendances(query models.GetAttendancesParameters) (*models.GetAttendancesResults, error)
	CreateOrUpdateAttendance(attendanceTime *models.AttendanceTime, userId string) (*models.Attendance, error)
}

type attendanceService struct {
	ss sqlstore.SQLStore
}

func NewAttendanceService(ss sqlstore.SQLStore) AttendanceService {
	return attendanceService{
		ss: ss,
	}
}

func (f attendanceService) GetAttendances(params models.GetAttendancesParameters) (*models.GetAttendancesResults, error) {
	ctx := context.Background()
	maxCnt, err := sqlstore.FetchAttendancesCount(ctx, params.UserId)
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

func (f attendanceService) CreateOrUpdateAttendance(attendanceTime *models.AttendanceTime, userId string) (*models.Attendance, error) {
	var (
		attendance *models.Attendance
		err        error
	)

	if userId == "" {
		return nil, xerrors.New("userId is empty")
	}

	if attendanceTime == nil {
		return nil, xerrors.New("attendance time is empty")
	}

	err = f.ss.InTransaction(context.Background(), func(ctx context.Context) error {
		attendance, err = sqlstore.FetchLatestAttendance(ctx, userId)
		if err != nil {
			return err
		}

		if attendance == nil {
			attendance = &models.Attendance{}
			attendance.UserId = userId
			attendance.ClockedIn = attendanceTime
			attendance.CreatedAt = flextime.Now()
			attendance.UpdatedAt = flextime.Now()
			if err := sqlstore.CreateAttendance(ctx, attendance); err != nil {
				return err
			}
			attendanceTime.AttendanceId = attendance.Id
			attendanceTime.AttendanceKindId = uint8(models.AttendanceKindClockIn)
		} else {
			if err := sqlstore.UpdateOldAttendanceTime(ctx, attendance.Id, uint8(models.AttendanceKindClockOut)); err != nil {
				return err
			}
			attendance.ClockedOut = attendanceTime
			attendanceTime.AttendanceId = attendance.Id
			attendanceTime.AttendanceKindId = uint8(models.AttendanceKindClockOut)
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
