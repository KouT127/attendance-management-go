package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func Init() {
	log.Out = os.Stdout
}

func NewInfo(fields map[string]interface{}, msg string) {
	log.WithFields(fields).Info(msg)
}

func NewWarn(fields logrus.Fields, msg string) {
	log.WithFields(fields).Warn(msg)
}

func NewFatal(fields logrus.Fields, msg string) {
	log.WithFields(fields).Fatal(msg)
}
