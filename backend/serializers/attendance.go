package serializers

import (
	. "github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/utils/timezone"
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

func NewAttendanceSerializer(attendance *Attendance) *AttendanceResult {
	serializer := new(AttendanceResult)
	serializer.NewAttendanceResult(true, attendance)
	return serializer
}

type AttendancesResult struct {
	CommonResponses
	Attendances []*AttendanceResp `json:"attendances"`
}

func (r *AttendanceResp) NewAttendanceResp(a *Attendance) *AttendanceResp {
	loc := timezone.NewJSTLocation()
	r.Id = a.Id
	r.UserId = a.UserId
	r.ClockedInTime = a.ClockedIn
	r.ClockedOutTime = a.ClockedOut
	r.CreatedAt = a.CreatedAt.In(loc).Format("2006-01-02-15:04:05")
	r.UpdatedAt = a.UpdatedAt.In(loc).Format("2006-01-02-15:04:05")
	return r
}

func (s *AttendanceResult) NewAttendanceResult(isSuccessful bool, attendance *Attendance) {
	s.IsSuccessful = true
	s.IsClockedOut = attendance.IsClockedOut()
	res := new(AttendanceResp)
	s.Attendance = *res.NewAttendanceResp(attendance)
}

func (s *AttendancesResult) NewAttendancesResult(isSuccessful bool, hasNext bool, responses []*AttendanceResp) {
	s.IsSuccessful = true
	s.HasNext = hasNext
	s.Attendances = responses
}
