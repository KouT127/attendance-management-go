package repositories

import (
	. "github.com/KouT127/Attendance-management/backend/models"
	. "github.com/go-xorm/xorm"
)

const (
	UserTable = "users"
)

func NewUserRepository(e Engine) *userRepository {
	return &userRepository{
		engine: e,
	}
}

type UserRepository interface {
	FetchUser(userId string, u *User) (bool, error)
	FetchUsers(u *User) ([]*User, error)
	CreateUser(u *User) (int64, error)
	UpdateUser(u *User, q *User) (int64, error)
}

type userRepository struct {
	engine Engine
}

func (ur userRepository) FetchUsers(u *User) ([]*User, error) {
	users := make([]*User, 0)
	err := ur.engine.
		Table(UserTable).
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func (ur userRepository) FetchUser(userId string, u *User) (bool, error) {
	return ur.engine.
		Table(UserTable).
		Where("id = ?", userId).
		Get(u)
}

func (ur userRepository) CreateUser(u *User) (int64, error) {
	return ur.engine.
		Table(UserTable).
		Insert(u)
}

func (ur userRepository) UpdateUser(u *User, q *User) (int64, error) {
	return ur.engine.
		Table(UserTable).
		Update(u, q)
}
