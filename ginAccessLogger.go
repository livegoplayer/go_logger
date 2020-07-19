package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	myhelper "github.com/livegoplayer/go_helper"
)

//获取gin专用的access log文件输入
func GetGinAccessFileLogger(logPath string, accessLogFileName string) gin.HandlerFunc {
	if logPath == "" {
		logPath = "../log"
	}

	if accessLogFileName == "" {
		accessLogFileName = "access.log"
	}

	if !myhelper.Exists(myhelper.PathToCommon(logPath)) {
		err := os.MkdirAll(logPath, os.ModeDir)
		if err != nil {
			panic("创建日志文件目录失败")
		}
	}

	if !myhelper.Exists(myhelper.PathToCommon(logPath + "/" + accessLogFileName)) {
		file, err := os.Create(logPath + "/" + accessLogFileName)
		if err != nil {
			panic("创建日志文件失败")
		}
		err = os.Chmod(logPath+"/"+accessLogFileName, 0777)
		if err != nil {
			panic("修改文件权限失败")
		}
		file.Close()
	}

	accessLogFile, err := os.OpenFile(logPath+"/"+accessLogFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic("打开日志文件失败" + err.Error())
	}

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

	return myhelper.BytesToString(myhelper.JsonEncode(accessLogBody))
}
