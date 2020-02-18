package attendance

import (
	"errors"
	. "github.com/KouT127/attendance-management/handlers"
	"github.com/KouT127/attendance-management/middlewares"
	"github.com/KouT127/attendance-management/models"
	. "github.com/KouT127/attendance-management/responses"
	attendanceService "github.com/KouT127/attendance-management/services/attendance"
	"github.com/KouT127/attendance-management/utils/logger"
	. "github.com/KouT127/attendance-management/validators"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func V1LatestHandler(c *Context) {
	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}
	a := &models.Attendance{
		UserId: userId,
	}

	res, err := attendanceService.ViewLatestAttendance(a)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("attendances", err))
	}
	c.JSON(http.StatusOK, res)
	return
}

func V1ListHandler(c *Context) {
	p := NewPaginatorInput(0, 5)

	if err := c.Bind(p); err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(InvalidValueError)
		c.JSON(http.StatusBadRequest, NewError("seach", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		logger.NewFatal(c, err.Error())
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	a := &models.Attendance{UserId: userId}
	maxCnt, err := models.FetchAttendancesCount(a)
	if err != nil {
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("attendances", err))
	}
	attendances, err := attendanceService.ViewAttendances(p, a)
	if err != nil {
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("attendances", err))
	}

	responses := make([]*AttendanceResp, 0)
	for _, attendance := range attendances {
		resp := NewAttendanceResp(attendance)
		responses = append(responses, &resp)
	}

	res := new(AttendancesResult)
	res.HasNext = p.HasNext(maxCnt)
	res.IsSuccessful = true
	res.Attendances = responses
	c.JSON(http.StatusOK, res)
}

func V1MonthlyHandler(c *Context) {
	p := NewPaginatorInput(0, 31)
	s := NewSearchParams()

	if err := c.Bind(s); err != nil {
		err := errors.New(InvalidValueError)
		c.JSON(http.StatusBadRequest, NewError("search", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	q := &models.Attendance{
		UserId: userId,
	}

	res, err := attendanceService.ViewAttendancesMonthly(p, q)
	if err != nil {
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("attendances", err))
	}

	c.JSON(http.StatusOK, res)
}

func V1CreateHandler(c *Context) {
	var (
		input AttendanceInput
	)

	if err := c.Bind(&input); err != nil {
		err := errors.New(InvalidValueError)
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	userId, err := GetIdByKey(c, middlewares.AuthorizedUserIdKey)
	if err != nil {
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	query := new(models.Attendance)
	query.UserId = userId

	res, err := attendanceService.CreateOrUpdateAttendance(&input, query)
	if err != nil {
		err := errors.New(BadAccessError)
		c.JSON(http.StatusBadRequest, NewError("attendance", err))
		return
	}

	c.JSON(http.StatusOK, res)
}
