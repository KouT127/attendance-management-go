package domains

import "time"

type Attendance struct {
	Id        uint
	UserId    string
	Kind      uint8
	Remark    string
	PushedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
