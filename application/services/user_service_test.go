package services

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewAttendanceService(t *testing.T) {
	type args struct {
		ss sqlstore.SqlStore
	}
	tests := []struct {
		name string
		args args
		want AttendanceService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttendanceService(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttendanceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	type args struct {
		ss sqlstore.SqlStore
	}
	tests := []struct {
		name string
		args args
		want *userService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attendanceService_CreateOrUpdateAttendance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := sqlstore.NewMockAttendance(ctrl)
	store.EXPECT().CreateAttendanceTime(context.Background(), &models.AttendanceTime{})
	_ = store.CreateAttendanceTime(context.Background(), &models.AttendanceTime{})

	type fields struct {
		store sqlstore.SqlStore
	}
	type args struct {
		attendanceTime *models.AttendanceTime
		userId         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Attendance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &attendanceService{
				store: tt.fields.store,
			}
			got, err := s.CreateOrUpdateAttendance(tt.args.attendanceTime, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateAttendance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrUpdateAttendance() got = %v, want %v", got, tt.want)
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

func Test_userService_GetOrCreateUser(t *testing.T) {
	type fields struct {
		store sqlstore.SqlStore
	}
	type args struct {
		params models.GetOrCreateUserParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.GetOrCreateUserResults
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				store: tt.fields.store,
			}
			got, err := s.GetOrCreateUser(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrCreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	type fields struct {
		store sqlstore.SqlStore
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				store: tt.fields.store,
			}
			if err := s.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
