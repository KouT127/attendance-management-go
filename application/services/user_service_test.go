package services

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/Songmu/flextime"
	"github.com/google/go-cmp/cmp"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func Test_userService_GetOrCreateUser(t *testing.T) {
	store := sqlstore.InitTestDatabase()
	userID := uuid.NewV4().String()

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
		{
			name: "Should create user",
			fields: fields{
				store: store,
			},
			args: args{
				params: models.GetOrCreateUserParams{
					UserID: userID,
				},
			},
			want: &models.GetOrCreateUserResults{
				User: &models.User{
					ID: userID,
				},
			},
			wantErr: false,
		},
		{
			name: "Should get user",
			fields: fields{
				store: store,
			},
			args: args{
				params: models.GetOrCreateUserParams{
					UserID: userID,
				},
			},
			want: &models.GetOrCreateUserResults{
				User: &models.User{
					ID: userID,
				},
			},
			wantErr: false,
		},
		{
			name: "Should not create user when userId is empty",
			fields: fields{
				store: store,
			},
			args: args{
				params: models.GetOrCreateUserParams{
					UserID: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
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
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetOrCreateUser() diff %s", diff)
			}
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	store := sqlstore.InitTestDatabase()
	userID := uuid.NewV4().String()
	err := store.CreateUser(context.Background(), &models.User{ID: userID})
	if err != nil {
		t.Errorf("CreateUser() %s", err)
	}

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
		want    *models.User
		wantErr bool
	}{
		{
			name: "Should update user",
			fields: fields{
				store: store,
			},
			args: args{
				user: &models.User{
					ID:        userID,
					Name:      "updated",
					Email:     "updated",
					ImageURL:  "updated",
					CreatedAt: flextime.Now(),
					UpdatedAt: flextime.Now(),
				},
			},
			want: &models.User{
				ID:        userID,
				Name:      "updated",
				Email:     "updated",
				ImageURL:  "updated",
				CreatedAt: flextime.Now(),
				UpdatedAt: flextime.Now(),
			},
			wantErr: false,
		},
		{
			name: "Should not update user when user is empty",
			fields: fields{
				store: store,
			},
			args: args{
				user: &models.User{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				store: tt.fields.store,
			}
			if err := s.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := s.GetOrCreateUser(models.GetOrCreateUserParams{UserID: tt.args.user.ID})
			if got == nil {
				// Failed用
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got.User == nil) != tt.wantErr {
				t.Errorf("GetOrCreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got.User, tt.want); diff != "" {
				t.Errorf("UpdateUser() diff %s", diff)
			}
		})
	}
}
