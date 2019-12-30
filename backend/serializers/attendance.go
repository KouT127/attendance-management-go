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
	Id             int64          `json:"id"`
	UserId         string         `json:"userId"`
	ClockedInTime  AttendanceTime `json:"clockedInTime"`
	ClockedOutTime AttendanceTime `json:"clockedOutTime"`
	CreatedAt      string         `json:"createdAt"`
	UpdatedAt      string         `json:"updatedAt"`
}

type AttendancesResponse struct {
	CommonResponse
	Attendances []*AttendanceResponse `json:"attendances"`
}

func (r *AttendanceResponse) Build(a *Attendance) *AttendanceResponse {
	loc := timezone.NewJSTLocation()
	r.Id = a.Id
	r.UserId = a.UserId
	r.CreatedAt = a.CreatedAt.In(loc).Format("2006-01-02-15:04:05")
	r.UpdatedAt = a.UpdatedAt.In(loc).Format("2006-01-02-15:04:05")
	return r
}
