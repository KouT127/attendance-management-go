package services

import (
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_attendanceService_CreateOrUpdateAttendance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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
