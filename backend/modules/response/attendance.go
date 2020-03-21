package response

import (
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/timezone"
)

type AttendanceTimeResp struct {
	PushedAt  string `json:"pushedAt"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AttendanceResp struct {
	Id             int64                  `json:"id"`
	UserId         string                 `json:"userId"`
	ClockedInTime  *models.AttendanceTime `json:"clockedInTime"`
	ClockedOutTime *models.AttendanceTime `json:"clockedOutTime"`
	CreatedAt      string                 `json:"createdAt"`
	UpdatedAt      string                 `json:"updatedAt"`
}

type AttendanceResult struct {
	CommonResponse
	Attendance   *AttendanceResp `json:"attendance"`
	IsClockedOut bool            `json:"isClockedOut"`
}

type AttendancesResult struct {
	CommonResponses
	Attendances []*AttendanceResp `json:"attendances"`
}

func toAttendanceResp(a *models.Attendance) *AttendanceResp {
	resp := &AttendanceResp{}
	loc := timezone.NewJSTLocation()
	resp.Id = a.Id
	resp.UserId = a.UserId
	resp.ClockedInTime = a.ClockedIn
	resp.ClockedOutTime = a.ClockedOut
	resp.CreatedAt = a.CreatedAt.In(loc).Format("2006-01-02T15:04:05Z07:00")
	resp.UpdatedAt = a.UpdatedAt.In(loc).Format("2006-01-02T15:04:05Z07:00")
	return resp
}

func ToAttendanceResult(attendance *models.Attendance) *AttendanceResult {
	res := &AttendanceResult{}
	res.IsSuccessful = true
	if attendance != nil {
		res.Attendance = toAttendanceResp(attendance)
	}
	return res
}

func ToAttendancesResult(hasNext bool, attendances []*models.Attendance) *AttendancesResult {
	res := &AttendancesResult{}
	responses := make([]*AttendanceResp, 0)
	for _, attendance := range attendances {
		resp := toAttendanceResp(attendance)
		responses = append(responses, resp)
	}

	res.IsSuccessful = true
	res.HasNext = hasNext
	res.Attendances = responses
	return res
}
