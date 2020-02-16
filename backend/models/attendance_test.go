package models

import (
	"reflect"
	"testing"
	"time"
)

func createAttendance() *Attendance {
	return &Attendance{
		Id:        1,
		UserId:    "abcd1234",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func createAttendanceTime(id int64) *AttendanceTime {
	return &AttendanceTime{
		Id:        id,
		Remark:    "remark",
		PushedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestAttendance_ClockIn(t *testing.T) {
	t.Run("insert_attendance_time", func(t *testing.T) {
		a := createAttendance()
		time := createAttendanceTime(1)
		a.ClockIn(time)

		if !reflect.DeepEqual(a.ClockedIn, time) {
			t.Fatal()
		}
	})

	t.Run("insert_nil", func(t *testing.T) {
		a := createAttendance()
		var time *AttendanceTime
		a.ClockIn(time)

		if !reflect.DeepEqual(a.ClockedIn, time) {
			t.Fatal()
		}
	})
}

func TestAttendance_ClockOut(t *testing.T) {
	t.Run("insert attendance time", func(t *testing.T) {
		a := createAttendance()
		time1 := createAttendanceTime(1)
		time2 := createAttendanceTime(2)
		a.ClockIn(time1)
		a.ClockOut(time2)

		if a.ClockedIn == nil || !reflect.DeepEqual(a.ClockedIn, time1) {
			t.Fatal("missing clocked in time")
		}
		if !reflect.DeepEqual(a.ClockedOut, time2) {
			t.Fatal()
		}
	})

	t.Run("insert nil", func(t *testing.T) {
		a := createAttendance()
		var time *AttendanceTime
		a.ClockOut(time)

		if !reflect.DeepEqual(a.ClockedOut, time) {
			t.Fatal()
		}
	})
}

func TestAttendance_IsClockedOut(t *testing.T) {
	t.Run("time is clocked out", func(t *testing.T) {
		a := createAttendance()
		time := createAttendanceTime(1)
		a.ClockOut(time)

		if !a.IsClockedOut() {
			t.Fatal("isn't clocked out")
		}
	})

	t.Run("time isn't clocked out", func(t *testing.T) {
		a := createAttendance()

		if a.IsClockedOut() {
			t.Fatal("clocked out")
		}
	})
}

//func insertUser(eng *xorm.Engine) {
//	mockUser := models.User{
//		Id:   "1",
//		Name: "test",
//	}
//	_, err := eng.Table(database.UserTable).Insert(&mockUser)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func insertTime(eng *xorm.Engine) {
//	mockTime := AttendanceTime{
//		Id:        1,
//		Remark:    "test",
//		PushedAt:  time.Now(),
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//	_, _ = eng.Table(database.AttendanceTimeTable).Insert(&mockTime)
//}
//
//func TestAttendanceRepository_CreateAttendance(t *testing.T) {
//	t.Run("success", func(t *testing.T) {
//		database.ConnectDatabase()
//		tearDown := database.PrepareTestDatabase()
//		defer tearDown()
//		eng := database.NewDB()
//
//		insertUser(eng)
//		insertTime(eng)
//
//		repo := NewAttendanceRepository()
//		sess := repo.NewSession(eng)
//
//		mockAttendanceTime := models.AttendanceTime{
//			Id:        1,
//			Remark:    "test",
//			PushedAt:  time.Now(),
//			CreatedAt: time.Now(),
//			UpdatedAt: time.Now(),
//		}
//
//		mockAttendance := models.Attendance{
//			UserId:    "1",
//			ClockedIn: &mockAttendanceTime,
//			CreatedAt: time.Time{},
//			UpdatedAt: time.Time{},
//		}
//
//		cnt, err := repo.CreateAttendance(sess, &mockAttendance)
//		if err != nil {
//			log.Fatal(err)
//		}
//		err = repo.Commit(sess)
//		if err != nil {
//			log.Fatal(err)
//		}
//		assert.Equal(t, int64(1), cnt)
//		assert.NotNil(t, &mockAttendance.Id)
//		assert.NotNil(t, mockAttendance.ClockedIn)
//		assert.Nil(t, mockAttendance.ClockedOut)
//	})
//
//	t.Run("failure", func(t *testing.T) {
//		database.ConnectDatabase()
//		tearDown := database.PrepareTestDatabase()
//		defer tearDown()
//		eng := database.NewDB()
//
//		insertTime(eng)
//
//		repo := NewAttendanceRepository()
//		sess := repo.NewSession(eng)
//
//		mockAttendanceTime := models.AttendanceTime{
//			Id:        1,
//			Remark:    "test",
//			PushedAt:  time.Now(),
//			CreatedAt: time.Now(),
//			UpdatedAt: time.Now(),
//		}
//
//		mockAttendance := models.Attendance{
//			UserId:    "1",
//			ClockedIn: &mockAttendanceTime,
//			CreatedAt: time.Time{},
//			UpdatedAt: time.Time{},
//		}
//
//		_, err := repo.CreateAttendance(sess, &mockAttendance)
//		assert.NotNil(t, err)
//	})
//}
