package attendance

import (
	. "github.com/KouT127/attendance-management/models"
	. "github.com/KouT127/attendance-management/responses"
	. "github.com/KouT127/attendance-management/usecases"
)

func ViewAttendances(pagination *PaginatorInput, attendance *Attendance) ([]*Attendance, error) {
	attendances := make([]*Attendance, 0)
	attendances, err := FetchAttendances(attendance, pagination.BuildPaginator())
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func ViewLatestAttendance(attendance *Attendance) (*AttendanceResult, error) {
	attendance, err := FetchLatestAttendance(attendance)
	if err != nil {
		return nil, err
	}

	s := new(AttendanceResult)
	s.NewAttendanceResult(true, attendance)
	return s, nil
}

func ViewAttendancesMonthly(pagination *PaginatorInput, attendance *Attendance) (*AttendancesResult, error) {
	attendances, err := FetchAttendances(attendance, pagination.BuildPaginator())
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

func CreateOrUpdateAttendance(input *AttendanceInput, query *Attendance) (*AttendanceResult, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	attendance, err := FetchLatestAttendance(query)
	if err != nil {
		return nil, err
	}

	if attendance == nil {
		attendance = new(Attendance)
		attendance.UserId = query.UserId
		if err := CreateAttendance(attendance); err != nil {
			return nil, err
		}
	}

	time := input.BuildAttendanceTime(attendance.Id, attendance.IsClockedOut())

	if err := CreateAttendanceTime(time); err != nil {
		return nil, err
	}

	serializer := NewAttendanceResult(attendance)
	return serializer, nil
}
