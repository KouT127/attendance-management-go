package domains

type Attendance struct {
	Id     int64
	UserId string
	Kind   uint8
	Remark string
}
