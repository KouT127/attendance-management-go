package models

type Pagination struct {
	Page  int64
	Limit int64
}

func (p *Pagination) CalculatePage() int64 {
	return p.Page * p.Limit
}

func (p *Pagination) HasNext(max int64) bool {
	cnt := p.Page * p.Limit
	return max > cnt
}
