package models

import (
	"github.com/KouT127/attendance-management/database"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	ImageUrl  string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

func NewUser(u *User) *User {
	user := new(User)
	user.Id = u.Id
	user.Name = u.Name
	user.ImageUrl = u.ImageUrl
	user.UpdatedAt = time.Now()
	return user
}

func (u *User) build(user *User) {
	user.Id = u.Id
	user.Name = u.Name
	user.ImageUrl = u.ImageUrl
}

func FetchUsers(u *User) ([]*User, error) {
	users := make([]*User, 0)
	err := engine.
		Table(database.UserTable).
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func FetchUser(userId string) (*User, error) {
	u := new(User)
	_, err := engine.
		Table(database.UserTable).
		Where("id = ?", userId).
		Get(u)
	return u, err
}

func CreateUser(user *User) (int64, error) {
	u := NewUser(user)
	cnt, err := engine.
		Table(database.UserTable).
		Insert(u)
	u.build(user)
	return cnt, err
}

func UpdateUser(user *User) (int64, error) {
	cnt, err := engine.
		Table(database.UserTable).
		Where("id = ?", user.Id).
		Update(user)
	return cnt, err
}
