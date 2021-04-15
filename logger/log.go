package logger

import (
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

// 初始化日志
var LoggerMap map[logrus.Level]*logrus.Logger

//主动log的一些方法
func Panic(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Panic(getFileLine(), message)
}

func getFileLine() string {
	_, file, line, _ := runtime.Caller(2)
	return file + " : " + strconv.Itoa(line) + " "
}

func Fatal(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Fatal(getFileLine(), message)
}

func Error(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Error(getFileLine(), message)
}

func Warning(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Warning(getFileLine(), message)
}

func Info(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Info(getFileLine(), message)
}

func Debug(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Debug(getFileLine(), message)
}

func Trace(message ...interface{}) {
	level := logrus.PanicLevel
	logger := GetLoggerByLevel(level)
	logger.Trace(getFileLine(), message)
}

func SetLogger(level logrus.Level, logger *logrus.Logger) {
	if _, ok := LoggerMap[level]; ok {
		text, _ := level.MarshalText()
		panic(string(text) + "等级的日志已经设置完毕，请不要重复设置")
	}

	LoggerMap[level] = logger
}

func GetLoggerByLevel(level logrus.Level) *logrus.Logger {
	if m, ok := LoggerMap[level]; ok {
		return m
	} else {
		text, _ := level.MarshalText()
		panic(string(text) + "等级的日志尚未设置，请调用SetLogger方法设置")
	}
	return nil
}

func GetGinAccessLogger(level logrus.Level) *logrus.Logger {
	if m, ok := LoggerMap[level]; ok {
		return m
	} else {
		text, _ := level.MarshalText()
		panic(string(text) + "等级的日志尚未设置，请调用SetLogger方法设置")
	}
	return nil
}
