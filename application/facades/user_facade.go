package facades

import (
	"context"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/models"
	"golang.org/x/xerrors"
)

type UserFacade interface {
	GetOrCreateUser(userID string) (*models.User, error)
	UpdateUser(user *models.User) error
}

type userFacade struct {
	ss sqlstore.SQLStore
}

func NewUserFacade(ss sqlstore.SQLStore) userFacade {
	return userFacade{
		ss: ss,
	}
}

func (f *userFacade) GetOrCreateUser(userID string) (*models.User, error) {
	var (
		user *models.User
		err  error
	)
	if userID == "" {
		return nil, xerrors.New("user id is empty")
	}

	err = f.ss.InTransaction(context.Background(), func(ctx context.Context) error {
		user, err = sqlstore.GetUser(ctx, userID)
		if err != nil {
			return err
		}

		if user.ID == "" {
			user.ID = userID
			if err = sqlstore.CreateUser(ctx, user); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return user, nil
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
