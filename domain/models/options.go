package models

import (
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type DefaultSearchOption struct {
	Paginator *Pagination
}

func (opt *DefaultSearchOption) SetPaginatedSession(eng *xorm.Session) *xorm.Session {
	p := opt.Paginator
	if opt.Paginator == nil {
		p = &Pagination{}
	}
	if p.Limit == 0 {
		p.Limit = 15
	}
	page := p.CalculatePage()

	return eng.Limit(int(p.Limit), int(page))
}

type GetAttendancesParameters struct {
	UserID string
	Month  int
}

func (p GetAttendancesParameters) Validate() error {
	if p.UserID == "" {
		return xerrors.New("user id is empty")
	}
	if p.Month == 0 {
		return xerrors.New("month is zero")
	}
	return nil
}

type GetAttendanceSummaryParameters struct {
	UserID string
}

type GetAttendancesResults struct {
	MaxCnt      int64
	Attendances []*Attendance
}

type GetAttendanceSummaryResults struct {
	LatestAttendance Attendance
	TotalHours       float64
	RequiredHours    float64
}
