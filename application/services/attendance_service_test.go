package services

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	sqlstore "github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/utilities/timezone"
	"github.com/Songmu/flextime"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
	"time"
)

func Test_attendanceService_CreateOrUpdateAttendance(t *testing.T) {
	store := sqlstore.InitTestDatabase()
	timezone.Set("Asia/Tokyo")
	flextime.Fix(time.Date(2020, 1, 1, 0, 0, 0, 0, timezone.JSTLocation()))
	options := cmp.Options{
		cmpopts.IgnoreFields(models.Attendance{}, "ID"),
		cmpopts.IgnoreFields(models.AttendanceTime{}, "ID", "AttendanceID"),
	}

	userID := uuid.NewV4().String()
	if err := store.CreateUser(context.Background(), &models.User{
		ID:        userID,
		Name:      "insert user",
		Email:     "insert",
		ImageURL:  "insert",
		CreatedAt: flextime.Now(),
		UpdatedAt: flextime.Now(),
	}); err != nil {
		t.Errorf("CreateUser() %s", err)
	}

	type fields struct {
		store sqlstore.SqlStore
	}
	type args struct {
		ctx            context.Context
		attendanceTime *models.AttendanceTime
		userID         string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		want             *models.Attendance
		shouldChangeDate bool
		wantErr          bool
	}{
		{
			name: "Should check in",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:     "test",
					IsModified: false,
				},
				userID: userID,
			},
			want: &models.Attendance{
				UserID:    userID,
				CreatedAt: flextime.Now(),
				UpdatedAt: flextime.Now(),
				ClockedIn: &models.AttendanceTime{
					Remark:           "test",
					IsModified:       false,
					AttendanceKindID: uint8(models.AttendanceKindClockIn),
					PushedAt:         flextime.Now(),
					CreatedAt:        flextime.Now(),
					UpdatedAt:        flextime.Now(),
				},
				ClockedOut: nil,
			},
			wantErr: false,
		},
		{
			name: "Should check out",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:     "test1",
					IsModified: false,
				},
				userID: userID,
			},
			want: &models.Attendance{
				UserID:    userID,
				CreatedAt: flextime.Now(),
				UpdatedAt: flextime.Now(),
				ClockedIn: &models.AttendanceTime{
					Remark:           "test",
					IsModified:       false,
					AttendanceKindID: uint8(models.AttendanceKindClockIn),
					PushedAt:         flextime.Now(),
					CreatedAt:        flextime.Now(),
					UpdatedAt:        flextime.Now(),
				},
				ClockedOut: &models.AttendanceTime{
					Remark:           "test1",
					IsModified:       false,
					AttendanceKindID: uint8(models.AttendanceKindClockOut),
					PushedAt:         flextime.Now(),
					CreatedAt:        flextime.Now(),
					UpdatedAt:        flextime.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "Should check out when second time",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:     "test2",
					IsModified: false,
				},
				userID: userID,
			},
			want: &models.Attendance{
				UserID:    userID,
				CreatedAt: flextime.Now(),
				UpdatedAt: flextime.Now(),
				ClockedIn: &models.AttendanceTime{
					Remark:           "test",
					IsModified:       false,
					AttendanceKindID: uint8(models.AttendanceKindClockIn),
					PushedAt:         flextime.Now(),
					CreatedAt:        flextime.Now(),
					UpdatedAt:        flextime.Now(),
				},
				ClockedOut: &models.AttendanceTime{
					Remark:           "test2",
					IsModified:       false,
					AttendanceKindID: uint8(models.AttendanceKindClockOut),
					PushedAt:         flextime.Now(),
					CreatedAt:        flextime.Now(),
					UpdatedAt:        flextime.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "Should check out when dates changes",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:     "test",
					IsModified: false,
				},
				userID: userID,
			},
			want: &models.Attendance{
				UserID:    userID,
				CreatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, timezone.JSTLocation()),
				UpdatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, timezone.JSTLocation()),
				ClockedIn: &models.AttendanceTime{
					Remark:           "test",
					IsModified:       false,
					AttendanceKindID: uint8(models.AttendanceKindClockIn),
					PushedAt:         time.Date(2020, 1, 2, 0, 0, 0, 0, timezone.JSTLocation()),
					CreatedAt:        time.Date(2020, 1, 2, 0, 0, 0, 0, timezone.JSTLocation()),
					UpdatedAt:        time.Date(2020, 1, 2, 0, 0, 0, 0, timezone.JSTLocation()),
				},
				ClockedOut: nil,
			},
			shouldChangeDate: true,
			wantErr:          false,
		},
		{
			name: "Should not create attendance when userID is empty",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:     "test",
					IsModified: false,
				},
				userID: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should not create attendance when time is empty",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:            context.Background(),
				attendanceTime: nil,
				userID:         userID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &attendanceService{
				store: tt.fields.store,
			}
			if tt.shouldChangeDate {
				flextime.Fix(time.Date(2020, 1, 2, 0, 0, 0, 0, timezone.JSTLocation()))
			}
			got, err := s.CreateOrUpdateAttendance(tt.args.ctx, tt.args.attendanceTime, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateAttendance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, options); diff != "" {
				t.Errorf("CreateOrUpdateAttendance() diff + want - got %v", diff)
			}
		})
	}
}

func Test_attendanceService_GetAttendances(t *testing.T) {
	type fields struct {
		store sqlstore.SqlStore
	}
	type args struct {
		params models.GetAttendancesParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.GetAttendancesResults
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &attendanceService{
				store: tt.fields.store,
			}
			got, err := s.GetAttendances(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttendances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAttendances() got = %v, want %v", got, tt.want)
			}
		})
	}
}
