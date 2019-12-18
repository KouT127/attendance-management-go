package models

import "time"

type Base struct {
	id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
