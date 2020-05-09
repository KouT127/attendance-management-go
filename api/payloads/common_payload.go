package payloads

import "github.com/KouT127/attendance-management/domain/models"

type PaginationPayload struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func (i *PaginationPayload) ToPagination() *models.Pagination {
	p := &models.Pagination{}
	return p
}

func NewPaginatorPayload(page int64, limit int64) *PaginationPayload {
	return &PaginationPayload{
		page, limit,
	}
}

func (i *PaginationPayload) CalculatePage() int64 {
	return i.Page * i.Limit
}

func (i *PaginationPayload) HasNext(max int64) bool {
	cnt := i.Page * i.Limit
	return max > cnt
}

type SearchParams struct {
	Date int64 `form:"date"`
}

func NewSearchParams() *SearchParams {
	return &SearchParams{}
}
