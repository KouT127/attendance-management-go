package models

import (
	"errors"
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/modules/logger"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return database.UserTable
}

func getUser(eng Engine, userId string) (*User, error) {
	u := User{Id: userId}
	if _, err := eng.Get(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func createUser(eng Engine, user *User) error {
	if _, err := eng.Insert(user); err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User) error {
	sess := engine.NewSession()
	defer sess.Close()
	has, err := sess.Where("id = ?", user.Id).Exist(&User{})
	if err != nil {
		return err
	}
	if !has {
		return errors.New("user is empty")
	}

	if _, err := sess.Update(user, &User{Id: user.Id}); err != nil {
		return err
	}
	if err := sess.Commit(); err != nil {
		return err
	}
	logger.NewInfo("updated user_id: " + user.Id)
	return nil
}

func GetOrCreateUser(userId string) (*User, error) {
	sess := engine.NewSession()
	defer sess.Close()
	if userId == "" {
		return nil, errors.New("user id is empty")
	}

	user, err := getUser(sess, userId)
	if err != nil {
		return nil, err
	}

	if user.Id == "" {
		user.Id = userId
		if err := createUser(sess, user); err != nil {
			return nil, err
		}
	}

	if err := sess.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
