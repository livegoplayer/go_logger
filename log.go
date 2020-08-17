package logger

import (
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

//主动log的一些方法
func Panic(message ...interface{}) {
	Logger = GetLogger()
	Logger.Panic(getFileLine(), message)
}

func getFileLine() string {
	_, file, line, _ := runtime.Caller(1)
	return file + " : " + strconv.Itoa(line) + " "
}

func Fatal(message ...interface{}) {
	Logger = GetLogger()
	Logger.Fatal(getFileLine(), message)
}

func Error(message ...interface{}) {
	Logger = GetLogger()
	Logger.Error(getFileLine(), message)

}

func Warning(message ...interface{}) {
	Logger = GetLogger()
	Logger.Warn(getFileLine(), message)
}

func Info(message ...interface{}) {
	Logger = GetLogger()
	Logger.Info(getFileLine(), message)
}

func Debug(message ...interface{}) {
	Logger = GetLogger()
	Logger.Debug(getFileLine(), message)
}

func Trace(message ...interface{}) {
	Logger = GetLogger()
	Logger.Trace(getFileLine(), Logger)
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
