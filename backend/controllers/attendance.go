package controllers

import (
	"github.com/KouT127/Attendance-management/backend/database"
	. "github.com/KouT127/Attendance-management/backend/domains"
	"github.com/KouT127/Attendance-management/backend/middlewares"
	. "github.com/KouT127/Attendance-management/backend/responses"
	. "github.com/KouT127/Attendance-management/backend/validators"
	. "github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

type AttendanceController struct{}

func (uc AttendanceController) AttendanceListController(c *Context) {
	var (
		responses   []*AttendanceResponse
		attendances []*Attendance
		userId      string
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
	err := mapstructure.Decode(value, &userId)
	if err != nil {
		c.JSON(http.StatusNotFound, H{
			"message": "user not found",
		})
		return
	}
	page := p.CalculatePage()
	engine := database.NewDB()

	maxCnt, err := engine.
		Table("attendances").
		Count(&Attendance{})
	if err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}
	err = engine.
		Table("attendances").
		Limit(int(p.Limit), int(page)).
		OrderBy("-id").
		Iterate(&Attendance{UserId: userId}, func(idx int, bean interface{}) error {
			attendance := bean.(*Attendance)
			attendances = append(attendances, attendance)
			return nil
		})

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

func (uc AttendanceController) AttendanceCreateController(c *Context) {
	var (
		userId     string
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

	err := mapstructure.Decode(value, &userId)
	if err != nil {
		c.JSON(http.StatusNotFound, H{})
		return
	}

	attendance = Attendance{
		UserId:    userId,
		Kind:      input.Kind,
		Remark:    input.Remark,
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
