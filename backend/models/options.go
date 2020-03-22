package models

import (
	"github.com/go-xorm/xorm"
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
	UserId string
	Date   *time.Time
}

func (opt *AttendanceSearchOption) setQueriedSession(eng Engine) Engine {
	return eng.
		Where("attendances.user_id = ?", opt.UserId)
}
