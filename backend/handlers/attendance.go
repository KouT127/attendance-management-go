package handlers

import (
	"errors"
	"github.com/KouT127/attendance-management/middlewares"
	. "github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/responses"
	. "github.com/KouT127/attendance-management/usecases"
	"github.com/KouT127/attendance-management/utils/logger"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func NewAttendanceHandler(usecase AttendanceUsecase) *attendanceHandler {
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
	usecase AttendanceUsecase
}

func (ac *attendanceHandler) AttendanceLatestHandler(c *Context) {
	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}
	a := &Attendance{
		UserId: userId,
	}

	res, err := ac.usecase.ViewLatestAttendance(a)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendances", err))
	}
	c.JSON(http.StatusOK, res)
	return
}

func (ac *attendanceHandler) AttendanceListHandler(c *Context) {
	p := NewPaginatorInput(0, 5)

	if err := c.Bind(p); err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.InvalidValueError)
		c.JSON(http.StatusBadRequest, responses.NewError("seach", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	a := &Attendance{
		UserId: userId,
	}

	res, err := ac.usecase.ViewAttendances(p, a)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func (ac *attendanceHandler) AttendanceMonthlyHandler(c *Context) {
	p := NewPaginatorInput(0, 31)
	s := NewSearchParams()

	if err := c.Bind(s); err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.InvalidValueError)
		c.JSON(http.StatusBadRequest, responses.NewError("search", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	q := &Attendance{
		UserId: userId,
	}

	res, err := ac.usecase.ViewAttendancesMonthly(p, q)
	if err != nil {
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func (ac *attendanceHandler) AttendanceCreateHandler(c *Context) {
	var (
		input AttendanceInput
	)

	if err := c.Bind(&input); err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.InvalidValueError)
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	query := new(Attendance)
	query.UserId = userId

	res, err := ac.usecase.CreateAttendance(&input, query)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendance", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
