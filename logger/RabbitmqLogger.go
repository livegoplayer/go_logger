package logger

import (
	"github.com/sirupsen/logrus"
)

func GetRabbitMqLogger(url string, routerKey string, exchange string, appType int) *logrus.Logger {
	logger := logrus.New()
	logger.Out = GetRabbitmqWriter(url, routerKey, exchange, appType)
	return logger
}
