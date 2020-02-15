package repositories

import (
	"context"
	"database/sql"
	database "github.com/KouT127/attendance-management/database/gen"
	"github.com/KouT127/attendance-management/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

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

func NewUser(u *models.User) *database.User {
	user := new(database.User)
	user.ID = u.Id
	user.Name = null.StringFrom(u.Name)
	user.ImageURL = null.StringFrom(u.ImageUrl)
	user.CreatedAt = null.TimeFrom(time.Now())
	user.UpdatedAt = null.TimeFrom(time.Now())
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
	FetchUser(ctx context.Context, db *sql.DB, userId string, user *models.User) (bool, error)
	CreateUser(ctx context.Context, db *sql.DB, user *models.User) error
	UpdateUser(ctx context.Context, db *sql.DB, user *models.User, id string) error
}

type userRepository struct{}

func (r *userRepository) FetchUser(ctx context.Context, db *sql.DB, userId string, user *models.User) (bool, error) {
	u, err := database.FindUser(ctx, db, userId)
	if err != nil {
		return false, err
	}
	has := u != nil
	user.Id = u.ID
	user.Name = u.Name.String
	user.ImageUrl = u.ImageURL.String
	user.Email = u.Email
	return has, err
}

func (r *userRepository) CreateUser(ctx context.Context, db *sql.DB, user *models.User) error {
	u := NewUser(user)
	if err := u.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}
	user.Id = u.ID
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, db *sql.DB, user *models.User, id string) error {
	u := NewUser(user)
	if _, err := u.Update(ctx, db, boil.Infer()); err != nil {
		return err
	}
	return nil
}
