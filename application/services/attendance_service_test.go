package services

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/utilities/timezone"
	"github.com/Songmu/flextime"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	uuid "github.com/satori/go.uuid"
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
		store sqlstore.SQLStore
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
	store := sqlstore.InitTestDatabase()
	timezone.Set("Asia/Tokyo")
	userID := uuid.NewV4().String()

	user := &models.User{
		ID:   userID,
		Name: "test1",
	}

	if err := store.CreateUser(context.Background(), user); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	attendance := &models.Attendance{
		UserID:    user.ID,
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Truncate(time.Second),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Truncate(time.Second),
	}

	if err := store.CreateAttendance(context.Background(), attendance); err != nil {
		t.Errorf("CreateAttendance() failed%s", err)
	}

	time := &models.AttendanceTime{
		Remark:           "test",
		AttendanceKindID: uint8(models.AttendanceKindClockIn),
		IsModified:       false,
		AttendanceID:     attendance.ID,
		PushedAt:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Truncate(time.Second),
		CreatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Truncate(time.Second),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Truncate(time.Second),
	}

	if err := store.CreateAttendanceTime(context.Background(), time); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	attendance.ClockedIn = time
	type fields struct {
		store sqlstore.SQLStore
	}
	type args struct {
		ctx    context.Context
		params models.GetAttendancesParameters
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.GetAttendancesResults
		wantErr bool
	}{
		{
			name: "Should get attendances",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				params: models.GetAttendancesParameters{
					UserID: userID,
					Month:  202001,
				},
			},
			want: &models.GetAttendancesResults{
				MaxCnt: 1,
				Attendances: []*models.Attendance{
					attendance,
				},
			},
			wantErr: false,
		},
		{
			name: "Should not equal max count",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				params: models.GetAttendancesParameters{
					UserID: userID,
					Month:  202002,
				},
			},
			want: &models.GetAttendancesResults{
				MaxCnt:      0,
				Attendances: []*models.Attendance{},
			},
			wantErr: false,
		},
		{
			name: "Should not get attendances",
			fields: fields{
				store: store,
			},
			args: args{
				ctx: context.Background(),
				params: models.GetAttendancesParameters{
					UserID: "",
					Month:  0,
				},
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
			got, err := s.GetAttendances(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttendances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetAttendances() diff %s", diff)
			}
		})
	}
}
