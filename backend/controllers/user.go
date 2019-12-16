package controllers

import (
	. "github.com/KouT127/Attendance-management/backend/domains"
	"github.com/KouT127/Attendance-management/backend/middlewares"
	. "github.com/KouT127/Attendance-management/backend/repositories"
	. "github.com/KouT127/Attendance-management/backend/validators"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func NewUserController(repo UserRepository) *userController {
	return &userController{
		repository: repo,
	}
}

type UserController interface {
	UserListController(c *Context)
	UserMineController(c *Context)
	UserUpdateController(c *Context)
}

type userController struct {
	repository UserRepository
}

func (uc userController) UserListController(c *Context) {
	u := &User{}
	users, err := uc.repository.FetchUsers(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"users": users,
	})
}

func getOrCreateUser(r UserRepository, userId string) (*User, error) {
	u := &User{}
	has, err := r.FetchUser(userId, u)

	if err != nil {
		return nil, err
	}

	if !has {
		u.Id = userId
		_, err := r.CreateUser(u)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

func (uc userController) UserMineController(c *Context) {
	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}
	userId := value.(string)
	u, err := getOrCreateUser(uc.repository, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"user": u,
	})
}

func (uc userController) UserCreateController(c *Context) {
	id := c.Request.Header.Get("id")

	u := &User{
		Id: id,
	}

	_, err := uc.repository.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"user": u,
	})
}

func (uc userController) UserUpdateController(c *Context) {
	var (
		input UserInput
	)

	err := c.Bind(&input)
	if err != nil {
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

	userId := value.(string)

	u := &User{}
	has, err := uc.repository.FetchUser(userId, u)
	if err != nil || !has {
		c.JSON(http.StatusNotFound, H{})
		return
	}

	u.Name = input.Name
	_, err = uc.repository.UpdateUser(u, &User{Id: u.Id})
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"user": u,
	})
}
