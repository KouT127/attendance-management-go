package models

import (
	"github.com/Songmu/flextime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttendance(t *testing.T) {
	t.Run("Testing Attendance data access", func(t *testing.T) {
		assert.Nil(t, SetTestDatabase())

		t.Run("Should not create attendance when userId is empty", func(t *testing.T) {
			attendanceTime := &AttendanceTime{
				Remark:     "test",
				IsModified: false,
				PushedAt:   flextime.Now(),
				CreatedAt:  flextime.Now(),
				UpdatedAt:  flextime.Now(),
			}
			attendance, err := CreateOrUpdateAttendance(attendanceTime, "")
			assert.NotNil(t, err)
			assert.Nil(t, attendance)
		})

		t.Run("Should not create attendance when time is empty", func(t *testing.T) {
			attendance, err := CreateOrUpdateAttendance(nil, "asdfjeijrlkjadf23laidf")
			assert.NotNil(t, err)
			assert.Nil(t, attendance)
		})

		t.Run("Should create clockIn", func(t *testing.T) {
			userId := "asdiekawei42lasedi356ladfkjfity3"
			user, err := createTestUser(userId)
			assert.Nil(t, err)

			attendanceTime := &AttendanceTime{
				Remark:     "test",
				IsModified: false,
				PushedAt:   flextime.Now(),
				CreatedAt:  flextime.Now(),
				UpdatedAt:  flextime.Now(),
			}
			attendance, err := CreateOrUpdateAttendance(attendanceTime, user.Id)
			assert.Nil(t, err)
			assert.NotNil(t, attendance.Id)
			assert.Equal(t, attendance.UserId, user.Id)
			assert.NotNil(t, attendance.ClockedIn)
		})

		t.Run("Should create clockOut", func(t *testing.T) {
			userId := "asdiekawei42lasedi356ladfkjfity2"
			user, err := createTestUser(userId)
			assert.Nil(t, err)
			attendanceTime := &AttendanceTime{
				Remark:     "test",
				IsModified: false,
				PushedAt:   flextime.Now(),
				CreatedAt:  flextime.Now(),
				UpdatedAt:  flextime.Now(),
			}
			attendance, err := CreateOrUpdateAttendance(attendanceTime, user.Id)
			assert.Nil(t, err)
			assert.NotNil(t, attendance.Id)
			assert.Equal(t, attendance.UserId, userId)
			assert.NotNil(t, attendance.ClockedIn)
			now := flextime.Now()
			attendanceTime2 := &AttendanceTime{
				Remark:     "test2",
				IsModified: false,
				PushedAt:   now,
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			attendance2, err := CreateOrUpdateAttendance(attendanceTime2, user.Id)
			assert.Nil(t, err)
			assert.NotNil(t, attendance2.Id)
			assert.Equal(t, attendance2.UserId, userId)
			assert.NotNil(t, attendance2.ClockedOut)
			assert.Equal(t, attendance2.ClockedOut.Remark, attendanceTime2.Remark)
		})

		t.Run("Should create clockOut", func(t *testing.T) {
			userId := "asdiekawei42lasedi356ladfkjfity4"
			user, err := createTestUser(userId)
			assert.Nil(t, err)
			attendanceTime := &AttendanceTime{
				Remark:     "test",
				IsModified: false,
				PushedAt:   flextime.Now(),
				CreatedAt:  flextime.Now(),
				UpdatedAt:  flextime.Now(),
			}
			attendance, err := CreateOrUpdateAttendance(attendanceTime, user.Id)
			assert.Nil(t, err)
			assert.NotNil(t, attendance.Id)
			assert.Equal(t, attendance.UserId, userId)
			assert.NotNil(t, attendance.ClockedIn)
			now := flextime.Now()
			attendanceTime2 := &AttendanceTime{
				Remark:     "test2",
				IsModified: false,
				PushedAt:   now,
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			attendance2, err := CreateOrUpdateAttendance(attendanceTime2, user.Id)
			assert.Nil(t, err)
			assert.NotNil(t, attendance2.Id)
			assert.Equal(t, attendance2.UserId, userId)
			assert.NotNil(t, attendance2.ClockedOut)
			assert.Equal(t, attendance2.ClockedOut.Remark, attendanceTime2.Remark)

			attendanceTime3 := &AttendanceTime{
				Remark:     "test3",
				IsModified: false,
				PushedAt:   now,
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			attendance3, err := CreateOrUpdateAttendance(attendanceTime3, user.Id)
			assert.Nil(t, err)
			assert.NotNil(t, attendance3.Id)
			assert.Equal(t, attendance3.UserId, userId)
			assert.NotNil(t, attendance3.ClockedOut)
			assert.Equal(t, attendance3.ClockedOut.Remark, attendanceTime3.Remark)
		})
	})
}
