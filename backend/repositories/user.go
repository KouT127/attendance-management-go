package repositories

import (
	. "github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"

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

func NewUserRepository() *userRepository {
	return &userRepository{}
}

type UserRepository interface {
	FetchUser(eng *Engine, userId string, user *models.User) (bool, error)
	CreateUser(eng *Engine, user *models.User) (int64, error)
	UpdateUser(eng *Engine, user *models.User, id string) (int64, error)
}

type userRepository struct{}

func (r *userRepository) FetchUsers(eng *Engine, u *models.User) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := eng.
		Table(UserTable).
		Iterate(u, func(idx int, bean interface{}) error {
			u := bean.(*models.User)
			users = append(users, u)
			return nil
		})
	return users, err
}

func (r *userRepository) FetchUser(eng *Engine, userId string, user *models.User) (bool, error) {
	u := new(User)
	has, err := eng.
		Table(UserTable).
		Where("id = ?", userId).
		Get(u)
	u.build(user)
	return has, err
}

func (r *userRepository) CreateUser(eng *Engine, user *models.User) (int64, error) {
	u := NewUser(user)
	cnt, err := eng.
		Table(UserTable).
		Insert(u)
	u.build(user)
	return cnt, err
}

func (r *userRepository) UpdateUser(eng *Engine, user *models.User, id string) (int64, error) {
	u := NewUser(user)
	cnt, err := eng.
		Table(UserTable).
		Where("id = ?", user.Id).
		Update(u)
	u.build(user)
	return cnt, err
}
