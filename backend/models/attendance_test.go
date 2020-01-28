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

func createAttendanceTime() *AttendanceTime {
	return &AttendanceTime{
		Id:        1,
		Remark:    "remark",
		PushedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestAttendance_ClockIn(t *testing.T) {
	t.Run("insert_attendance_time", func(t *testing.T) {
		a := createAttendance()
		time := createAttendanceTime()
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
