package middlewares

import (
	"fmt"
	"github.com/KouT127/attendance-management/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

const AuthorizedUserKey = "AuthorizedUser"

func FetchAuthorizedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId string

		value, exists := c.Get(AuthorizedUserIdKey)
		if !exists {
			u := fmt.Sprintf("not found user id")
			c.AbortWithStatusJSON(http.StatusNotFound, u)
			return
		}

		if err := mapstructure.Decode(value, userId); err != nil {
			u := fmt.Sprintf("not found user id")
			c.AbortWithStatusJSON(http.StatusNotFound, u)
			return
		}

		user, err := fetchUser(userId)
		if err != nil {
			u := fmt.Sprintf("not found user id: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, u)
			return
		}

		c.Set(AuthorizedUserKey, user)
		c.Next()
	}
}

func fetchUser(userId string) (*models.User, error) {
	//engine := database.NewDB()
	var user models.User


	//if err != nil || len(results) == 0 {
	//	return nil, err
	//}
	//
	//err = mapstructure.Decode(results[0], &user)
	//
	//if err != nil {
	//	return nil	//results, err := engine.
	//	//	Table("users").
	//	//	Select("users.id, users.name").
	//	//	Where("id = ?", userId).
	//	//	QueryString()
	//	, err
	//}
	return &user, nil
}
