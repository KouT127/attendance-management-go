package responses

import (
	"github.com/KouT127/attendance-management/domain/models"
	"time"
)

type AttendanceTimeResponse struct {
	ID               int64  `json:"id"`
	AttendanceID     int64  `json:"attendance_id"`
	AttendanceKindID uint8  `json:"attendance_kind_id"`
	IsModified       bool   `json:"is_modified"`
	PushedAt         string `json:"pushed_at"`
	Remark           string `json:"remark"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type AttendanceResponse struct {
	ID             int64                   `json:"id"`
	UserID         string                  `json:"user_id"`
	ClockedInTime  *AttendanceTimeResponse `json:"clocked_in_time"`
	ClockedOutTime *AttendanceTimeResponse `json:"clocked_out_time"`
	CreatedAt      string                  `json:"created_at"`
	UpdatedAt      string                  `json:"updated_at"`
}

type AttendanceCreatedResponse struct {
	CommonResponse
	Attendance   *AttendanceResponse `json:"attendance"`
	IsClockedOut bool                `json:"is_clocked_out"`
}

type AttendancesResponses struct {
	CommonResponse
	Attendances []*AttendanceResponse `json:"attendances"`
}

type AttendanceSummaryResponse struct {
	CommonResponse
	LatestAttendance *models.Attendance `json:"latest_attendance"`
	RequiredHours    float64            `json:"required_time"`
	TotalHours       float64            `json:"total_time"`
}

func toAttendanceResponse(a *models.Attendance) *AttendanceResponse {
	resp := &AttendanceResponse{}
	resp.ID = a.ID
	resp.UserID = a.UserID
	if a.ClockedIn != nil {
		resp.ClockedInTime = toAttendanceTimeResponse(a.ClockedIn)
	}
	if a.ClockedOut != nil {
		resp.ClockedOutTime = toAttendanceTimeResponse(a.ClockedOut)
	}
	resp.CreatedAt = a.CreatedAt.Format(time.RFC3339)
	resp.UpdatedAt = a.UpdatedAt.Format(time.RFC3339)
	return resp
}

func toAttendanceTimeResponse(t *models.AttendanceTime) *AttendanceTimeResponse {
	return &AttendanceTimeResponse{
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

func ToAttendanceCreatedResponse(attendance *models.Attendance) *AttendanceCreatedResponse {
	res := &AttendanceCreatedResponse{}
	res.IsSuccessful = true
	if attendance != nil {
		res.Attendance = toAttendanceResponse(attendance)
	}
	return res
}

func ToAttendancesResponses(attendances []*models.Attendance) *AttendancesResponses {
	res := &AttendancesResponses{}
	responses := make([]*AttendanceResponse, 0)
	for _, attendance := range attendances {
		resp := toAttendanceResponse(attendance)
		responses = append(responses, resp)
	}

	res.IsSuccessful = true
	res.Attendances = responses
	return res
}

func ToAttendanceSummaryResponse(results *models.GetAttendanceSummaryResults) *AttendanceSummaryResponse {
	res := AttendanceSummaryResponse{
		TotalHours:    results.TotalHours,
		RequiredHours: results.RequiredHours,
	}
	if results.LatestAttendance.ID != 0 {
		res.LatestAttendance = &results.LatestAttendance
	}
	res.IsSuccessful = true
	return &res
}
