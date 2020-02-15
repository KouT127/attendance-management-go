package usecases

import (
	"context"
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/repositories"
	. "github.com/KouT127/attendance-management/responses"
)

func NewAttendanceUsecase(ar repositories.AttendanceRepository) *attendanceUsecase {
	return &attendanceUsecase{
		ar: ar,
	}
}

type AttendanceUsecase interface {
	ViewAttendances(pagination *PaginatorInput, attendance *models.Attendance) (*AttendancesResult, error)
	ViewLatestAttendance(attendance *models.Attendance) (*AttendanceResult, error)
	ViewAttendancesMonthly(pagination *PaginatorInput, attendance *models.Attendance) (*AttendancesResult, error)
	CreateAttendance(input *AttendanceInput, query *models.Attendance) (*AttendanceResult, error)
}

type attendanceUsecase struct {
	ar repositories.AttendanceRepository
}

func (i *attendanceUsecase) ViewAttendances(pagination *PaginatorInput, attendance *models.Attendance) (*AttendancesResult, error) {
	db := database.NewDB()
	ctx := context.Background()
	maxCnt, err := i.ar.FetchAttendancesCount(ctx, db, attendance)
	if err != nil {
		return nil, err
	}

	attendances := make([]*models.Attendance, 0)
	attendances, err = i.ar.FetchAttendances(ctx, db, attendance, pagination.BuildPaginator())
	if err != nil {
		return nil, err
	}

	responses := make([]*AttendanceResp, 0)

	for _, attendance := range attendances {
		resp := NewAttendanceResp(attendance)
		responses = append(responses, &resp)
	}

	res := new(AttendancesResult)
	res.HasNext = pagination.HasNext(maxCnt)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (i *attendanceUsecase) ViewLatestAttendance(attendance *models.Attendance) (*AttendanceResult, error) {
	db := database.NewDB()
	ctx := context.Background()
	err := i.ar.FetchLatestAttendance(ctx, db, attendance)
	if err != nil {
		return nil, err
	}
	attendance.IsClockedOut()

	s := new(AttendanceResult)
	s.NewAttendanceResult(true, attendance)
	return s, nil
}

func (i *attendanceUsecase) ViewAttendancesMonthly(pagination *PaginatorInput, attendance *models.Attendance) (*AttendancesResult, error) {
	db := database.NewDB()
	ctx := context.Background()

	attendances, err := i.ar.FetchAttendances(ctx, db, attendance, pagination.BuildPaginator())
	if err != nil {
		return nil, err
	}

	responses := make([]*AttendanceResp, 0)

	for _, attendance := range attendances {
		resp := NewAttendanceResp(attendance)
		responses = append(responses, &resp)
	}

	res := new(AttendancesResult)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (i *attendanceUsecase) CreateAttendance(input *AttendanceInput, attendance *models.Attendance) (*AttendanceResult, error) {
	db := database.NewDB()
	ctx := context.Background()
	if err := input.Validate(); err != nil {
		return nil, err
	}
	time := input.BuildAttendanceTime()

	err := i.ar.FetchLatestAttendance(ctx, db, attendance)
	if err != nil {
		return nil, err
	}

	if err := i.ar.CreateAttendanceTime(ctx, db, time); err != nil {
		return nil, err
	}

	if attendance.Id == 0 {
		attendance.ClockIn(time)
		if err := i.ar.CreateAttendance(ctx, db, attendance); err != nil {
			return nil, err
		}

		serializer := NewAttendanceResult(attendance)
		return serializer, nil
	}

	attendance.ClockOut(time)
	if err := i.ar.UpdateAttendance(ctx, db, attendance); err != nil {
		return nil, err
	}

	serializer := NewAttendanceResult(attendance)
	return serializer, nil
}
