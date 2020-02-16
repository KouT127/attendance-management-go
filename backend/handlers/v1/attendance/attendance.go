package handlers

import (
	"errors"
	. "github.com/KouT127/attendance-management/handlers"
	"github.com/KouT127/attendance-management/middlewares"
	. "github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/responses"
	. "github.com/KouT127/attendance-management/service/attendance"
	. "github.com/KouT127/attendance-management/usecases"
	"github.com/KouT127/attendance-management/utils/logger"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func AttendanceLatestHandler(c *Context) {
	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}
	a := &Attendance{
		UserId: userId,
	}

	res, err := ViewLatestAttendance(a)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendances", err))
	}
	c.JSON(http.StatusOK, res)
	return
}

func AttendanceListHandler(c *Context) {
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

	res, err := ViewAttendances(p, a)
	if err != nil {
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func AttendanceMonthlyHandler(c *Context) {
	p := NewPaginatorInput(0, 31)
	s := NewSearchParams()

	if err := c.Bind(s); err != nil {
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

	res, err := ViewAttendancesMonthly(p, q)
	if err != nil {
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func AttendanceCreateHandler(c *Context) {
	var (
		input AttendanceInput
	)

	if err := c.Bind(&input); err != nil {
		err := errors.New(responses.InvalidValueError)
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	query := new(Attendance)
	query.UserId = userId

	res, err := CreateOrUpdateAttendance(&input, query)
	if err != nil {
		err := errors.New(responses.BadAccessError)
		c.JSON(http.StatusBadRequest, responses.NewError("attendance", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
