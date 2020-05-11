package services

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/utilities/timeutil"
	"github.com/Songmu/flextime"
	"golang.org/x/xerrors"
)

type AttendanceService interface {
	GetAttendances(ctx context.Context, params models.GetAttendancesParameters) (*models.GetAttendancesResults, error)
	CreateOrUpdateAttendance(ctx context.Context, attendanceTime *models.AttendanceTime, userID string) (*models.Attendance, error)
	GetAttendanceSummary(ctx context.Context, params models.GetAttendanceSummaryParameters) (*models.GetAttendanceSummaryResults, error)
}

type attendanceService struct {
	store sqlstore.SQLStore
}

func NewAttendanceService(ss sqlstore.SQLStore) AttendanceService {
	return &attendanceService{
		store: ss,
	}
}

func (s *attendanceService) GetAttendances(ctx context.Context, params models.GetAttendancesParameters) (*models.GetAttendancesResults, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	maxCnt, err := s.store.GetAttendancesCount(ctx, &params)
	if err != nil {
		return nil, err
	}
	attendances, err := s.store.GetAttendances(ctx, params.UserID, params.Month)
	if err != nil {
		return nil, err
	}

	res := models.GetAttendancesResults{
		MaxCnt:      maxCnt,
		Attendances: attendances,
	}
	return &res, nil
}

func (s *attendanceService) CreateOrUpdateAttendance(ctx context.Context, attendanceTime *models.AttendanceTime, userID string) (*models.Attendance, error) {
	if userID == "" {
		return nil, xerrors.New("userID is empty")
	}
	if attendanceTime == nil {
		return nil, xerrors.New("attendance time is empty")
	}

	ctx, err := s.store.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer s.store.Close(ctx)

	attendance, err := s.store.GetLatestAttendance(ctx, userID)
	if err != nil {
		return nil, err
	}

	if attendance == nil {
		attendance = &models.Attendance{}
		attendance.UserID = userID
		attendance.ClockedIn = attendanceTime
		attendance.CreatedAt = flextime.Now()
		attendance.UpdatedAt = flextime.Now()
		if err = s.store.CreateAttendance(ctx, attendance); err != nil {
			return nil, err
		}
		attendanceTime.AttendanceKindID = uint8(models.AttendanceKindClockIn)
	} else {
		if err = s.store.UpdateOldAttendanceTime(ctx, attendance.ID, uint8(models.AttendanceKindClockOut)); err != nil {
			return nil, err
		}
		attendance.ClockedOut = attendanceTime
		attendanceTime.AttendanceKindID = uint8(models.AttendanceKindClockOut)
	}
	attendanceTime.PushedAt = flextime.Now()
	attendanceTime.CreatedAt = flextime.Now()
	attendanceTime.UpdatedAt = flextime.Now()
	attendanceTime.AttendanceID = attendance.ID

	if err = s.store.CreateAttendanceTime(ctx, attendanceTime); err != nil {
		return nil, err
	}

	if err = s.store.Commit(ctx); err != nil {
		return nil, err
	}
	return attendance, nil
}
func (s *attendanceService) GetAttendanceSummary(ctx context.Context, params models.GetAttendanceSummaryParameters) (*models.GetAttendanceSummaryResults, error) {
	var (
		res models.GetAttendanceSummaryResults
	)

	month, err := timeutil.GetDefaultMonth()
	if err != nil {
		return nil, err
	}

	hour, err := s.store.GetWorkingHours(ctx, flextime.Now())
	if err != nil {
		return nil, err
	}
	if hour.WorkingHours == 0 {
		return nil, xerrors.New("No working hours set for this month")
	}
	res.RequiredHours = hour.WorkingHours

	attendance, err := s.store.GetLatestAttendance(ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	attendances, err := s.store.GetAttendances(ctx, params.UserID, month)
	if err != nil {
		return nil, err
	}

	res.TotalHours = attendances.ManipulateTotalWorkHours()

	if attendance != nil {
		res.LatestAttendance = *attendance
	}
	return &res, nil
}
