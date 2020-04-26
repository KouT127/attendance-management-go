package models

import (
	"time"
	"xorm.io/xorm"
)

type DefaultSearchOption struct {
	Paginator *Paginator
}

func (opt *DefaultSearchOption) SetPaginatedSession(eng *xorm.Session) *xorm.Session {
	p := opt.Paginator
	if opt.Paginator == nil {
		p = &Paginator{}
	}
	if p.Limit == 0 {
		p.Limit = 15
	}
	page := p.CalculatePage()

	return eng.Limit(int(p.Limit), int(page))
}

type GetAttendancesParameters struct {
	UserId    string
	Date      *time.Time
	Paginator *Paginator
}

type GetAttendancesResults struct {
	MaxCnt      int64
	Attendances []*Attendance
}
