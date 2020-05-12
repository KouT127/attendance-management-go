package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/xerrors"
	"testing"
)

func TestCreateUser(t *testing.T) {
	store := InitTestDatabase()
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Should create user",
			args{
				ctx: context.Background(),
				user: &models.User{
					ID:   "asdiekawei42lasedi356ladfkjfity",
					Name: "test1",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := store.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	store := InitTestDatabase()

	user := &models.User{
		ID:   "asdiekawei42lasedi356ladfkjfity",
		Name: "test1",
	}

	if err := store.CreateUser(context.Background(), user); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Should not create user when have not created user",
			args{
				ctx: context.Background(),
				user: &models.User{
					ID:   "qawsedreftgyhujuiqadnsrt2376sd",
					Name: "test1",
				},
			},
			true,
		},
		{
			"Should create user",
			args{
				ctx: context.Background(),
				user: &models.User{
					ID:       "asdiekawei42lasedi356ladfkjfity",
					Name:     "updatedName",
					Email:    "updatedEmail",
					ImageURL: "updatedImage",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := store.UpdateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if tt.args.user.Name != "updatedName" ||
					tt.args.user.Email != "updatedEmail" ||
					tt.args.user.ImageURL != "updatedImage" {
					t.Errorf("UpdateUser() error = %v, wantErr %v", xerrors.New("Did not updated"), tt.wantErr)
				}
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	store := InitTestDatabase()
	user := &models.User{
		ID:   "asdiekawei42lasedi356ladfkjfity",
		Name: "test1",
	}

	if err := store.CreateUser(context.Background(), user); err != nil {
		t.Errorf("CreateAttendanceTime() failed%s", err)
	}

	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			"Should get user",
			args{
				ctx:    context.Background(),
				userID: "asdiekawei42lasedi356ladfkjfity",
			},
			user,
			false,
		},
		{
			"Should not get user",
			args{
				ctx:    context.Background(),
				userID: "asdiekawei42lasedi356ladfkjfity",
			},
			user,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, IgnoreGlobalOptions); diff != "" {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
