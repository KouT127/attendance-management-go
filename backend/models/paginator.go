package models

type Paginator struct {
	Page  int64
	Limit int64
}

func (p Paginator) CalculatePage() int64 {
	return p.Page * p.Limit
}

func (p Paginator) HasNext(max int64) bool {
	cnt := p.Page * p.Limit
	return max > cnt
}
