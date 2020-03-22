package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func SetUp() {
	log.Out = os.Stdout
}

func NewInfo(msg string) {
	log.Info(msg)
}

func NewWarn(fields logrus.Fields, msg string) {
	log.WithFields(fields).Warn(msg)
}

func NewFatal(c *gin.Context, msg string) {
	log.WithFields(logrus.Fields{
		"host":   c.Request.URL.Host,
		"path":   c.Request.URL.Path,
		"params": c.Params,
		"header": c.Request.Header,
		"body":   c.Request.Body,
	}).Fatal(msg)
}
