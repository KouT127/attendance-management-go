package usecases

import (
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/repositories"
	"github.com/KouT127/attendance-management/utils/logger"
	"github.com/go-xorm/xorm"
)

func NewUserUsecase(userRepo repositories.UserRepository, attendanceRepo repositories.AttendanceRepository) *userUsecase {
	return &userUsecase{
		userRepository:       userRepo,
		attendanceRepository: attendanceRepo,
	}
}

type UserUsecase interface {
	ViewUser(eng *xorm.Engine, userId string) (*models.User, *models.Attendance, error)
	UpdateUser(eng *xorm.Engine, userId string, userName string) (*models.User, error)
}

type userUsecase struct {
	userRepository       repositories.UserRepository
	attendanceRepository repositories.AttendanceRepository
}

func (i *userUsecase) ViewUser(eng *xorm.Engine, userId string) (*models.User, *models.Attendance, error) {
	user := new(models.User)
	attendance := new(models.Attendance)
	has, err := i.userRepository.FetchUser(eng, userId, user)
	if err != nil {
		return nil, nil, err
	}

	if !has {
		user.Id = userId
		_, err := i.userRepository.CreateUser(eng, user)
		if err != nil {
			return nil, nil, err
		}
	}
	attendance.UserId = user.Id
	attendance, err = i.attendanceRepository.FetchLatestAttendance(eng, attendance)
	if err != nil {
		return nil, nil, err
	}
	return user, attendance, nil
}

func (i *userUsecase) UpdateUser(eng *xorm.Engine, userId string, userName string) (*models.User, error) {
	u := new(models.User)
	has, err := i.userRepository.FetchUser(eng, userId, u)
	if err != nil || !has {
		return nil, err
	}
	u.Name = userName
	_, err = i.userRepository.UpdateUser(eng, u, u.Id)
	if err != nil {
		return nil, err
	}
	logger.NewInfo("updated user-" + u.Id)

	return u, nil
}
