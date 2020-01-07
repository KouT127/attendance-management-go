package usecases

import . "github.com/KouT127/attendance-management/backend/repositories"

type PaginatorInput struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func (i PaginatorInput) BuildPaginator() *Paginator {
	p := new(Paginator)
	return p
}

func NewPaginatorInput(page int64, limit int64) *PaginatorInput {
	return &PaginatorInput{
		page, limit,
	}
}

func (p PaginatorInput) CalculatePage() int64 {
	return p.Page * p.Limit
}

func (p PaginatorInput) HasNext(max int64) bool {
	cnt := p.Page * p.Limit
	return max > cnt
}

type SearchParams struct {
	Month int64 `form:"month"`
}

func NewSearchParams() *SearchParams {
	return new(SearchParams)
}
