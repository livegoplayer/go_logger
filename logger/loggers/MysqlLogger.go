package loggers

import (
	"github.com/livegoplayer/go_logger/logger/writer"
	"github.com/sirupsen/logrus"
)

func GetMysqlLogger(host string, port string, dbName string, tableName string, username string, password string) *logrus.Logger {
	logger := logrus.New()
	logger.Out = writer.GetMysqlWriter(host, port, dbName, tableName, username, password)
	return logger
}
