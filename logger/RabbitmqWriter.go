package logger

import (
	"encoding/json"
	"strings"

	myHelper "github.com/livegoplayer/go_helper"
	mqHelper "github.com/livegoplayer/go_mq_helper/rabbitmq"
	"github.com/sirupsen/logrus"
)

//todo 以下两行移动到别的项目中去
const (
	GO_FILE_STORE = iota + 1
	GO_USER
	GO_USER_RPC
)

type LogMessage struct {
	Message string       `json:"message"`
	Level   logrus.Level `json:"level"`
	AppType int          `json:"app_type"`
}

//通过rabbitmq写入数据库
type RabbitmqWriter struct {
	Url        string
	RoutingKey string
	Exchange   string
	AppType    int
}

func GetRabbitmqWriter(url string, routerKey string, exchange string, appType int) *RabbitmqWriter {
	mqHelper.InitMqChannel(url)

	rw := &RabbitmqWriter{
		Url:        url,
		RoutingKey: routerKey,
		Exchange:   exchange,
		AppType:    appType,
	}

	return rw
}

func (rw *RabbitmqWriter) Write(p []byte) (n int, err error) {
	n = 0
	err = nil

	//解析出level_no
	str := myHelper.BytesToString(p)
	level := myHelper.GetSubStringBetween(str, "level=", " ")
	msg := strings.Trim(strings.Trim(myHelper.GetSubStringBetween(str, "msg=", ""), "\""), "\"\n")
	time := strings.Trim(myHelper.GetSubStringBetween(str, "time=", " "), "\"")

	message := time + " " + msg

	levelNo, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	//组装message
	logMessage := &LogMessage{
		Message: message,
		Level:   levelNo,
		AppType: rw.AppType,
	}

	logMsg, err := json.Marshal(logMessage)
	if err != nil {
		panic(err)
	}

	mqMessage := &mqHelper.Message{
		Message:    logMsg,
		RetryTimes: 0,
		Exchange:   rw.Exchange,
		RoutingKey: rw.RoutingKey,
	}

	mqHelper.Publish(mqMessage)

	return len(p), nil
}
