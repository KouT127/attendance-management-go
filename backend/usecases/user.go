package usecases

import (
	"context"
	. "github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/repositories"
	"github.com/KouT127/attendance-management/utils/logger"
)

func NewUserUsecase(userRepo repositories.UserRepository, attendanceRepo repositories.AttendanceRepository) *userUsecase {
	return &userUsecase{
		userRepository:       userRepo,
		attendanceRepository: attendanceRepo,
	}
}

type UserUsecase interface {
	ViewUser(userId string) (*models.User, *models.Attendance, error)
	UpdateUser(userId string, userName string) (*models.User, error)
}

type userUsecase struct {
	userRepository       repositories.UserRepository
	attendanceRepository repositories.AttendanceRepository
}

func (i *userUsecase) ViewUser(userId string) (*models.User, *models.Attendance, error) {
	user := new(models.User)
	attendance := new(models.Attendance)
	ctx := context.Background()
	db := NewDB()
	has, err := i.userRepository.FetchUser(ctx, db, userId, user)
	if err != nil {
		return nil, nil, err
	}

	if !has {
		user.Id = userId
		err := i.userRepository.CreateUser(ctx, db, user)
		if err != nil {
			return nil, nil, err
		}
	}
	attendance.UserId = user.Id
	attendance, err = i.attendanceRepository.FetchLatestAttendance(ctx, db, attendance)
	if err != nil {
		return nil, nil, err
	}
	return user, attendance, nil
}

func (i *userUsecase) UpdateUser(userId string, userName string) (*models.User, error) {
	u := new(models.User)
	ctx := context.Background()
	db := NewDB()
	has, err := i.userRepository.FetchUser(ctx, db, userId, u)
	if err != nil || !has {
		return nil, err
	}
	u.Name = userName
	err = i.userRepository.UpdateUser(ctx, db, u, u.Id)
	if err != nil {
		return nil, err
	}
	logger.NewInfo("updated user-" + u.Id)
	return u, nil
}
