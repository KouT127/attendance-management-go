package models

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/utils/logger"
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

func FetchUsers(u *User) ([]*User, error) {
	users := make([]*User, 0)
	err := engine.
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func GetUser(userId string) (*User, error) {
	return getUser(engine, userId)
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
	if has, err := sess.Exist(user); err != nil || !has {
		return err
	}

	if _, err := sess.Update(user, &User{Id: user.Id}); err != nil {
		return err
	}
	logger.NewInfo("updated user-" + user.Id)
	return nil
}

func GetOrCreateUser(userId string) (*User, error) {
	sess := engine.NewSession()
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

	return user, nil
}
