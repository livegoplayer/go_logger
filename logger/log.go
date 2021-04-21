package logger

import (
	"github.com/livegoplayer/go_logger/logger/writer"
	"github.com/rifflock/lfshook"
	"os"
	"runtime"
	"strconv"
	"time"

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
	level := logrus.FatalLevel
	logger := GetLoggerByLevel(level)
	logger.Fatal(getFileLine(), message)
}

func Error(message ...interface{}) {
	level := logrus.ErrorLevel
	logger := GetLoggerByLevel(level)
	logger.Error(getFileLine(), message)
}

func Warning(message ...interface{}) {
	level := logrus.WarnLevel
	logger := GetLoggerByLevel(level)
	logger.Warning(getFileLine(), message)
}

func Info(message ...interface{}) {
	level := logrus.InfoLevel
	logger := GetLoggerByLevel(level)
	logger.Info(getFileLine(), message)
}

func Debug(message ...interface{}) {
	level := logrus.DebugLevel
	logger := GetLoggerByLevel(level)
	logger.Debug(getFileLine(), message)
}

func Trace(message ...interface{}) {
	level := logrus.TraceLevel
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

var levelList = []logrus.Level{
	logrus.TraceLevel,
	logrus.DebugLevel,
	logrus.InfoLevel,
	logrus.WarnLevel,
	logrus.ErrorLevel,
	logrus.FatalLevel,
	logrus.PanicLevel,
}

// 初始化所有log为filelog
func InitBaseFileLogByPath(path string, cleanTime time.Duration, splitTime time.Duration) {
	for _, level := range levelList {
		appLogger := logrus.New()
		appLogger.Out = writer.GetFileWriter(path, level, cleanTime, splitTime)
		appLogger.Formatter = &logrus.JSONFormatter{}
		SetLogger(level, appLogger)
	}
}

// 初始化所有log为mysqllog
func InitBaseMysqlLogByConfig(host string, port string, dbName string, tableName string, username string, password string) {
	for _, level := range levelList {
		appLogger := logrus.New()
		appLogger.Out = writer.GetMysqlWriter(host, port, dbName, tableName, username, password)
		appLogger.Formatter = &logrus.JSONFormatter{}
		SetLogger(level, appLogger)
	}
}

// 初始化所有log为rabbitmq log
func InitBaseRabbitmqLogByConfig(url string, routerKey string, exchange string, appType int) {
	for _, level := range levelList {
		appLogger := logrus.New()
		appLogger.Out = writer.GetRabbitmqWriter(url, routerKey, exchange, appType)
		appLogger.Formatter = &logrus.JSONFormatter{}
		SetLogger(level, appLogger)
	}
}

// 为debug模式增加控制台输出
func AddDebugLogHook() {
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: os.Stdout, // 为不同级别设置不同的输出目的,这些都是ioWriter
		logrus.DebugLevel: os.Stdout, // 为不同级别设置不同的输出目的,这些都是ioWriter
		logrus.InfoLevel:  os.Stdout,
		logrus.WarnLevel:  os.Stdout,
		logrus.ErrorLevel: os.Stdout,
		logrus.FatalLevel: os.Stdout,
		logrus.PanicLevel: os.Stdout,
	}, &logrus.TextFormatter{})
	for _, level := range levelList {
		l := GetLoggerByLevel(level)
		l.AddHook(lfHook)
	}
}

func init() {
	LoggerMap = make(map[logrus.Level]*logrus.Logger, 0)
}
