package facades

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"golang.org/x/xerrors"
)

type UserFacade interface {
	GetOrCreateUser(params models.GetOrCreateUserParams) (*models.GetOrCreateUserResults, error)
	UpdateUser(user *models.User) error
}

type userFacade struct {
	ss *sqlstore.SQLStore
}

func NewUserFacade(ss *sqlstore.SQLStore) *userFacade {
	return &userFacade{
		ss: ss,
	}
}

func (f *userFacade) GetOrCreateUser(params models.GetOrCreateUserParams) (*models.GetOrCreateUserResults, error) {
	var (
		user       *models.User
		attendance *models.Attendance
		err        error
	)
	if params.UserId == "" {
		return nil, xerrors.New("user id is empty")
	}

	err = f.ss.InTransaction(context.Background(), func(ctx context.Context) error {
		user, err = sqlstore.GetUser(ctx, params.UserId)
		if err != nil {
			return err
		}

		if user.Id == "" {
			user.Id = params.UserId
			if err = sqlstore.CreateUser(ctx, user); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if attendance, err = sqlstore.FetchLatestAttendance(context.Background(), params.UserId); err != nil {
		return nil, err
	}

	res := models.GetOrCreateUserResults{
		User:             user,
		LatestAttendance: attendance,
	}
	return &res, nil
}

func (f *userFacade) UpdateUser(user *models.User) error {
	err := f.ss.InTransaction(context.Background(), func(ctx context.Context) error {
		return sqlstore.UpdateUser(ctx, user)
	})

	if err != nil {
		return err
	}
	return nil
}
