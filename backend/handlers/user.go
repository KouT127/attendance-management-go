package handlers

import (
	"errors"
	"github.com/KouT127/attendance-management/middlewares"
	"github.com/KouT127/attendance-management/responses"
	. "github.com/KouT127/attendance-management/usecases"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func NewUserHandler(uc UserUsecase) *userHandler {
	return &userHandler{
		usecase: uc,
	}
}

type UserHandler interface {
	UserMineHandler(c *Context)
	UserUpdateHandler(c *Context)
}

type userHandler struct {
	usecase UserUsecase
}

// @Description get mine information
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Accept  json
// @Produce  json
// @Success 200 {object} domains.User
// @Failure 400 {object} responses.CommonError
// @Router /users/mine [get]
func (uc userHandler) UserMineHandler(c *Context) {
	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}
	userId := value.(string)
	u, err := uc.usecase.ViewUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}

	c.JSON(http.StatusOK, H{
		"user": responses.NewUserResp(u),
	})
}

func (uc userHandler) UserUpdateHandler(c *Context) {
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

	u, err := uc.usecase.UpdateUser(userId, input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError("user", err))
		return
	}
	c.JSON(http.StatusOK, H{
		"user": responses.NewUserResp(u),
	})
}
