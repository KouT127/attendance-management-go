package usecases

import (
	"github.com/KouT127/attendance-management/database"
	. "github.com/KouT127/attendance-management/models"
	. "github.com/KouT127/attendance-management/repositories"
	. "github.com/KouT127/attendance-management/serializers"
)

func NewAttendanceUsecase(ar AttendanceRepository) *attendanceUsecase {
	return &attendanceUsecase{
		ar: ar,
	}
}

type AttendanceUsecase interface {
	ViewAttendances(pagination *PaginatorInput, attendance *Attendance) (*AttendancesResult, error)
	ViewLatestAttendance(attendance *Attendance) (*AttendanceResult, error)
	ViewAttendancesMonthly(pagination *PaginatorInput, attendance *Attendance) (*AttendancesResult, error)
	CreateAttendance(input *AttendanceInput, query *Attendance) (*AttendanceResult, error)
}

type attendanceUsecase struct {
	ar AttendanceRepository
}

func (i *attendanceUsecase) ViewAttendances(pagination *PaginatorInput, attendance *Attendance) (*AttendancesResult, error) {
	eng := database.NewDB()
	maxCnt, err := i.ar.FetchAttendancesCount(eng, attendance)
	if err != nil {
		return nil, err
	}

	attendances := make([]*Attendance, 0)
	attendances, err = i.ar.FetchAttendances(eng, attendance, pagination.BuildPaginator())
	if err != nil {
		return nil, err
	}

	responses := make([]*AttendanceResp, 0)

	for _, attendance := range attendances {
		res := &AttendanceResp{}
		res.NewAttendanceResp(attendance)
		responses = append(responses, res)
	}

	res := new(AttendancesResult)
	res.HasNext = pagination.HasNext(maxCnt)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (i *attendanceUsecase) ViewLatestAttendance(attendance *Attendance) (*AttendanceResult, error) {
	eng := database.NewDB()

	attendance, err := i.ar.FetchLatestAttendance(eng, attendance)
	if err != nil {
		return nil, err
	}
	attendance.IsClockedOut()

	s := new(AttendanceResult)
	s.NewAttendanceResult(true, attendance)
	return s, nil
}

func (i *attendanceUsecase) ViewAttendancesMonthly(pagination *PaginatorInput, attendance *Attendance) (*AttendancesResult, error) {
	eng := database.NewDB()

	attendances, err := i.ar.FetchAttendances(eng, attendance, pagination.BuildPaginator())
	if err != nil {
		return nil, err
	}

	responses := make([]*AttendanceResp, 0)

	for _, attendance := range attendances {
		res := new(AttendanceResp)
		res.NewAttendanceResp(attendance)
		responses = append(responses, res)
	}

	res := new(AttendancesResult)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (i *attendanceUsecase) CreateAttendance(input *AttendanceInput, query *Attendance) (*AttendanceResult, error) {
	eng := database.NewDB()
	sess := i.ar.NewSession(eng)
	defer i.ar.Close(sess)
	if err := i.ar.Begin(sess); err != nil {
		return nil, err
	}

	if err := input.Validate(); err != nil {
		return nil, err
	}
	time := input.BuildAttendanceTime()

	attendance, err := i.ar.FetchLatestAttendance(eng, query)
	if err != nil {
		return nil, err
	}

	if err := i.ar.CreateAttendanceTime(sess, time); err != nil {
		return nil, err
	}

	if attendance == nil {
		attendance = new(Attendance)
		attendance.ClockIn(query.UserId, time)
		if _, err := i.ar.CreateAttendance(sess, attendance); err != nil {
			return nil, err
		}

		serializer := NewAttendanceSerializer(attendance)
		if err := i.ar.Commit(sess); err != nil {
			return nil, err
		}
		return serializer, nil
	}

	attendance.ClockOut(time)
	if _, err := i.ar.UpdateAttendance(sess, attendance); err != nil {
		return nil, err
	}

	serializer := NewAttendanceSerializer(attendance)
	if err := i.ar.Commit(sess); err != nil {
		return nil, err
	}
	return serializer, nil
}
