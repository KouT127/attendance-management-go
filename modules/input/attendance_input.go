package validators

import (
	. "github.com/KouT127/attendance-management/models"
	"github.com/Songmu/flextime"
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type AttendanceInput struct {
	Remark string
}

func (i AttendanceInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Remark, validation.Length(0, 1000)),
	)
}

func (i AttendanceInput) ToAttendanceTime() *AttendanceTime {
	t := new(AttendanceTime)
	t.Remark = i.Remark
	t.PushedAt = flextime.Now()
	t.CreatedAt = flextime.Now()
	t.UpdatedAt = flextime.Now()
	return t
}
