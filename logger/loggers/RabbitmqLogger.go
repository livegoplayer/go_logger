package loggers

import (
	"github.com/livegoplayer/go_logger/logger/writer"
	"github.com/sirupsen/logrus"
)

func GetRabbitMqLogger(url string, routerKey string, exchange string, appType int) *logrus.Logger {
	logger := logrus.New()
	logger.Out = writer.GetRabbitmqWriter(url, routerKey, exchange, appType)
	return logger
}
