package usecases

import (
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/repositories"
)

func NewUserInteractor(repo UserRepository) *userInteractor {
	return &userInteractor{
		repository: repo,
	}
}

type UserInteractor interface {
	ViewUser(userId string) (*User, error)
	UpdateUser(userId string, userName string) (*User, error)
}

type userInteractor struct {
	repository UserRepository
}

func (i *userInteractor) ViewUser(userId string) (*User, error) {
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

func (i *userInteractor) UpdateUser(userId string, userName string) (*User, error) {
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
