package repositories

import (
	"github.com/KouT127/attendance-management/database"
	"github.com/KouT127/attendance-management/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestAttendanceRepository_CreateAttendance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		database.ConnectDatabase()
		tearDown := database.PrepareTestDatabase()
		defer tearDown()
		eng := database.NewDB()
		mockUser := models.User{
			Id:   "1",
			Name: "test",
		}
		_, err := eng.Insert(&mockUser)
		if err != nil {
			log.Fatal(err)
		}
		mockTime := AttendanceTime{
			Id:        1,
			Remark:    "test",
			PushedAt:  time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, _ = eng.Insert(&mockTime)

		repo := NewAttendanceRepository()
		sess := repo.NewSession(eng)

		mockAttendanceTime := models.AttendanceTime{
			Id:        1,
			Remark:    "",
			PushedAt:  time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockAttendance := models.Attendance{
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
	})

	t.Run("failure", func(t *testing.T) {
		assert.Equal(t, "", "")
	})
}
