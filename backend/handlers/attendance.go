package handlers

import (
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/serializers"
	. "github.com/KouT127/attendance-management/backend/usecases"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func NewAttendanceHandler(usecase AttendanceInteractor) *attendanceHandler {
	return &attendanceHandler{
		usecase: usecase,
	}
}

type AttendanceHandler interface {
	AttendanceListHandler(c *Context)
	AttendanceCreateHandler(c *Context)
	AttendanceMonthlyHandler(c *Context)
}

type attendanceHandler struct {
	usecase AttendanceInteractor
}

func (ac attendanceHandler) AttendanceListHandler(c *Context) {
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

func (ac attendanceHandler) AttendanceMonthlyHandler(c *Context) {
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

func (ac attendanceHandler) AttendanceCreateHandler(c *Context) {
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
