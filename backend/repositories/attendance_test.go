package repositories

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/domains"
	"github.com/go-xorm/xorm"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func insertUser(eng *xorm.Engine) {
	mockUser := domains.User{
		Id:   "1",
		Name: "test",
	}
	_, err := eng.Table(database.UserTable).Insert(&mockUser)
	if err != nil {
		log.Fatal(err)
	}
}

func insertTime(eng *xorm.Engine) {
	mockTime := AttendanceTime{
		Id:        1,
		Remark:    "test",
		PushedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, _ = eng.Table(database.AttendanceTimeTable).Insert(&mockTime)
}

func TestAttendanceRepository_CreateAttendance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		database.ConnectDatabase()
		tearDown := database.PrepareTestDatabase()
		defer tearDown()
		eng := database.NewDB()

		insertUser(eng)
		insertTime(eng)

		repo := NewAttendanceRepository()
		sess := repo.NewSession(eng)

		mockAttendanceTime := domains.AttendanceTime{
			Id:        1,
			Remark:    "test",
			PushedAt:  time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockAttendance := domains.Attendance{
			UserId:    "1",
			ClockedIn: &mockAttendanceTime,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		cnt, err := repo.CreateAttendance(sess, &mockAttendance)
		if err != nil {
			log.Fatal(err)
		}
		err = repo.Commit(sess)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, int64(1), cnt)
		assert.NotNil(t, &mockAttendance.Id)
		assert.NotNil(t, mockAttendance.ClockedIn)
		assert.Nil(t, mockAttendance.ClockedOut)
	})

	t.Run("failure", func(t *testing.T) {
		database.ConnectDatabase()
		tearDown := database.PrepareTestDatabase()
		defer tearDown()
		eng := database.NewDB()

		insertTime(eng)

		repo := NewAttendanceRepository()
		sess := repo.NewSession(eng)

		mockAttendanceTime := domains.AttendanceTime{
			Id:        1,
			Remark:    "test",
			PushedAt:  time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockAttendance := domains.Attendance{
			UserId:    "1",
			ClockedIn: &mockAttendanceTime,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		_, err := repo.CreateAttendance(sess, &mockAttendance)
		assert.NotNil(t, err)
	})
}
