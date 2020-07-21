package logger

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

//主动log的一些方法
func Panic(message ...interface{}) {
	Logger = GetLogger()
	Logger.Panic(message)
}

func Fatal(message ...interface{}) {
	Logger = GetLogger()
	Logger.Fatal(message)
}

func Error(message ...interface{}) {
	Logger = GetLogger()
	Logger.Error(message)

}

func Warning(message ...interface{}) {
	Logger = GetLogger()
	Logger.Warn(message)
}

func Info(message ...interface{}) {
	Logger = GetLogger()
	Logger.Info(message)
}

func Debug(message ...interface{}) {
	Logger = GetLogger()
	Logger.Debug(message)
}

func Trace(message ...interface{}) {
	Logger = GetLogger()
	Logger.Trace(Logger)
}

func GetLogger() *logrus.Logger {
	if Logger == nil {
		return logrus.New()
	}

	return Logger
}

//设置默认的logger
func SetLogger(logger *logrus.Logger) {
	Logger = logger
}
