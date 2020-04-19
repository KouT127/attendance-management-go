package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/KouT127/attendance-management/modules/auth"
	"github.com/KouT127/attendance-management/modules/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := auth.NewCredential()
		app, err := firebase.NewApp(context.Background(), nil, *opt)
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error invalid credential file")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error firebase unauthorized")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		header := c.Request.Header.Get("Authorization")
		replacedToken := strings.Replace(header, "Bearer ", "", 1)
		if replacedToken == "" {
			logger.NewWarn(logrus.Fields{"err": err}, "error verifying ID token")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		verifiedToken, err := client.VerifyIDToken(context.Background(), replacedToken)
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error verifying id token")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		c.Set(auth.AuthorizedUserIDKey, verifiedToken.UID)
		c.Next()
	}
}
