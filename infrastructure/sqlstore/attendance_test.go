package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestCreateAttendanceTime(t *testing.T) {
	InitTestDatabase()
	type args struct {
		ctx            context.Context
		attendanceTime *models.AttendanceTime
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Should succeed creating attendance time",
			args{
				ctx: context.Background(),
				attendanceTime: &models.AttendanceTime{
					Remark:           "test",
					AttendanceKindId: uint8(models.AttendanceKindClockIn),
					IsModified:       false,
					PushedAt:         time.Now(),
					CreatedAt:        time.Now(),
					UpdatedAt:        time.Now(),
				},
			},
			false,
		},
		{
			"Should failed creating attendance time",
			args{
				ctx:            context.Background(),
				attendanceTime: &models.AttendanceTime{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAttendanceTime(tt.args.ctx, tt.args.attendanceTime); (err != nil) != tt.wantErr {
				t.Errorf("CreateAttendanceTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateAttendance(t *testing.T) {
	InitTestDatabase()
	time1 := &models.AttendanceTime{
		Remark:           "test",
		AttendanceKindId: uint8(models.AttendanceKindClockIn),
		IsModified:       false,
		PushedAt:         time.Now(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if err := CreateAttendanceTime(context.Background(), time1); err != nil {
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
			"Should failed creating attendance",
			args{
				ctx:        context.Background(),
				attendance: &models.Attendance{},
			},
			true,
		},
		{
			"Should succeed creating attendance",
			args{
				ctx: context.Background(),
				attendance: &models.Attendance{
					ClockedIn: time1,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAttendance(tt.args.ctx, tt.args.attendance); (err != nil) != tt.wantErr {
				t.Errorf("CreateAttendance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchAttendances(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchAttendances(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchAttendances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchAttendances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchAttendancesCount(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
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
			got, err := FetchAttendancesCount(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchAttendancesCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FetchAttendancesCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchLatestAttendance(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
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
			got, err := FetchLatestAttendance(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchLatestAttendance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchLatestAttendance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOldAttendanceTime(t *testing.T) {
	type args struct {
		ctx    context.Context
		id     int64
		kindId uint8
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
			if err := UpdateOldAttendanceTime(tt.args.ctx, tt.args.id, tt.args.kindId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOldAttendanceTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
