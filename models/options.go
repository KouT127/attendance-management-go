package models

import (
	"xorm.io/xorm"
	"time"
)

type DefaultSearchOption struct {
	Paginator *Paginator
}

func (opt *DefaultSearchOption) setPaginatedSession(eng *xorm.Session) *xorm.Session {
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

type AttendanceSearchOption struct {
	DefaultSearchOption
	UserID string
	Date   *time.Time
}
