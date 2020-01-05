package validators

type Pagination struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

func NewPagination(page int64, limit int64) *Pagination {
	return &Pagination{
		page, limit,
	}
}

func (p Pagination) CalculatePage() int64 {
	return p.Page * p.Limit
}

func (p Pagination) HasNext(max int64) bool {
	cnt := p.Page * p.Limit
	return max > cnt
}

type SearchParams struct {
	Month int64 `form:"month"`
}

func NewSearchParams() *SearchParams {
	return new(SearchParams)
}
