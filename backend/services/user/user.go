package user

import (
	"fmt"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/utils/logger"
)

func ViewUser(userId string) (*models.User, *models.Attendance, error) {
	user := new(models.User)
	has, err := models.FetchUser(userId, user)
	if err != nil {
		return nil, nil, err
	}

	if !has {
		user.Id = userId
		_, err := models.CreateUser(user)
		if err != nil {
			return nil, nil, err
		}
	}

	attendance, err := models.FetchLatestAttendance(userId)
	fmt.Println(attendance)
	if err != nil {
		return nil, nil, err
	}
	return user, attendance, nil
}

func UpdateUser(userId string, userName string) (*models.User, error) {
	u := new(models.User)
	has, err := models.FetchUser(userId, u)
	if err != nil || !has {
		return nil, err
	}
	u.Name = userName
	_, err = models.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	logger.NewInfo("updated user-" + u.Id)

	return u, nil
}
