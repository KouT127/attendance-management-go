package usecases

import (
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/repositories"
	. "github.com/KouT127/attendance-management/backend/serializers"
	. "github.com/KouT127/attendance-management/backend/validators"
)

func NewAttendanceInteractor(repo AttendanceRepository) *attendanceInteractor {
	return &attendanceInteractor{
		repository: repo,
	}
}

type AttendanceInteractor interface {
	ViewAttendances(pagination *Pagination, attendance *Attendance) (*AttendancesResponse, error)
	CreateAttendance(query *Attendance, time *AttendanceTime) (*AttendanceResponse, error)
}

type attendanceInteractor struct {
	repository AttendanceRepository
}

func (interactor *attendanceInteractor) ViewAttendances(pagination *Pagination, attendance *Attendance) (*AttendancesResponse, error) {
	maxCnt, err := interactor.repository.FetchAttendancesCount(attendance)
	if err != nil {
		return nil, err
	}

	attendances := make([]*Attendance, 0)
	attendances, err = interactor.repository.FetchAttendances(attendance, pagination)
	if err != nil {
		return nil, err
	}

	responses := make([]*AttendanceResponse, 0)

	for _, attendance := range attendances {
		res := &AttendanceResponse{}
		res.Build(attendance)
		responses = append(responses, res)
	}

	res := new(AttendancesResponse)
	res.HasNext = pagination.HasNext(maxCnt)
	res.IsSuccessful = true
	res.Attendances = responses
	return res, nil
}

func (interactor *attendanceInteractor) CreateAttendance(query *Attendance, time *AttendanceTime) (*AttendanceResponse, error) {
	attendance, err := interactor.repository.FetchLatestAttendance(query)
	if err != nil {
		return nil, err
	}

	if err := interactor.repository.CreateAttendanceTime(time); err != nil {
		return nil, err
	}

	if attendance == nil {
		attendance = &Attendance{
			UserId:    query.UserId,
			ClockedIn: *time,
		}
		if _, err := interactor.repository.CreateAttendance(attendance); err != nil {
			return nil, err
		}
	} else {
		attendance = &Attendance{
			Id:         attendance.Id,
			UserId:     attendance.UserId,
			ClockedIn:  attendance.ClockedIn,
			ClockedOut: *time,
		}
		if _, err := interactor.repository.UpdateAttendance(attendance); err != nil {
			return nil, err
		}
	}

	res := new(AttendanceResponse)
	res.Build(attendance)
	return res, nil
}
