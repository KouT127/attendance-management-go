package payloads

import (
	"github.com/KouT127/attendance-management/models"
	"github.com/Songmu/flextime"
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type AttendancePayload struct {
	Remark string
}

func (i *AttendancePayload) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Remark, validation.Length(0, 1000)),
	)
}

func (i *AttendancePayload) ToAttendanceTime() *models.AttendanceTime {
	t := &models.AttendanceTime{}
	t.Remark = i.Remark
	t.PushedAt = flextime.Now()
	t.CreatedAt = flextime.Now()
	t.UpdatedAt = flextime.Now()
	return t
}
