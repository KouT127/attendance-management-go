package payloads

import (
	"github.com/KouT127/attendance-management/domain/models"
)

type QueryParam struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (i *QueryParam) ToPagination() *models.Pagination {
	p := &models.Pagination{}
	return p
}

func NewPaginatorPayload(page, limit int) *QueryParam {
	return &QueryParam{
		page, limit,
	}
}

func (i *QueryParam) CalculatePage() int {
	return i.Page * i.Limit
}

func (i *QueryParam) HasNext(max int) bool {
	cnt := i.Page * i.Limit
	return max > cnt
}

type AttendancesQueryParam struct {
	QueryParam
	Month int `form:"month"`
}

func NewAttendancesQueryParam(month int) AttendancesQueryParam {
	return AttendancesQueryParam{
		QueryParam{
			Page:  1,
			Limit: 31,
		},
		month,
	}
}
