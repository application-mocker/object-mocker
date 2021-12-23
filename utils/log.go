package utils

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.TraceLevel)
	Logger.Info("Init logger complete")
}
