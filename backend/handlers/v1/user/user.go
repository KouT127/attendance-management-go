package user

import (
	"errors"
	"github.com/KouT127/attendance-management/middlewares"
	"github.com/KouT127/attendance-management/responses"
	userService "github.com/KouT127/attendance-management/service/user"
	. "github.com/KouT127/attendance-management/usecases"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func UserMineHandler(c *Context) {
	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}
	userId := value.(string)
	u, a, err := userService.ViewUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	c.JSON(http.StatusOK, H{
		"user":       responses.NewUserResp(u),
		"attendance": responses.NewAttendanceResult(a),
	})
}

func UserUpdateHandler(c *Context) {
	var (
		input UserInput
	)

	err := c.Bind(&input)
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

	u, err := userService.UpdateUser(userId, input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}
	c.JSON(http.StatusOK, H{
		"user": responses.NewUserResp(u),
	})
}
