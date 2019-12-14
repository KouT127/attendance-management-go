package controllers

import (
	"github.com/KouT127/Attendance-management/backend/database"
	. "github.com/KouT127/Attendance-management/backend/domains"
	"github.com/KouT127/Attendance-management/backend/middlewares"
	. "github.com/KouT127/Attendance-management/backend/repositories"
	. "github.com/KouT127/Attendance-management/backend/responses"
	. "github.com/KouT127/Attendance-management/backend/validators"
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
		attendance Attendance
		input      AttendanceInput
	)
	engine := database.NewDB()
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

	attendance = Attendance{
		UserId:    userId,
		Kind:      input.Kind,
		Remark:    input.Remark,
		PushedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := engine.Table("attendances").Insert(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	res := AttendanceResponse{}
	res.SetAttendance(&attendance)

	c.JSON(http.StatusOK, H{
		"attendance": res,
	})
}
