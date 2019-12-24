package controllers

import (
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/models"
	. "github.com/KouT127/attendance-management/backend/repositories"
	. "github.com/KouT127/attendance-management/backend/responses"
	. "github.com/KouT127/attendance-management/backend/validators"
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
	responses := make([]*AttendanceResponse, 0)

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

	res := new(response)
	res.HasNext = p.HasNext(maxCnt)
	res.IsSuccessful = true
	res.Attendances = responses

	c.JSON(http.StatusOK, res)
}

type response struct {
	IsSuccessful bool `json:"isSuccessful"`
	HasNext      bool `json:"hasNext"`
	Attendances  []*AttendanceResponse `json:"attendances"`
}

func (ac attendanceController) AttendanceCreateController(c *Context) {
	var (
		input AttendanceInput
	)
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}
	t := AttendanceTime{
		Remark:    input.Remark,
		PushedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userId := value.(string)

	query := Attendance{
		UserId: userId,
	}

	attendance, err := ac.repository.FetchLatestAttendance(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{
			"message": err,
		})
		return
	}

	if err := ac.repository.CreateAttendanceTime(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if attendance == nil {
		attendance = &Attendance{
			UserId:    userId,
			ClockedIn: t,
		}
	} else {
		attendance = &Attendance{
			Id:         attendance.Id,
			UserId:     attendance.UserId,
			ClockedIn:  attendance.ClockedIn,
			ClockedOut: t,
		}
	}

	if _, err := ac.repository.CreateAttendance(attendance); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res := AttendanceResponse{}
	res.SetAttendance(attendance)

	c.JSON(http.StatusOK, H{
		"attendance": res,
	})
}

func (ac attendanceController) AttendanceMonthlyController(c *Context) {
	//p := NewPagination(0, 100)

	//if err := c.Bind(p); err != nil {
	//	c.JSON(http.StatusBadRequest, H{
	//		"message": err,
	//	})
	//	return
	//}
	//
	//value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	//if !exists {
	//	c.JSON(http.StatusNotFound, H{
	//		"message": "user not found",
	//	})
	//	return
	//}
	//
	//userId := value.(string)
	//
	//q := &Attendance{
	//	UserId: userId,
	//}
	//
	//attendances := make([]*Attendance, 0)
	//attendances, err := ac.repository.FetchAttendances(q, p)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, H{
	//		"message": err,
	//	})
	//	return
	//}

}
