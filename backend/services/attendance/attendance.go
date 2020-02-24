package attendance

import (
	. "github.com/KouT127/attendance-management/models"
	. "github.com/KouT127/attendance-management/responses"
	. "github.com/KouT127/attendance-management/validators"
)

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
