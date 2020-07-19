package logger

import (
	"github.com/sirupsen/logrus"
)

func GetMysqlLogger(host string, port string, dbName string) *logrus.Logger {
	logger := logrus.New()
	logger.Out = GetMysqlWriter(host, port, dbName)
	return logger
}
