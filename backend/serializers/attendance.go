package serializers

import (
	. "github.com/KouT127/attendance-management/backend/models"
	"github.com/KouT127/attendance-management/backend/utils/timezone"
)

type AttendanceTimeResponse struct {
	PushedAt  string `json:"pushedAt"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AttendanceResponse struct {
	Id             int64           `json:"id"`
	UserId         string          `json:"userId"`
	ClockedInTime  *AttendanceTime `json:"clockedInTime"`
	ClockedOutTime *AttendanceTime `json:"clockedOutTime"`
	CreatedAt      string          `json:"createdAt"`
	UpdatedAt      string          `json:"updatedAt"`
}

type AttendanceSerializer struct {
	CommonResponse
	Attendance AttendanceResponse `json:"attendance"`
}

type AttendancesSerializer struct {
	CommonResponse
	Attendances []*AttendanceResponse `json:"attendances"`
}

func (r *AttendanceResponse) Build(a *Attendance) *AttendanceResponse {
	loc := timezone.NewJSTLocation()
	r.Id = a.Id
	r.UserId = a.UserId
	r.ClockedInTime = a.ClockedIn
	r.ClockedOutTime = a.ClockedOut
	r.CreatedAt = a.CreatedAt.In(loc).Format("2006-01-02-15:04:05")
	r.UpdatedAt = a.UpdatedAt.In(loc).Format("2006-01-02-15:04:05")
	return r
}

func (s *AttendanceSerializer) Serialize(isSuccessful bool, hasNext bool, attendance *Attendance) {
	s.IsSuccessful = true
	s.HasNext = hasNext
	res := new(AttendanceResponse)
	s.Attendance = *res.Build(attendance)
}

func (s *AttendancesSerializer) Serialize(isSuccessful bool, hasNext bool, responses []*AttendanceResponse) {
	s.IsSuccessful = true
	s.HasNext = hasNext
	s.Attendances = responses
}
