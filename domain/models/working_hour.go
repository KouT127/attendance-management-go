package models

import "time"

type WorkingHour struct {
	ID           int64
	StartedAt    time.Time
	FinishedAt   time.Time
	WorkingHours float64
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

func (WorkingHour) TableName() string {
	return "working_hours"
}
