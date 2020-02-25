package middlewares

import (
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
		app, err := firebase.NewApp(c, nil, *opt)
		if err != nil {
			logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, "error invalid credential file")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		client, err := app.Auth(c)
		if err != nil {
			logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, "error firebase unauthorized")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		header := c.Request.Header.Get("Authorization")
		replacedToken := strings.Replace(header, "Bearer ", "", 1)
		if replacedToken == "" {
			logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, "error verifying ID token")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		verifiedToken, err := client.VerifyIDToken(c, replacedToken)
		if err != nil {
			logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, "error verifying id token")
			u := fmt.Sprintf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		c.Set(auth.AuthorizedUserIdKey, verifiedToken.UID)
		c.Next()
	}
}
