package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/Songmu/flextime"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
	"xorm.io/xorm"
)

func Test_sqlStore_CreateWorkingHours(t *testing.T) {
	InitTestDatabase()

	type fields struct {
		engine *xorm.Engine
	}
	type args struct {
		ctx  context.Context
		hour *models.WorkingHour
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should create working hours",
			fields: fields{
				engine: eng,
			},
			args: args{
				ctx: context.Background(),
				hour: &models.WorkingHour{
					StartedAt:    flextime.Now(),
					FinishedAt:   flextime.Now(),
					WorkingHours: 180,
				},
			},
			wantErr: false,
		},
		{
			name: "Should not create working hours",
			fields: fields{
				engine: eng,
			},
			args: args{
				ctx: context.Background(),
				hour: &models.WorkingHour{
					StartedAt: flextime.Now(),
				},
			},
			wantErr: true,
		},
		{
			name: "Should not create working hours",
			fields: fields{
				engine: eng,
			},
			args: args{
				ctx: context.Background(),
				hour: &models.WorkingHour{
					FinishedAt: flextime.Now(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := sqlStore{
				engine: tt.fields.engine,
			}
			if err := sq.CreateWorkingHour(tt.args.ctx, tt.args.hour); (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkingHour() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqlStore_GetWorkingHour(t *testing.T) {
	store := InitTestDatabase()
	wh := &models.WorkingHour{
		StartedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		FinishedAt:   time.Date(2020, 1, 30, 0, 0, 0, 0, time.UTC),
		WorkingHours: 180,
	}
	if err := store.CreateWorkingHour(context.Background(), wh); err != nil {
		t.Errorf("CreateWorkingHour() err = %v", err)
	}

	type fields struct {
		engine *xorm.Engine
	}
	type args struct {
		ctx context.Context
		now time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.WorkingHour
		wantErr bool
	}{
		{
			name: "Should get working hours",
			fields: fields{
				engine: eng,
			},
			args: args{
				ctx: context.Background(),
				now: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want:    wh,
			wantErr: false,
		},
		{
			name: "Should not get working hours",
			fields: fields{
				engine: eng,
			},
			args: args{
				ctx: context.Background(),
				now: time.Date(2020, 2, 2, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := sqlStore{
				engine: tt.fields.engine,
			}
			got, err := sq.GetWorkingHours(tt.args.ctx, tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkingHours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, IgnoreGlobalOptions); diff != "" {
				t.Errorf("GetWorkingHours() diff %s", diff)
			}
		})
	}
}
