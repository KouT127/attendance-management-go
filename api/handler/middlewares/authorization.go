package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/KouT127/attendance-management/infrastructure/auth"
	"github.com/KouT127/attendance-management/utilities/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			app *firebase.App
			err error
		)
		opt, err := auth.NewCredential()
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error credential is not exists")
			err = fmt.Errorf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		app, err = firebase.NewApp(context.Background(), nil, *opt)
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error invalid credential file")
			err = fmt.Errorf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error firebase unauthorized")
			err = fmt.Errorf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		header := c.Request.Header.Get("Authorization")
		replacedToken := strings.Replace(header, "Bearer ", "", 1)
		if replacedToken == "" {
			logger.NewWarn(logrus.Fields{"err": err}, "error verifying ID token")
			err = fmt.Errorf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		verifiedToken, err := client.VerifyIDToken(context.Background(), replacedToken)
		if err != nil {
			logger.NewWarn(logrus.Fields{"err": err}, "error verifying id token")
			err = fmt.Errorf("unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		c.Set(auth.AuthorizedUserIDKey, verifiedToken.UID)
		c.Next()
	}
}
