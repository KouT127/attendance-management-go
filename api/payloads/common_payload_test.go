package payloads

import (
	"github.com/KouT127/attendance-management/domain/models"
	"reflect"
	"testing"
)

func TestNewPaginatorPayload(t *testing.T) {
	type args struct {
		page  int
		limit int
	}
	tests := []struct {
		name string
		args args
		want *QueryParam
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
		want *AttendancesQueryParam
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttendancesQueryParam(202001); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttendancesQueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaginationPayload_CalculatePage(t *testing.T) {
	type fields struct {
		Page  int
		Limit int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &QueryParam{
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
		Page  int
		Limit int
	}
	type args struct {
		max int
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
			i := &QueryParam{
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
		Page  int
		Limit int
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Pagination
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &QueryParam{
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if got := i.ToPagination(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}
