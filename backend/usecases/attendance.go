package usecases

import (
	"github.com/KouT127/attendance-management/backend/database"
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/repositories"
	. "github.com/KouT127/attendance-management/backend/serializers"
)

func NewAttendanceUsecase(ar AttendanceRepository) *attendanceUsecase {
	return &attendanceUsecase{
		ar: ar,
	}
}

type AttendanceInteractor interface {
	ViewAttendances(pagination *PaginatorInput, attendance *Attendance) (*AttendancesSerializer, error)
	ViewAttendancesMonthly(pagination *PaginatorInput, attendance *Attendance) (*AttendancesSerializer, error)
	CreateAttendance(input *AttendanceInput, query *Attendance) (*AttendanceSerializer, error)
}

type attendanceUsecase struct {
	ar AttendanceRepository
}

func (i *attendanceUsecase) ViewAttendances(pagination *PaginatorInput, attendance *Attendance) (*AttendancesSerializer, error) {
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

	responses := make([]*AttendanceResponse, 0)

	for _, attendance := range attendances {
		res := &AttendanceResponse{}
		res.Build(attendance)
		responses = append(responses, res)
	}

	res := new(AttendancesSerializer)
	res.HasNext = pagination.HasNext(maxCnt)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (i *attendanceUsecase) ViewAttendancesMonthly(pagination *PaginatorInput, attendance *Attendance) (*AttendancesSerializer, error) {
	eng := database.NewDB()

	attendances, err := i.ar.FetchAttendances(eng, attendance, pagination.BuildPaginator())
	if err != nil {
		return nil, err
	}

	responses := make([]*AttendanceResponse, 0)

	for _, attendance := range attendances {
		res := new(AttendanceResponse)
		res.Build(attendance)
		responses = append(responses, res)
	}

	res := new(AttendancesSerializer)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (i *attendanceUsecase) CreateAttendance(input *AttendanceInput, query *Attendance) (*AttendanceSerializer, error) {
	sess := i.ar.NewSession(database.NewDB())
	defer i.ar.Close(sess)
	if err := i.ar.Begin(sess); err != nil {
		return nil, err
	}

	if err := input.Validate(); err != nil {
		return nil, err
	}
	time := input.BuildAttendanceTime()

	attendance, err := i.ar.FetchLatestAttendance(sess, query)
	if err != nil {
		return nil, err
	}

	if err := i.ar.CreateAttendanceTime(sess, time); err != nil {
		return nil, err
	}

	if attendance == nil {
		attendance = &Attendance{
			UserId:    query.UserId,
			ClockedIn: time,
		}
		if _, err := i.ar.CreateAttendance(sess, attendance); err != nil {
			return nil, err
		}
	} else {
		attendance = &Attendance{
			Id:         attendance.Id,
			UserId:     attendance.UserId,
			ClockedIn:  attendance.ClockedIn,
			ClockedOut: time,
		}
		if _, err := i.ar.UpdateAttendance(sess, attendance); err != nil {
			return nil, err
		}
	}

	serializer := new(AttendanceSerializer)
	serializer.Serialize(true, false, attendance)

	if err := i.ar.Commit(sess); err != nil {
		return nil, err
	}
	return serializer, nil
}
