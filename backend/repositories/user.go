package repositories

import (
	"github.com/KouT127/attendance-management/backend/models"
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
	FetchUser(userId string, u *models.User) (bool, error)
	FetchUsers(u *models.User) ([]*models.User, error)
	CreateUser(u *models.User) (int64, error)
	UpdateUser(u *models.User, q *models.User) (int64, error)
}

type userRepository struct {
	engine Engine
}

func (ur userRepository) FetchUsers(u *models.User) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := ur.engine.
		Table(UserTable).
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*models.User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func (ur userRepository) FetchUser(userId string, u *models.User) (bool, error) {
	return ur.engine.
		Table(UserTable).
		Where("id = ?", userId).
		Get(u)
}

func (ur userRepository) CreateUser(u *models.User) (int64, error) {
	return ur.engine.
		Table(UserTable).
		Insert(u)
}

func (ur userRepository) UpdateUser(u *models.User, q *models.User) (int64, error) {
	return ur.engine.
		Table(UserTable).
		Update(u, q)
}
