package responses

import (
	. "github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/timezone"
)

type AttendanceTimeResp struct {
	PushedAt  string `json:"pushedAt"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AttendanceResp struct {
	Id             int64           `json:"id"`
	UserId         string          `json:"userId"`
	ClockedInTime  *AttendanceTime `json:"clockedInTime"`
	ClockedOutTime *AttendanceTime `json:"clockedOutTime"`
	CreatedAt      string          `json:"createdAt"`
	UpdatedAt      string          `json:"updatedAt"`
}

type AttendanceResult struct {
	CommonResponse
	Attendance   AttendanceResp `json:"attendance"`
	IsClockedOut bool           `json:"isClockedOut"`
}

func NewAttendanceResult(attendance *Attendance) *AttendanceResult {
	serializer := new(AttendanceResult)
	serializer.NewAttendanceResult(true, attendance)
	return serializer
}

type AttendancesResult struct {
	CommonResponses
	Attendances []*AttendanceResp `json:"attendances"`
}

func NewAttendanceResp(a *Attendance) AttendanceResp {
	resp := AttendanceResp{}
	loc := timezone.NewJSTLocation()
	resp.Id = a.Id
	resp.UserId = a.UserId
	resp.ClockedInTime = a.ClockedIn
	resp.ClockedOutTime = a.ClockedOut
	resp.CreatedAt = a.CreatedAt.In(loc).Format("2006-01-02-15:04:05")
	resp.UpdatedAt = a.UpdatedAt.In(loc).Format("2006-01-02-15:04:05")
	return resp
}

func (s *AttendanceResult) NewAttendanceResult(isSuccessful bool, attendance *Attendance) {
	s.IsSuccessful = true
	if attendance != nil {
		s.Attendance = NewAttendanceResp(attendance)
	}
}

func (s *AttendancesResult) NewAttendancesResult(isSuccessful bool, hasNext bool, responses []*AttendanceResp) {
	s.IsSuccessful = true
	s.HasNext = hasNext
	s.Attendances = responses
}
