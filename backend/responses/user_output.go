package responses

import . "github.com/KouT127/Attendance-management/backend/domains"

type AttendanceResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"userId"`
	Kind      uint8  `json:"kind"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (r *AttendanceResponse) SetAttendance(a *Attendance) *AttendanceResponse {
	r.Id = a.Id
	r.UserId = a.UserId
	r.Kind = a.Kind
	r.Remark = a.Remark
	r.CreatedAt = a.CreatedAt.Format("2006-01-02")
	r.UpdatedAt = a.UpdatedAt.Format("2006-01-02")
	return r
}
