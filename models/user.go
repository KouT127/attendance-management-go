package models

import (
	"errors"
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/modules/logger"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return database.UserTable
}

func getUser(eng Engine, userID string) (*User, error) {
	u := User{ID: userID}
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
	has, err := sess.Where("id = ?", user.ID).Exist(&User{})
	if err != nil {
		return err
	}
	if !has {
		return errors.New("user is empty")
	}

	if _, err := sess.Update(user, &User{ID: user.ID}); err != nil {
		return err
	}
	if err := sess.Commit(); err != nil {
		return err
	}
	logger.NewInfo("updated user_id: " + user.ID)
	return nil
}

func GetOrCreateUser(userID string) (*User, error) {
	sess := engine.NewSession()
	defer sess.Close()
	if userID == "" {
		return nil, errors.New("user id is empty")
	}

	user, err := getUser(sess, userID)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		user.ID = userID
		if err := createUser(sess, user); err != nil {
			return nil, err
		}
	}

	if err := sess.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
