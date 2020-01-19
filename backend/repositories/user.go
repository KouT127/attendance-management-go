package repositories

import (
	"github.com/KouT127/attendance-management/models"
	. "github.com/go-xorm/xorm"
	"time"
)

const (
	UserTable = "users"
)

type User struct {
	Id        string
	Name      string
	Email     string
	ImageUrl  string    `xorm:"image_url"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

func NewUser(u *models.User) *User {
	user := new(User)
	user.Id = u.Id
	user.Name = u.Name
	user.ImageUrl = u.ImageUrl
	user.UpdatedAt = time.Now()
	return user
}

func (u *User) build(user *models.User) {
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
	FetchUser(userId string, user *models.User) (bool, error)
	CreateUser(user *models.User) (int64, error)
	UpdateUser(user *models.User, id string) (int64, error)
}

type userRepository struct {
	engine Engine
}

func (r userRepository) FetchUsers(u *models.User) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := r.engine.
		Table(UserTable).
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*models.User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func (r *userRepository) FetchUser(userId string, user *models.User) (bool, error) {
	u := new(User)
	has, err := r.engine.
		Table(UserTable).
		Where("id = ?", userId).
		Get(u)
	u.build(user)
	return has, err
}

func (r *userRepository) CreateUser(user *models.User) (int64, error) {
	u := NewUser(user)
	cnt, err := r.engine.
		Table(UserTable).
		Insert(u)
	u.build(user)
	return cnt, err
}

func (r *userRepository) UpdateUser(user *models.User, id string) (int64, error) {
	u := NewUser(user)
	cnt, err := r.engine.
		Table(UserTable).
		Where("id = ?", user.Id).
		Update(u)
	u.build(user)
	return cnt, err
}
