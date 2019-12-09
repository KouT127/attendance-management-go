package validators

type AttendanceInput struct {
	Kind   uint8
	Remark string
}

//func (a AttendanceInput) Validate() error {
//	return validation.ValidateStruct(&a,
//		validation.Field(&a.UserId, validation.Required, validation.Length(1, 50)),
//		validation.Field(&a.Kind, validation.Required, validation.Length(1, 50)),
//		validation.Field(&a.Remark, validation.Length(0, 1000)),
//	)
//}
