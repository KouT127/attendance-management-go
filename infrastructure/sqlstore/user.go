package sqlstore

import (
	"context"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/modules/logger"
	"golang.org/x/xerrors"
)

func GetUser(ctx context.Context, userId string) (*models.User, error) {
	var user models.User
	err := withDBSession(ctx, func(sess *DBSession) error {
		_, err := sess.Get(&user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(ctx context.Context, user *models.User) error {
	err := withDBSession(ctx, func(sess *DBSession) error {
		if _, err := sess.Insert(user); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(ctx context.Context, user *models.User) error {
	err := withDBSession(ctx, func(sess *DBSession) error {
		has, err := sess.Where("id = ?", user.Id).Exist(&models.User{})
		if err != nil {
			return err
		}
		if !has {
			return xerrors.New("user is empty")
		}

		if _, err := sess.Update(user, &models.User{Id: user.Id}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	logger.NewInfo("updated user_id: " + user.Id)
	return nil
}
