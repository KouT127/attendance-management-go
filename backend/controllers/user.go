package controllers

import (
	"github.com/KouT127/Attendance-management/backend/database"
	. "github.com/KouT127/Attendance-management/backend/domains"
	"github.com/KouT127/Attendance-management/backend/middlewares"
	. "github.com/KouT127/Attendance-management/backend/validators"
	. "github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type UserController struct{}

func (uc UserController) UserListController(c *Context) {
	var users []*User
	engine := database.NewDB()

	results, err := engine.
		Table("users").
		Select("users.*").
		QueryString()

	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	for _, result := range results {
		var user User
		err := mapstructure.Decode(result, user)
		if err != nil {
			c.JSON(http.StatusBadRequest, H{})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, H{
		"users": users,
	})
}

func (uc UserController) UserMineController(c *Context) {
	var (
		user   *User
		userId string
	)

	value, exists := c.Get(middlewares.AuthorizedUserIdKey)
	if !exists {
		c.JSON(http.StatusBadRequest, H{
			"message": "user not found",
		})
		return
	}
	err := mapstructure.Decode(value, &userId)
	if err != nil {
		c.JSON(http.StatusNotFound, H{})
		return
	}
	user, err = getOrCreateUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"user": user,
	})
}

func getOrCreateUser(userId string) (*User, error) {
	engine := database.NewDB()
	user := User{
		Id: userId,
	}

	has, err := engine.
		Table("users").
		Where("id = ?", userId).
		Get(&user)

	if err != nil {
		return nil, err
	}

	if !has {
		_, err := engine.Table("users").Insert(&user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (uc UserController) UserCreateController(c *Context) {
	engine := database.NewDB()
	id := c.Request.Header.Get("id")
	user := User{
		id,
		"",
		"",
		"",
	}

	_, err := engine.Table("users").Insert(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"user": user,
	})
}

func (uc UserController) UserUpdateController(c *Context) {
	var (
		user  User
		input UserInput
	)

	id := c.Param("id")
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	//err = input.Validate()
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, H{
	//		"message": err,
	//	})
	//	return
	//}

	engine := database.NewDB()
	has, err := engine.
		Table("users").
		Select("users.*").
		Where("id = ?", id).
		Get(&user)

	if err != nil || !has {
		c.JSON(http.StatusNotFound, H{})
		return
	}

	user.Name = input.Name
	_, err = engine.Table("users").Update(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{})
		return
	}

	c.JSON(http.StatusOK, H{
		"user": user,
	})
}
