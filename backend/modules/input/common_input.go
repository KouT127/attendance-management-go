package validators

import (
	. "github.com/KouT127/attendance-management/models"
)

type PaginatorInput struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func (i *PaginatorInput) BuildPaginator() *Paginator {
	p := new(Paginator)
	return p
}

func NewPaginatorInput(page int64, limit int64) *PaginatorInput {
	return &PaginatorInput{
		page, limit,
	}
}

func (i *PaginatorInput) CalculatePage() int64 {
	return i.Page * i.Limit
}

func (i *PaginatorInput) HasNext(max int64) bool {
	cnt := i.Page * i.Limit
	return max > cnt
}

type SearchParams struct {
	Month int64 `form:"month"`
}

func NewSearchParams() *SearchParams {
	return new(SearchParams)
}
