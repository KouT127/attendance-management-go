package user

import (
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/utils/logger"
)

func GetOrCreateUser(userId string) (*models.User, error) {
	user, err := models.FetchUser(userId)
	if err != nil {
		return nil, err
	}

	if user.Id == "" {
		user.Id = userId
		_, err := models.CreateUser(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func UpdateUser(user *models.User, userName string) error {
	user.Name = userName
	if _, err := models.UpdateUser(user); err != nil {
		return err
	}
	logger.NewInfo("updated user-" + user.Id)
	return nil
}
