package domains

import "time"

type Attendance struct {
	Id        uint
	UserId    string
	Kind      uint8
	Remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
