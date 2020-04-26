package payloads

import (
	"github.com/KouT127/attendance-management/domain/models"
	"reflect"
	"testing"
)

func TestNewPaginatorPayload(t *testing.T) {
	type args struct {
		page  int64
		limit int64
	}
	tests := []struct {
		name string
		args args
		want *PaginationPayload
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaginatorPayload(tt.args.page, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaginatorPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSearchParams(t *testing.T) {
	tests := []struct {
		name string
		want *SearchParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaginationPayload_CalculatePage(t *testing.T) {
	type fields struct {
		Page  int64
		Limit int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PaginationPayload{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := i.CalculatePage(); got != tt.want {
				t.Errorf("CalculatePage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaginationPayload_HasNext(t *testing.T) {
	type fields struct {
		Page  int64
		Limit int64
	}
	type args struct {
		max int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PaginationPayload{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := i.HasNext(tt.args.max); got != tt.want {
				t.Errorf("HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaginationPayload_ToPaginator(t *testing.T) {
	type fields struct {
		Page  int64
		Limit int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Paginator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PaginationPayload{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := i.ToPaginator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPaginator() = %v, want %v", got, tt.want)
			}
		})
	}
}
