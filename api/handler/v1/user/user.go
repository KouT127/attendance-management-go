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

type Handler interface {
	MineHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) Handler {
	return userHandler{
		service: service,
	}
}

func (h userHandler) MineHandler(c *gin.Context) {
	value, exists := c.Get(auth.AuthorizedUserIDKey)
	if !exists {
		logger.NewInfo("トークンが存在しません")
		c.JSON(http.StatusBadRequest, responses.NewError("不正なリクエストです"))
		return
	}

	userID, has := value.(string)
	if !has {
		logger.Warn("ユーザーIDが存在しません")
		c.JSON(http.StatusBadRequest, responses.NewError("不正なリクエストです"))
		return
	}

	params := models.GetOrCreateUserParams{UserID: userID}
	res, err := h.service.GetOrCreateUser(params)
	if err != nil {
		logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, err.Error())
		c.JSON(http.StatusBadRequest, responses.NewError("ユーザーが取得できませんでした"))
		return
	}

	c.JSON(http.StatusOK, responses.ToUserMineResult(res.User))
}

func (h userHandler) UpdateHandler(c *gin.Context) {
	input := payloads.UserPayload{}
	user := &models.User{}

	if err := c.Bind(&input); err != nil {
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	value, exists := c.Get(auth.AuthorizedUserIDKey)
	if !exists {
		err := xerrors.New("user not found")
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	user.ID = value.(string)
	user.Name = input.Name
	user.Email = input.Email
	user.ImageURL = input.ImageURL

	if err := h.service.UpdateUser(user); err != nil {
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.ToUserResult(user))
}
