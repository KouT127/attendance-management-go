package controllers

import (
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/serializers"
	. "github.com/KouT127/attendance-management/backend/usecases"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func NewAttendanceController(usecase AttendanceInteractor) *attendanceController {
	return &attendanceController{
		usecase: usecase,
	}
}

type AttendanceController interface {
	AttendanceListController(c *Context)
	AttendanceCreateController(c *Context)
	AttendanceMonthlyController(c *Context)
}

type attendanceController struct {
	usecase AttendanceInteractor
}

func (ac attendanceController) AttendanceListController(c *Context) {
	p := NewPaginatorInput(0, 5)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	a := &Attendance{
		UserId: userId,
	}

	res, err := ac.usecase.ViewAttendances(p, a)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func (ac attendanceController) AttendanceMonthlyController(c *Context) {
	p := NewPaginatorInput(0, 31)
	s := NewSearchParams()

	if err := c.Bind(s); err != nil {
		c.JSON(http.StatusBadRequest, NewError("search", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	q := &Attendance{
		UserId: userId,
	}

	res, err := ac.usecase.ViewAttendancesMonthly(p, q)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func (ac attendanceController) AttendanceCreateController(c *Context) {
	var (
		input AttendanceInput
	)

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	query := new(Attendance)
	query.UserId = userId

	res, err := ac.usecase.CreateAttendance(&input, query)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("attendance", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
