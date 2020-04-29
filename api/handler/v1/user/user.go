package user

import (
	"github.com/KouT127/attendance-management/api/payloads"
	"github.com/KouT127/attendance-management/api/responses"
	"github.com/KouT127/attendance-management/application/services"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/auth"
	"github.com/KouT127/attendance-management/utilities/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/http"
)

type UserHandler interface {
	MineHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}

type userHandler struct {
	facade services.UserService
}

func NewUserHandler(facade services.UserService) UserHandler {
	return userHandler{
		facade: facade,
	}
}

func (h userHandler) MineHandler(c *gin.Context) {
	value, exists := c.Get(auth.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	userId := value.(string)
	params := models.GetOrCreateUserParams{UserId: userId}
	res, err := h.facade.GetOrCreateUser(params)
	if err != nil {
		logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, err.Error())
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	c.JSON(http.StatusOK, responses.ToUserMineResult(res.User, res.LatestAttendance))
}

func (h userHandler) UpdateHandler(c *gin.Context) {
	input := payloads.UserPayload{}
	user := &models.User{}

	err := c.Bind(&input)
	if err != nil {
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	value, exists := c.Get(auth.AuthorizedUserIdKey)
	if !exists {
		err := xerrors.New("user not found")
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	user.Id = value.(string)
	user.Name = input.Name
	user.Email = input.Email

	if err := h.facade.UpdateUser(user); err != nil {
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.ToUserResult(user))
}
