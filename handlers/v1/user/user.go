package user

import (
	"errors"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/auth"
	. "github.com/KouT127/attendance-management/modules/input"
	"github.com/KouT127/attendance-management/modules/logger"
	"github.com/KouT127/attendance-management/modules/response"
	. "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func MineHandler(c *Context) {
	value, exists := c.Get(auth.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}

	userId := value.(string)
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, err.Error())
		c.JSON(http.StatusBadRequest, response.NewValidationError("user", err))
		return
	}

	attendance, err := models.FetchLatestAttendance(userId)
	if err != nil {
		logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, err.Error())
		c.JSON(http.StatusBadRequest, response.NewValidationError("user", err))
		return
	}

	c.JSON(http.StatusOK, response.ToUserMineResult(user, attendance))
}

func UpdateHandler(c *Context) {
	input := new(UserInput)
	user := new(models.User)

	err := c.Bind(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewValidationError("user", err))
		return
	}

	value, exists := c.Get(auth.AuthorizedUserIdKey)
	if !exists {
		err := errors.New("user not found")
		c.JSON(http.StatusBadRequest, response.NewValidationError("user", err))
		return
	}

	user.Id = value.(string)
	user.Name = input.Name
	user.Email = input.Email

	if err := models.UpdateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, response.NewError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ToUserResult(user))
}
