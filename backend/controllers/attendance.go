package controllers

import (
	"github.com/KouT127/Attendance-management/backend/database"
	. "github.com/KouT127/Attendance-management/backend/domains"
	"github.com/KouT127/Attendance-management/backend/middlewares"
	. "github.com/KouT127/Attendance-management/backend/validators"
	. "github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type AttendanceController struct{}

func (uc AttendanceController) AttendanceListController(c *Context) {
	var (
		attendances []*Attendance
		userId      string
	)

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

	engine := database.NewDB()
	err = engine.
		Table("attendances").
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

	c.JSON(http.StatusOK, H{
		"attendances": attendances,
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
		UserId: userId,
		Kind:   input.Kind,
		Remark: input.Remark,
	}

	if _, err := engine.Table("attendances").Insert(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"attendance": attendance,
	})
}
