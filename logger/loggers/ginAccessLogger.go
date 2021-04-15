package loggers

import (
	"github.com/gin-gonic/gin"
	"github.com/livegoplayer/go_logger/logger/writer"
	"github.com/sirupsen/logrus"
	"time"
)

//获取gin专用的access log文件输入
func GetGinAccessFileLogger(logPath string, accessLogFileName string) gin.HandlerFunc {
	if logPath == "" {
		logPath = "../access_log"
	}

	accessLogFile := writer.GetFileWriter(logPath, logrus.InfoLevel, time.Hour*24*60, time.Hour*24)

	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: MyGinLoggerFormatter, Output: accessLogFile,
	})
}

type loggerParams struct {
	QueryUrl      string              //请求url
	Method        string              //请求方法
	Proto         string              //请求方法
	PostForm      map[string][]string //post body
	RequestHeader map[string][]string //requestHeader
	TimeStamp     string              //请求时间
	StatusCode    int                 //请求回应状态
	Latency       string              //请求时长
	ClientIP      string              //客户端IP
}

//专门给gin框架定制的access log文件输出模式
//直接json encode
func MyGinLoggerFormatter(params gin.LogFormatterParams) string {
	accessLogBody := &loggerParams{
		QueryUrl:      params.Request.Host + params.Path,
		Proto:         params.Request.Proto,
		Method:        params.Method,
		PostForm:      params.Request.Form,
		RequestHeader: params.Request.Header,
		TimeStamp:     params.TimeStamp.String(),
		StatusCode:    params.StatusCode,
		Latency:       params.Latency.String(),
		ClientIP:      params.ClientIP,
	}

	return writer.JsonEncode(accessLogBody)
}
