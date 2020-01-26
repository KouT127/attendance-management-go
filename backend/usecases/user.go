package usecases

import (
	"github.com/KouT127/attendance-management/domains"
	"github.com/KouT127/attendance-management/repositories"
	"github.com/KouT127/attendance-management/utils/logger"
)

func NewUserUsecase(repo repositories.UserRepository) *userUsecase {
	return &userUsecase{
		repository: repo,
	}
}

type UserUsecase interface {
	ViewUser(userId string) (*domains.User, error)
	UpdateUser(userId string, userName string) (*domains.User, error)
}

type userUsecase struct {
	repository repositories.UserRepository
}

func (i *userUsecase) ViewUser(userId string) (*domains.User, error) {
	u := new(domains.User)
	has, err := i.repository.FetchUser(userId, u)
	if err != nil {
		return nil, err
	}

	if !has {
		u.Id = userId
		_, err := i.repository.CreateUser(u)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}

func (i *userUsecase) UpdateUser(userId string, userName string) (*domains.User, error) {
	u := new(domains.User)
	has, err := i.repository.FetchUser(userId, u)
	if err != nil || !has {
		return nil, err
	}
	u.Name = userName
	_, err = i.repository.UpdateUser(u, u.Id)
	if err != nil {
		return nil, err
	}
	logger.NewInfo("updated user-" + u.Id)

	return u, nil
}
