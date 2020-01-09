package handlers

import (
	"errors"
	"github.com/KouT127/attendance-management/backend/middlewares"
	. "github.com/KouT127/attendance-management/backend/serializers"
	. "github.com/KouT127/attendance-management/backend/usecases"
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
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	c.JSON(http.StatusOK, H{
		"user": u,
	})
}

func (uc userHandler) UserUpdateHandler(c *Context) {
	var (
		input UserInput
	)

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		err := errors.New("user not found")
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}

	userId := value.(string)

	u, err := uc.usecase.UpdateUser(userId, input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError("user", err))
		return
	}
	c.JSON(http.StatusOK, H{
		"user": u,
	})
}
