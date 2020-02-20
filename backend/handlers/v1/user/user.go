package user

import (
	"errors"
	"github.com/KouT127/attendance-management/middlewares"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/responses"
	userService "github.com/KouT127/attendance-management/services/user"
	. "github.com/KouT127/attendance-management/validators"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func V1MineHandler(c *Context) {
	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}

	userId := value.(string)
	u, err := userService.GetOrCreateUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	attendance, err := models.FetchLatestAttendance(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	c.JSON(http.StatusOK, H{
		"user":       responses.NewUserResp(u),
		"attendance": responses.NewAttendanceResult(attendance),
	})
}

func V1UpdateHandler(c *Context) {
	input := new(UserInput)

	err := c.Bind(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		err := errors.New("user not found")
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	userId := value.(string)
	user, err := models.FetchUser(userId)
	if err != nil || user.Id == "" {
		err := errors.New("user not found")
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	if err := userService.UpdateUser(user, input.Name); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	c.JSON(http.StatusOK, H{
		"user": responses.NewUserResp(user),
	})
}
