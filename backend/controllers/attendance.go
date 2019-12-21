package controllers

import (
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/repositories"
	. "github.com/KouT127/attendance-management/backend/responses"
	. "github.com/KouT127/attendance-management/backend/validators"
	. "github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewAttendanceController(repo AttendanceRepository) *attendanceController {
	return &attendanceController{repository: repo}
}

type AttendanceController interface {
	AttendanceListController(c *Context)
	AttendanceCreateController(c *Context)
}

type attendanceController struct {
	repository AttendanceRepository
}

func (ac attendanceController) AttendanceListController(c *Context) {
	var (
		responses []*AttendanceResponse
	)

	p := NewPagination(0, 5)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusNotFound, H{
			"message": "user not found",
		})
		return
	}

	userId := value.(string)

	a := &Attendance{
		UserId: userId,
	}

	maxCnt, err := ac.repository.FetchAttendancesCount(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}

	attendances := make([]*Attendance, 0)
	attendances, err = ac.repository.FetchAttendances(a, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}

	for _, attendance := range attendances {
		res := &AttendanceResponse{}
		res.SetAttendance(attendance)
		responses = append(responses, res)
	}

	c.JSON(http.StatusOK, H{
		"hasNext":     p.HasNext(maxCnt),
		"attendances": responses,
	})
}

func (ac attendanceController) AttendanceCreateController(c *Context) {
	var (
		input AttendanceInput
	)
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}

	userId := value.(string)

	a := Attendance{
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	t := AttendanceTime{
		Remark:    input.Remark,
		PushedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := ac.repository.CreateAttendanceTime(&t); err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}
	if input.Kind == 10 {
		a.ClockedInId = t
	} else {
		a.ClockedOutId = t
	}

	if _, err := ac.repository.CreateAttendance(&a); err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	res := AttendanceResponse{}
	res.SetAttendance(&a)

	c.JSON(http.StatusOK, H{
		"attendance": res,
	})
}

type DailyAttendance struct {
	Day          string
	ClockedInAt  time.Time
	ClockedOutAt time.Time
}

func (ac attendanceController) AttendanceMonthlyController(c *Context) {
	//p := NewPagination(0, 100)

	//if err := c.Bind(p); err != nil {
	//	c.JSON(http.StatusBadRequest, H{
	//		"message": err,
	//	})
	//	return
	//}
	//
	//value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	//if !exists {
	//	c.JSON(http.StatusNotFound, H{
	//		"message": "user not found",
	//	})
	//	return
	//}
	//
	//userId := value.(string)
	//
	//q := &Attendance{
	//	UserId: userId,
	//}
	//
	//attendances := make([]*Attendance, 0)
	//attendances, err := ac.repository.FetchAttendances(q, p)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, H{
	//		"message": err,
	//	})
	//	return
	//}

}
