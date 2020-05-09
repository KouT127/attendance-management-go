package responses

import (
	"github.com/KouT127/attendance-management/domain/models"
	"time"
)

type AttendanceTimeResp struct {
	ID               int64  `json:"id"`
	AttendanceID     int64  `json:"attendance_id"`
	AttendanceKindID uint8  `json:"attendance_kind_id"`
	IsModified       bool   `json:"is_modified"`
	PushedAt         string `json:"pushed_at"`
	Remark           string `json:"remark"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type AttendanceResp struct {
	ID             int64               `json:"id"`
	UserID         string              `json:"user_id"`
	ClockedInTime  *AttendanceTimeResp `json:"clocked_in_time"`
	ClockedOutTime *AttendanceTimeResp `json:"clocked_out_time"`
	CreatedAt      string              `json:"created_at"`
	UpdatedAt      string              `json:"updated_at"`
}

type AttendanceResult struct {
	CommonResponse
	Attendance   *AttendanceResp `json:"attendance"`
	IsClockedOut bool            `json:"is_clocked_out"`
}

type AttendancesResponses struct {
	CommonResponse
	Attendances []*AttendanceResp `json:"attendances"`
}

func toAttendanceResp(a *models.Attendance) *AttendanceResp {
	resp := &AttendanceResp{}
	resp.ID = a.ID
	resp.UserID = a.UserID
	if a.ClockedIn != nil {
		resp.ClockedInTime = toAttendanceTimeResp(a.ClockedIn)
	}
	if a.ClockedOut != nil {
		resp.ClockedOutTime = toAttendanceTimeResp(a.ClockedOut)
	}
	resp.CreatedAt = a.CreatedAt.Format(time.RFC3339)
	resp.UpdatedAt = a.UpdatedAt.Format(time.RFC3339)
	return resp
}

func toAttendanceTimeResp(t *models.AttendanceTime) *AttendanceTimeResp {
	return &AttendanceTimeResp{
		ID:               t.ID,
		AttendanceID:     t.AttendanceID,
		AttendanceKindID: t.AttendanceKindID,
		IsModified:       t.IsModified,
		PushedAt:         t.PushedAt.Format(time.RFC3339),
		Remark:           t.Remark,
		CreatedAt:        t.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        t.UpdatedAt.Format(time.RFC3339),
	}
}

func ToAttendanceResult(attendance *models.Attendance) *AttendanceResult {
	res := &AttendanceResult{}
	res.IsSuccessful = true
	if attendance != nil {
		res.Attendance = toAttendanceResp(attendance)
	}
	return res
}

func ToAttendancesResponses(hasNext bool, attendances []*models.Attendance) *AttendancesResponses {
	res := &AttendancesResponses{}
	responses := make([]*AttendanceResp, 0)
	for _, attendance := range attendances {
		resp := toAttendanceResp(attendance)
		responses = append(responses, resp)
	}

	res.IsSuccessful = true
	res.Attendances = responses
	return res
}
