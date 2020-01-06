package controllers

import (
	"errors"
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/serializers"
	. "github.com/KouT127/attendance-management/backend/usecases"
	. "github.com/KouT127/attendance-management/backend/validators"
	. "github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewAttendanceController(usecase AttendanceInteractor) *attendanceController {
	return &attendanceController{
		usecase: usecase,
	}
}

type AttendanceController interface {
	AttendanceListController(c *Context)
	AttendanceCreateController(c *Context)
}

type attendanceController struct {
	usecase AttendanceInteractor
}

func (ac attendanceController) AttendanceListController(c *Context) {
	p := NewPagination(0, 5)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		err := errors.New("invalid user id")
		c.JSON(http.StatusNotFound, NewError("user_id", err))
		return
	}

	userId := value.(string)
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
	p := NewPagination(0, 31)
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
		// TODO: Validation Error Method
		c.JSON(http.StatusBadRequest, err)
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		err := errors.New("invalid user id")
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}
	t := new(AttendanceTime)
	t.Remark = input.Remark
	t.PushedAt = time.Now()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	userId := value.(string)
	query := new(Attendance)
	query.UserId = userId

	res, err := ac.usecase.CreateAttendance(query, t)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("attendance", err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetIdByKey(ctx *Context, key string) (string, error) {
	value, exists := ctx.Get(key)
	if !exists {
		return "", errors.New("user not found")
	}
	id := value.(string)
	return id, nil
}
