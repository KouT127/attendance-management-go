package usecases

import (
	. "github.com/KouT127/attendance-management/models"
	. "github.com/KouT127/attendance-management/repositories"
)

func NewUserUsecase(repo UserRepository) *userUsecase {
	return &userUsecase{
		repository: repo,
	}
}

type UserUsecase interface {
	ViewUser(userId string) (*User, error)
	UpdateUser(userId string, userName string) (*User, error)
}

type userUsecase struct {
	repository UserRepository
}

func (i *userUsecase) ViewUser(userId string) (*User, error) {
	u := &User{}
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

func (i *userUsecase) UpdateUser(userId string, userName string) (*User, error) {
	u := &User{}
	has, err := i.repository.FetchUser(userId, u)
	if err != nil || !has {
		return nil, err
	}

	u.Name = userName
	_, err = i.repository.UpdateUser(u, &User{Id: u.Id})
	if err != nil {
		return nil, err
	}
	return u, nil
}
