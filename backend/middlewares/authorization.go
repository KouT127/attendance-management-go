package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/KouT127/attendance-management/configs"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"net/http"
	"strings"
)

const AuthorizedUserIdKey = "authorized_user_id"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := configs.NewConfig()
		if conf == nil {
			u := fmt.Sprintf("error firebase unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		filename := fmt.Sprintf("%s/firebase-service-stg.json", "./backend/configs")
		opt := option.WithCredentialsFile(filename)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			u := fmt.Sprintf("error invalid credential file")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			u := fmt.Sprintf("error firebase unauthorized: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		header := c.Request.Header.Get("Authorization")
		replacedToken := strings.Replace(header, "Bearer ", "", 1)
		if replacedToken == "" {
			u := fmt.Sprintf("error verifying ID token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		verifiedToken, err := client.VerifyIDToken(context.Background(), replacedToken)
		if err != nil {
			u := fmt.Sprintf("error verifying ID token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		c.Set(AuthorizedUserIdKey, verifiedToken.UID)
		c.Next()
	}
}
