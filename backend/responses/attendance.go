package responses

import . "github.com/KouT127/attendance-management/backend/models"

type AttendanceResponse struct {
	Id             int64          `json:"id"`
	UserId         string         `json:"userId"`
	ClockedInTime  AttendanceTime `json:"clockedInTime"`
	ClockedOutTime AttendanceTime `json:"clockedOutTime"`
	CreatedAt      string         `json:"createdAt"`
	UpdatedAt      string         `json:"updatedAt"`
}
type AttendanceTimeResponse struct {
	PushedAt string `json:"pushedAt"`
	Remark   string `json:"remark"`
}

func (r *AttendanceResponse) Build(a *Attendance) *AttendanceResponse {
	r.Id = a.Id
	r.UserId = a.UserId
	r.CreatedAt = a.CreatedAt.Format("2006-01-02-15:04:05")
	r.UpdatedAt = a.UpdatedAt.Format("2006-01-02-15:04:05")
	return r
}
