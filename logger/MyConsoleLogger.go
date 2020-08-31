package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetConsoleLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	return logger
}
