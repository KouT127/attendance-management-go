package repositories

import (
	. "github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/domains"
	. "github.com/go-xorm/xorm"
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

func NewUser(u *domains.User) *User {
	user := new(User)
	user.Id = u.Id
	user.Name = u.Name
	user.ImageUrl = u.ImageUrl
	user.UpdatedAt = time.Now()
	return user
}

func (u *User) build(user *domains.User) {
	user.Id = u.Id
	user.Name = u.Name
	user.ImageUrl = u.ImageUrl
}

func NewUserRepository(e Engine) *userRepository {
	return &userRepository{
		engine: e,
	}
}

type UserRepository interface {
	FetchUser(userId string, user *domains.User) (bool, error)
	CreateUser(user *domains.User) (int64, error)
	UpdateUser(user *domains.User, id string) (int64, error)
}

type userRepository struct {
	engine Engine
}

func (r userRepository) FetchUsers(u *domains.User) ([]*domains.User, error) {
	users := make([]*domains.User, 0)
	err := r.engine.
		Table(UserTable).
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*domains.User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func (r *userRepository) FetchUser(userId string, user *domains.User) (bool, error) {
	u := new(User)
	has, err := r.engine.
		Table(UserTable).
		Where("id = ?", userId).
		Get(u)
	u.build(user)
	return has, err
}

func (r *userRepository) CreateUser(user *domains.User) (int64, error) {
	u := NewUser(user)
	cnt, err := r.engine.
		Table(UserTable).
		Insert(u)
	u.build(user)
	return cnt, err
}

func (r *userRepository) UpdateUser(user *domains.User, id string) (int64, error) {
	u := NewUser(user)
	cnt, err := r.engine.
		Table(UserTable).
		Where("id = ?", user.Id).
		Update(u)
	u.build(user)
	return cnt, err
}
