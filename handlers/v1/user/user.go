package user

import (
	"github.com/KouT127/attendance-management/application/facades"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/auth"
	"github.com/KouT127/attendance-management/modules/logger"
	"github.com/KouT127/attendance-management/modules/payloads"
	"github.com/KouT127/attendance-management/modules/responses"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/http"
)

func MineHandler(c *gin.Context) {
	store := sqlstore.InitDatabase()
	facade := facades.NewUserFacade(store)
	value, exists := c.Get(auth.AuthorizedUserIDKey)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	userID := value.(string)
	params := models.GetOrCreateUserParams{UserID: userID}
	res, err := facade.GetOrCreateUser(params)
	if err != nil {
		logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, err.Error())
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	c.JSON(http.StatusOK, responses.ToUserMineResult(res.User, res.LatestAttendance))
}

func UpdateHandler(c *gin.Context) {
	store := sqlstore.InitDatabase()
	facade := facades.NewUserFacade(store)

	input := payloads.UserPayload{}
	user := &models.User{}

	err := c.Bind(&input)
	if err != nil {
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

	if err := facade.UpdateUser(user); err != nil {
		logrus.Warnf("not exists: %s", err)
		c.JSON(http.StatusBadRequest, responses.NewError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.ToUserResult(user))
}
