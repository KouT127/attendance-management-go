package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/KouT127/attendance-management/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"net/http"
	"strings"
)

const AuthorizedUserIdKey = "authorized_user_id"

func loadCredFromFile(name string) *option.ClientOption {
	filename := fmt.Sprintf(name)
	opt := option.WithCredentialsFile(filename)
	return &opt
}

func loadCredFromCtx() *option.ClientOption {
	cred, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		return nil
	}
	opt := option.WithCredentials(cred)
	return &opt
}

func NewCredential() *option.ClientOption {
	opt := loadCredFromCtx()
	if opt == nil {
		opt = loadCredFromFile("./backend/configs/firebase-service-dev.json")
	}
	return opt
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := NewCredential()
		app, err := firebase.NewApp(context.Background(), nil, *opt)
		if err != nil {
			logger.NewFatal(c, "error invalid credential file")
			u := fmt.Sprintf("bad access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			logger.NewFatal(c, "error firebase unauthorized")
			u := fmt.Sprintf("bad access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		header := c.Request.Header.Get("Authorization")
		replacedToken := strings.Replace(header, "Bearer ", "", 1)
		if replacedToken == "" {
			logger.NewFatal(c, "error verifying ID token")
			u := fmt.Sprintf("bad access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		verifiedToken, err := client.VerifyIDToken(context.Background(), replacedToken)
		if err != nil {
			logger.NewWarn(logrus.Fields{"Header": c.Request.Header}, "error verifying id token")
			u := fmt.Sprintf("bad access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, u)
			return
		}
		c.Set(AuthorizedUserIdKey, verifiedToken.UID)
		c.Next()
	}
}
