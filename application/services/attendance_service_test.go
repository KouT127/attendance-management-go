package services

import (
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"reflect"
	"testing"
)

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
