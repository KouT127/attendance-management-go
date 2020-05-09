package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/utilities/timezone"
	"github.com/Songmu/flextime"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
	"time"
)

func TestCreateAttendance(t *testing.T) {
	store := InitTestDatabase()

	user := &models.User{
		ID:   "asdiekawei42lasedi356ladfkjfity",
		Name: "test1",
	}

	if err := store.CreateUser(context.Background(), user); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	type args struct {
		ctx        context.Context
		attendance *models.Attendance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Should not create attendance",
			args{
				ctx:        context.Background(),
				attendance: &models.Attendance{},
			},
			true,
		},
		{
			"Should create clocked in attendance",
			args{
				ctx: context.Background(),
				attendance: &models.Attendance{
					UserID: user.ID,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := store.CreateAttendance(tt.args.ctx, tt.args.attendance); (err != nil) != tt.wantErr {
				t.Errorf("CreateAttendance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateAttendanceTime(t *testing.T) {
	store := InitTestDatabase()
	type args struct {
		ctx            context.Context
		attendanceTime *models.AttendanceTime
	}
	user := &models.User{
		ID:   "asdiekawei42lasedi356ladfkjfity",
		Name: "test1",
	}

	if err := store.CreateUser(context.Background(), user); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	attendance := &models.Attendance{
		UserID: user.ID,
	}

	if err := store.CreateAttendance(context.Background(), attendance); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Should create attendance time",
			args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:           "test",
					AttendanceKindID: uint8(models.AttendanceKindClockIn),
					IsModified:       false,
					AttendanceID:     attendance.ID,
					PushedAt:         flextime.Now(),
					CreatedAt:        flextime.Now(),
					UpdatedAt:        flextime.Now(),
				},
			},
			false,
		},
		{
			"Should not create attendance time",
			args{
				ctx:            context.Background(),
				attendanceTime: &models.AttendanceTime{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := store.CreateAttendanceTime(tt.args.ctx, tt.args.attendanceTime); (err != nil) != tt.wantErr {
				t.Errorf("CreateAttendanceTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchAttendances(t *testing.T) {
	store := InitTestDatabase()
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

	type args struct {
		ctx   context.Context
		query *models.GetAttendancesParameters
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Attendance
		wantErr bool
	}{
		{
			name: "Should get attendances",
			args: args{
				ctx: context.Background(),
				query: &models.GetAttendancesParameters{
					UserID: userID,
					Month:  0,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should get attendances",
			args: args{
				ctx: context.Background(),
				query: &models.GetAttendancesParameters{
					UserID: "",
					Month:  202001,
				},
			},
			want:    []*models.Attendance{},
			wantErr: false,
		},
		{
			name: "Should get attendances by params",
			args: args{
				ctx: context.Background(),
				query: &models.GetAttendancesParameters{
					UserID: userID,
					Month:  202001,
				},
			},
			want: []*models.Attendance{
				attendance,
			},
			wantErr: false,
		},
		{
			name: "Should get attendances by params",
			args: args{
				ctx: context.Background(),
				query: &models.GetAttendancesParameters{
					UserID: userID,
					Month:  202002,
				},
			},
			want:    []*models.Attendance{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetAttendances(tt.args.ctx, tt.args.query)
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

func TestFetchAttendancesCount(t *testing.T) {
	store := InitTestDatabase()

	type args struct {
		ctx   context.Context
		query *models.GetAttendancesParameters
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetAttendancesCount(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttendancesCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAttendancesCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchLatestAttendance(t *testing.T) {
	store := InitTestDatabase()
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetLatestAttendance(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestAttendance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLatestAttendance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOldAttendanceTime(t *testing.T) {
	store := InitTestDatabase()
	type args struct {
		ctx    context.Context
		id     int64
		kindID uint8
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := store.UpdateOldAttendanceTime(tt.args.ctx, tt.args.id, tt.args.kindID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOldAttendanceTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
