package payloads

import (
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/Songmu/flextime"
	"reflect"
	"testing"
	"time"
)

func TestAttendancePayload_ToAttendanceTime(t *testing.T) {
	flextime.Fix(time.Date(2020, 01, 01, 1, 1, 1, 1, time.Local))

	type fields struct {
		Remark string
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.AttendanceTime
	}{
		{
			name: "Should convert attendance payload to attendance time",
			fields: fields{
				Remark: "remark",
			},
			want: &models.AttendanceTime{
				Remark:    "remark",
				PushedAt:  flextime.Now(),
				CreatedAt: flextime.Now(),
				UpdatedAt: flextime.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AttendancePayload{
				Remark: tt.fields.Remark,
			}
			if got := i.ToAttendanceTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToAttendanceTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttendancePayload_Validate(t *testing.T) {
	type fields struct {
		Remark string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Should validate",
			fields: fields{
				Remark: "test",
			},
			wantErr: false,
		},
		{
			name: "Should not validate",
			fields: fields{
				Remark: "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AttendancePayload{
				Remark: tt.fields.Remark,
			}
			if err := i.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
