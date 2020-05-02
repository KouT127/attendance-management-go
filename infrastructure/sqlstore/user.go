package sqlstore

//go:generate mockgen -source=user.go -destination=mock/mock_user.go -package=sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/utilities/logger"
	"golang.org/x/xerrors"
)

type User interface {
	GetUser(ctx context.Context, userId string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
}

func (sqlStore) GetUser(ctx context.Context, userId string) (*models.User, error) {
	sess, err := getDBSession(ctx)
	if err != nil {
		return nil, err
	}

	user := &models.User{Id: userId}
	_, err = sess.Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (sqlStore) CreateUser(ctx context.Context, user *models.User) error {
	sess, err := getDBSession(ctx)
	if err != nil {
		return err
	}

	if _, err := sess.Insert(user); err != nil {
		return err
	}
	return nil
}

func (sqlStore) UpdateUser(ctx context.Context, user *models.User) error {
	sess, err := getDBSession(ctx)
	if err != nil {
		return err
	}

	has, err := sess.Where("id = ?", user.Id).Exist(&models.User{})
	if err != nil {
		return err
	}
	if !has {
		return xerrors.New("user is not exists")
	}

	if _, err := sess.Update(user, &models.User{Id: user.Id}); err != nil {
		return err
	}
	logger.NewInfo("updated user_id: " + user.Id)

	return nil
}
