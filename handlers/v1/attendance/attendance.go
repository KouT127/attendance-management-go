package attendance

import (
	. "github.com/KouT127/attendance-management/handlers"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/auth"
	. "github.com/KouT127/attendance-management/modules/input"
	"github.com/KouT127/attendance-management/modules/logger"
	. "github.com/KouT127/attendance-management/modules/response"
	. "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ListHandler(c *Context) {
	p := NewPaginatorInput(0, 5)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	userId, err := GetIdByKey(c, auth.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	maxCnt, err := models.FetchAttendancesCount(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	opt := &models.AttendanceSearchOption{}
	opt.UserId = userId
	opt.Paginator = p.ToPaginator()
	attendances, err := models.FetchAttendances(opt)
	if err != nil {
		logger.NewWarn(logrus.Fields{}, err.Error())
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	hasNext := p.HasNext(maxCnt)
	res := ToAttendancesResult(hasNext, attendances)
	c.JSON(http.StatusOK, res)
}

func MonthlyHandler(c *Context) {
	p := NewPaginatorInput(0, 31)
	s := NewSearchParams()

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if err := c.Bind(s); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	userId, err := GetIdByKey(c, auth.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	maxCnt, err := models.FetchAttendancesCount(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	opt := &models.AttendanceSearchOption{}
	opt.UserId = userId
	opt.Paginator = p.ToPaginator()
	attendances, err := models.FetchAttendances(opt)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	hasNext := p.HasNext(maxCnt)
	res := ToAttendancesResult(hasNext, attendances)
	c.JSON(http.StatusOK, res)
}

func CreateHandler(c *Context) {
	input := new(AttendanceInput)
	if err := c.Bind(input); err != nil {
		c.JSON(http.StatusBadRequest, NewValidationError("user", err))
		return
	}

	userId, err := GetIdByKey(c, auth.AuthorizedUserIdKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	attendanceTime := input.ToAttendanceTime()

	attendance, err := models.CreateOrUpdateAttendance(attendanceTime, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	res := ToAttendanceResult(attendance)
	c.JSON(http.StatusOK, res)
}
